/*
 * @date: 2021/8/19
 * @desc: ...
 */

package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response base response
func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}

// Success  response
func Success(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

// Failed response
func Failed(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, 400, data, msg)
}
