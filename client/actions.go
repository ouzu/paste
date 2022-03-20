package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

type FileData struct {
	Name string
	Type string
	Size int64
}

type PutResp struct {
	Error string `json:error`
}

func buildURL(c *cli.Context, path string) string {
	url := url.URL{
		Scheme: "https",
		Host:   c.String("server"),
		Path:   path,
	}

	return url.String()
}

func UploadHandler(c *cli.Context) error {
	debug := c.Bool("debug")

	var err error
	var f FileData

	file := os.Stdin
	f.Name = c.Args().First()

	if f.Name != "" {
		if debug {
			log.Println("[file] reading from", f.Name)
		}
		file, err = os.Open(f.Name)
		if err != nil {
			return err
		}
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	f.Type = http.DetectContentType(data)
	f.Size = int64(len(data))

	if debug {
		log.Println("[meta] name:", f.Name)
		log.Println("[meta] type:", f.Type)
		log.Println("[meta] size:", f.Size)
	}

	iv := generateIV()
	key := generateKey()

	if debug {
		log.Println("[encode] iv:", iv)
		log.Println("[encode] key:", key)
	}

	ivStr := exportB58(iv)
	keyStr := exportB58(key)

	if debug {
		log.Println("[encode] iv:", ivStr)
		log.Println("[encode] key:", keyStr)
	}

	f.Name = exportB58(encrypt([]byte(f.Name), iv, key))
	f.Type = exportB58(encrypt([]byte(f.Type), iv, key))

	if debug {
		log.Println("[meta] name:", f.Name)
		log.Println("[meta] type:", f.Type)
		log.Println("[meta] size:", f.Size)
	}

	cryptData := encrypt(data, iv, key)

	if debug {
		log.Println("[info] creating multipart form")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", ivStr)
	if err != nil {
		return err
	}
	io.Copy(part, bytes.NewReader(cryptData))

	part, err = writer.CreateFormField("name")
	if err != nil {
		return err
	}
	io.Copy(part, strings.NewReader(f.Name))

	part, err = writer.CreateFormField("type")
	if err != nil {
		return err
	}
	io.Copy(part, strings.NewReader(f.Type))
	writer.Close()

	if debug {
		log.Println("[info] creating http request")
	}

	req, err := http.NewRequest(http.MethodPut, buildURL(c, "/api/files/"+ivStr), body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("[fatal] http request failed")
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		r := &PutResp{}
		err = json.Unmarshal(b, r)
		if err != nil {
			return err
		}

		return errors.New(r.Error)
	}

	fmt.Println(buildURL(c, "/") + "#/files/" + ivStr + "/" + keyStr)

	return nil
}

func DownloadHandler(c *cli.Context) error {
	debug := c.Bool("debug")

	u, err := url.Parse(c.Args().First())
	if err != nil {
		return err
	}

	if debug {
		log.Println("[URL] server:", u.Host)
	}

	if u.Host != c.String("server") {
		return errors.New("server mismatch")
	}

	r := regexp.MustCompile(`^\/files\/([1-9A-HJ-NP-Za-km-z]*)\/([1-9A-HJ-NP-Za-km-z]*)\/?$`)

	if !r.Match([]byte(u.Fragment)) {
		return errors.New("invalid url")
	}

	m := r.FindStringSubmatch(u.Fragment)
	ivStr := m[1]
	keyStr := m[2]

	if debug {
		log.Println("[URL] iv:", ivStr)
		log.Println("[URL] key:", keyStr)
	}

	iv := importB58(ivStr)
	key := importB58(keyStr)

	if debug {
		log.Println("[decode] iv:", iv)
		log.Println("[decode] key:", key)
	}

	name := c.String("name")

	if name == "" {
		metaResp, err := http.Get(buildURL(c, "api/meta/"+ivStr))
		if err != nil {
			return err
		}
		defer metaResp.Body.Close()

		metaBody, err := ioutil.ReadAll(metaResp.Body)
		if err != nil {
			return err
		}

		var f FileData
		err = json.Unmarshal(metaBody, &f)
		if err != nil {
			return err
		}

		if debug {
			log.Println("[meta] name:", f.Name)
			log.Println("[meta] type:", f.Type)
			log.Println("[meta] size:", f.Size)
		}

		name = string(decrypt(importB58(f.Name), iv, key))
	}

	if debug {
		log.Println("[meta] name:", name)
	}

	fileResp, err := http.Get(buildURL(c, "api/files/"+ivStr))
	if err != nil {
		return err
	}
	defer fileResp.Body.Close()

	cryptFile, err := ioutil.ReadAll(fileResp.Body)
	if err != nil {
		return err
	}

	_, err = os.Stat(name)
	if err == nil {
		return errors.New("file exists")
	}

	file := decrypt(cryptFile, iv, key)

	err = ioutil.WriteFile(name, file, 0600)
	if err != nil {
		return err
	}

	return nil
}
