package main

import (
	"github.com/gin-gonic/gin"

	"bj38web/web/controller"
)

// 添加gin框架开发3步骤
func main()  {
	// 初始化路由
	router := gin.Default()

	// 路由匹配
/*	router.GET("/", func(context *gin.Context) {
		context.Writer.WriteString("项目开始了....")
	})*/
	router.Static("/home", "view")

	router.GET("/api/v1.0/session", controller.GetSession)

	// 启动运行
	router.Run(":8080")
}
