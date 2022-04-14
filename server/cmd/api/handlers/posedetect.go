package handlers

import (
	"io"
	"log"
	"os"
	"server/cmd/posedetect"
	"time"

	"github.com/gin-gonic/gin"
)

func PoseDetect(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	file, header, _ := c.Request.FormFile("image")
	filename := header.Filename

	filePath := "uploads/" + filename

	out, err := os.Create(filePath)

	defer os.Remove(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	fileName := posedetect.Detect(filePath)

	c.JSON(200, gin.H{"createdAt": time.Now(), "path": "http://localhost:8000/uploads/" + fileName})
}