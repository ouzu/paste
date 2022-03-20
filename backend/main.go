package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	bolt "go.etcd.io/bbolt"
)

type FileData struct {
	Name string
	Type string
	Size int64
}

func (f FileData) JSON() []byte {
	b, _ := json.Marshal(&f)
	return b
}

func main() {
	dataDir := os.Getenv("PASTE_DATA_DIR")
	if dataDir == "" {
		dataDir = "./data"
	}

	fileDir := filepath.Join(dataDir, "files")

	os.MkdirAll(fileDir, 0755)

	frontendDir := os.Getenv("PASTE_FRONTEND_DIR")
	if frontendDir == "" {
		frontendDir = "../frontend/public"
	}

	releaseDir := os.Getenv("PASTE_RELEASE_DIR")
	if releaseDir == "" {
		releaseDir = "../client/release"
	}

	db, err := bolt.Open(filepath.Join(dataDir, "meta.db"), 0666, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	r := gin.Default()

	r.PUT("/api/files/:iv", func(c *gin.Context) {
		iv := c.Param("iv")

		filePath := filepath.Join(fileDir, iv)

		fileName := c.PostForm("name")
		if fileName == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no file name provided",
			})
			return
		}

		fileType := c.PostForm("type")
		if fileType == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no file type provided",
			})
			return
		}

		file, err := c.FormFile("file")

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "no file received",
			})
			return
		}

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not save",
			})
			return
		}

		fi, err := os.Stat(filePath)
		if err != nil {
			log.Println(err)
			os.Remove(filePath)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not stat",
			})
			return
		}

		if err = db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucketIfNotExists([]byte("meta"))
			if err != nil {
				log.Println(err)
				return err
			}

			b.Put([]byte(iv), FileData{
				fileName,
				fileType,
				fi.Size(),
			}.JSON())

			return nil

		}); err != nil {
			os.Remove(filePath)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not open bucket",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/api/meta/:iv", func(c *gin.Context) {
		iv := c.Param("iv")

		db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("meta"))

			data := b.Get([]byte(iv))

			if len(data) == 0 {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": "entry not found",
				})
				return nil
			}

			c.Data(http.StatusOK, "application/json", data)

			return nil
		})
	})

	r.StaticFS("/api/files", gin.Dir(fileDir, false))
	r.StaticFS("/releases", gin.Dir(releaseDir, true))

	r.StaticFS("/build", gin.Dir(filepath.Join(frontendDir, "build"), false))
	r.StaticFile("/", filepath.Join(frontendDir, "index.html"))
	r.StaticFile("/global.css", filepath.Join(frontendDir, "global.css"))
	r.StaticFile("/favicon.png", filepath.Join(frontendDir, "favicon.png"))

	if port := os.Getenv("HTTPS_PORT"); port != "" {
		go r.RunTLS(
			"0.0.0.0:"+port,
			filepath.Join(dataDir, "paste.cert"),
			filepath.Join(dataDir, "paste.key"),
		)
	}

	r.Run()
}
