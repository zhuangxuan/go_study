package main

import (
	"bj38web/web/controller"
	"bj38web/web/model"
	"github.com/gin-gonic/gin"
)

// 添加gin框架开发3步骤
func main() {

	// 初始化 Redis 链接池
	model.InitRedis()

	// 初始化路由
	router := gin.Default()

	// 路由匹配
	/*	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目开始了....")
	})*/
	router.Static("/home", "view")

	//router.GET("/api/v1.0/session", controller.GetSession)
	//router.GET("/api/v1.0/imagecode/:uuid", controller.GetImageCd)
	//router.GET("/api/v1.0/smscode/:phone")

	// 添加路由分组
	r1 := router.Group("/api/v1.0")
	{
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:phone", controller.GetSmscd)
		r1.POST("/users", controller.PostRet)
	}

	// 启动运行
	router.Run(":8080")
}
