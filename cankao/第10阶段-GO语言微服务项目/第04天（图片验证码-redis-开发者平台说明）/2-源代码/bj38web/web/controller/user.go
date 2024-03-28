package controller

import (
	"github.com/gin-gonic/gin"
	"bj38web/web/utils"
	"net/http"
	"fmt"
	getCaptcha "bj38web/web/proto/getCaptcha"   // 给包起别名
	"image/png"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
)

// 获取 session 信息.
func GetSession(ctx *gin.Context) {
	// 初始化错误返回的 map
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}

// 获取图片信息
func GetImageCd(ctx *gin.Context) {
	// 获取图片验证码 uuid
	uuid := ctx.Param("uuid")

	// 指定 consul 服务发现
	consulReg := consul.NewRegistry()
	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("getCaptcha", consulService.Client())

	// 调用远程函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid:uuid})
	if err != nil {
		fmt.Println("未找到远程服务...")
		return
	}

	// 将得到的数据,反序列化,得到图片数据
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 将图片写出到 浏览器.
	png.Encode(ctx.Writer, img)

	fmt.Println("uuid = ", uuid)
}
