/**
 * @date: 2022/3/12
 * @desc: ...
 */

package middleware

import (
	"DataCompare/initialize/initCron"
	"github.com/gin-gonic/gin"
)

//
// ReInitCron
// @Description: 重新初始化 cron
// @return gin.HandlerFunc:
//
func ReInitCron() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		initCron.ChangeTaskModelEvent()
	}
}
