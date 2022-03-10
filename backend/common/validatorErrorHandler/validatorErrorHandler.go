/*
 * @date: 2021/12/15
 * @desc: ...
 */

package validatorErrorHandler

import (
	"DataCompare/global"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

// ValidatorErrorHandler 格式化form表单验证信息
func ValidatorErrorHandler(ctx *gin.Context, err error) {
	// 返回错误信息
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
	return
}

// removeTopStruct 将返回的map的key格式化
func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
