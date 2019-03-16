package routers

import (
	"adminVideos/routers/api"
	"adminVideos/routers/api/v1"
	"adminVideos/routers/api/v3"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	//静态文件服务 提供图片路径
	r.Static("/images","./public/upload/images")
	r.Static("/static/images","./public/images")
	r.Static("/videos","./public/upload/videos")
	r.Static("/index.html","./public/dist")
	
	//登录模块
	api1 := r.Group("/api/v1")
	{
		api1.POST("/login",v1.Login)
		api1.POST("/registered",v1.Registered)
	}

	//用户管理模块
	api2 := r.Group("/api/v2")
	api2.Use(api.GetAuth)
	{
			api2.GET("/aa", func(c *gin.Context) {
				id,exists:=c.Get("persons_id")
				if !exists{
					c.JSON(http.StatusOK, gin.H{
						"status" :400,
						"error": "没有拿到person_id",
						"data": "失败",
					})
				}
				c.JSON(http.StatusOK, gin.H{
					"status" :200,
					"error": nil,
					"data": "成功",
					"persons_id": id,
				})
			})
	}

	//视频管理模块
	api3 := r.Group("/api/v3")
	//api3.Use(api.GetAuth)
	{
		//上传接口
		api3.POST("/upload/videos",v3.UploadVideos) //上传视频
		api3.POST("/upload/images",v3.UploadImages) //上传视频封面

		// 操作视频表
		api3.GET("/videos",v3.FindVideos)
		api3.GET("/videos/:id",v3.GetVideo)
		api3.POST("/videos",v3.PostVideos)
		api3.PUT("/videos/:id",v3.UpdateVideos)
		api3.DELETE("/videos/:id",v3.DeleteVideos)



	}
	
	
	return r
}
