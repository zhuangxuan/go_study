package controller

import (
	"github.com/gin-gonic/gin"
	"bj38web/web/utils"
	"net/http"
)

func GetSession(ctx *gin.Context)  {
	// 初始化错误返回的 map
	resp := make(map[string]string)

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)

	ctx.JSON(http.StatusOK, resp)
}
