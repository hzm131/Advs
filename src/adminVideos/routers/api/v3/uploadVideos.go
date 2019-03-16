package v3

import (
	"adminVideos/models/operating"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)
//上传视频
func UploadVideos(c *gin.Context) {
	/*
		客户端上传的文件是key-value形式
	*/
	//通过key拿到文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "错误请求")
		return
	}
	//文件的名称
	filename := header.Filename

	fmt.Println(file, err, filename)

	//创建文件
	out, err := os.Create("public/upload/videos/"+filename)
	if err != nil {
		log.Fatal(err)
	}

	operating.CreatedImage("http://127.0.0.1:8000/videos/"+filename)

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "上传成功")
}

//上传视频封面
func UploadImages(c *gin.Context) {
	/*
		客户端上传的文件是key-value形式
	*/
	//通过key拿到文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "错误请求")
		return
	}
	//文件的名称
	filename := header.Filename

	fmt.Println(file, err, filename)

	//创建文件
	out, err := os.Create("public/upload/images/"+filename)
	if err != nil {
		log.Fatal(err)
	}

	id,err := operating.CreatedImage("http://127.0.0.1:8000/images/"+filename)
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": id,
	})
}