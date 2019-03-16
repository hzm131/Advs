package v3

import (
	"adminVideos/models/createtable"
	"adminVideos/models/operating"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func FindVideos(c *gin.Context) {
	getOffset := c.Query("offset")
	getLimit := c.Query("limit")

	//将string转换成int
	offset,err := strconv.Atoi(getOffset)
	if err != nil{
		fmt.Fprintln(c.Writer,"offset有问题")
		return
	}

	limit,err := strconv.Atoi(getLimit)
	if err != nil{
		fmt.Fprintln(c.Writer,"offset有问题")
		return
	}
	data := operating.GetVideos(offset,limit)

	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": data,
	})
}

func GetVideo(c *gin.Context) {
}

func PostVideos(c *gin.Context){

	var video createtable.Videos

	err := c.BindJSON(video) //将post的json封装进结构体
	if err != nil{
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :200,
		"error": nil,
		"data": video,
	})
}

func UpdateVideos(c *gin.Context){

}

func DeleteVideos(c *gin.Context){

}