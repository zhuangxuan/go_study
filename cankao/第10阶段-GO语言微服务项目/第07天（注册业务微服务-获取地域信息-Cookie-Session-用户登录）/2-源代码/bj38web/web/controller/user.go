package controller

import (
	"github.com/gin-gonic/gin"
	"bj38web/web/utils"
	"net/http"
	"fmt"
	getCaptcha "bj38web/web/proto/getCaptcha" // 给包起别名
	"image/png"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro"
	"context"
	"encoding/json"
	"github.com/afocus/captcha"

	userMicro "bj38web/web/proto/user" // 给包起别名
	"bj38web/web/model"
	"github.com/gomodule/redigo/redis"
	"github.com/gin-contrib/sessions"
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
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
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

// 获取短信验证码
func GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	phone := ctx.Param("phone")
	// 拆分 GET 请求中 的 URL === 格式: 资源路径?k=v&k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")

	// 指定Consul 服务发现
	consulReg := consul.NewRegistry()

	consulService := micro.NewService(
		micro.Registry(consulReg),
	)

	// 初始化客户端
	microClient := userMicro.NewUserService("go.micro.srv.user", consulService.Client())

	// 调用远程函数:
	resp, err := microClient.SendSms(context.TODO(), &userMicro.Request{Phone: phone, ImgCode: imgCode, Uuid: uuid})
	if err != nil {
		fmt.Println("调用远程函数 SendSms 失败:", err)
		return
	}

	// 发送校验结果 给 浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 发送注册信息
func PostRet(ctx *gin.Context) {
	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}

	ctx.Bind(&regData)

	// 初始化consul
	microService := utils.InitMicro()
	microClient := userMicro.NewUserService("go.micro.srv.user", microService.Client())

	// 调用远程函数
	resp, err := microClient.Register(context.TODO(), &userMicro.RegReq{
		Mobile:   regData.Mobile,
		SmsCode:  regData.SmsCode,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("注册用户, 找不到远程服务!", err)
		return
	}

	// 写给浏览器
	ctx.JSON(http.StatusOK, resp)
}

// 获取地域信息
func GetArea(ctx *gin.Context) {
	// 先从MySQL中获取数据.
	var areas []model.Area

	// 从缓存redis 中, 获取数据
	conn := model.RedisPool.Get()
	// 当初使用 "字节切片" 存入, 现在使用 切片类型接收
	areaData, _ := redis.Bytes(conn.Do("get", "areaData"))
	// 没有从 Redis 中获取到数据
	if len(areaData) == 0 {

		fmt.Println("从 MySQL 中 获取数据...")
		model.GlobalConn.Find(&areas)
		// 把数据写入到 redis 中. , 存储结构体序列化后的 json 串
		areaBuf, _ := json.Marshal(areas)
		conn.Do("set", "areaData", areaBuf)

	} else {
		fmt.Println("从 redis 中 获取数据...")
		// redis 中有数据
		json.Unmarshal(areaData, &areas)
	}

	resp := make(map[string]interface{})

	resp["errno"] = "0"
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas

	ctx.JSON(http.StatusOK, resp)
}

// 处理登录业务
func PostLogin(ctx *gin.Context) {
	// 获取前端数据
	var loginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	ctx.Bind(&loginData)

	resp := make(map[string]interface{})

	//获取 数据库数据, 查询是否和数据的数据匹配
	userName, err := model.Login(loginData.Mobile, loginData.PassWord)
	if err == nil {
		// 登录成功!
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

		// 将 登录状态, 保存到Session中

		s := sessions.Default(ctx)		// 初始化session
		s.Set("userName", userName)   // 将用户名设置到session中.
		s.Save()

	} else {
		// 登录失败!
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}

	ctx.JSON(http.StatusOK, resp)
}
