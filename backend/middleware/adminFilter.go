/**
 * @date: 2022/2/22
 * @desc: 确认管理员身份(鉴权)
 */

package middleware

import (
	"DataCompare/common/response"
	"DataCompare/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//
// AdminFilter
// @Description: 管理员权限过滤
// @return gin.HandlerFunc:
//
func AdminFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentUserInfo, ok := utils.GetCurrentUserInfo(ctx)
		if !ok {
			response.Response(ctx, http.StatusUnauthorized, 401, nil, "未登录")
			ctx.Abort()
			return
		}
		userName := currentUserInfo.Username
		if userName != "admin" {
			response.Response(ctx, http.StatusForbidden, 403, nil, "无权限")
			ctx.Abort()
			return
		}
		ctx.Next()

	}
}
