package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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
	db, err := bolt.Open("./files/meta.db", 0666, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	r := gin.Default()

	r.PUT("/api/files/:iv", func(c *gin.Context) {
		iv := c.Param("iv")

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

		if err := c.SaveUploadedFile(file, "./files/"+iv); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "could not save",
			})
			return
		}

		fi, err := os.Stat("./files/" + iv)
		if err != nil {
			log.Println(err)
			os.Remove("./files/" + iv)
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
			os.Remove("./files/" + iv)
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

	r.StaticFS("/api/files", http.Dir("files"))

	r.Run()
}
