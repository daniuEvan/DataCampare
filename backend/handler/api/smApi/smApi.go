/**
 * @date: 2022/2/18
 * @desc: ...
 */

package smApi

import (
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/api/userApi"
	"DataCompare/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"regexp"
)

//
// GetSmCode
// @Description: 获取验证码
// @param ctx:
//
func GetSmCode(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取验证码", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	mobile := ctx.Param("mobile")
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "手机号输入错误")
		return
	}
	// 验证用户是否存在
	if userApi.IsMobilExist(db, mobile) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经存在")
		return
	}
	// 获取验证码
	smService := utils.NewSmService()
	codeJson, err := smService.SendSmCode(mobile)
	if err != nil {
		global.Logger.Error("获取验证码", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	type MsgCode struct {
		Code int
		Msg  string
		Obj  string
	}
	msgCodeObj := &MsgCode{}
	_ = json.Unmarshal([]byte(codeJson), msgCodeObj)
	if msgCodeObj.Code != 200 {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	response.Success(ctx, gin.H{"ms": msgCodeObj}, "验证码")
}
