/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"DataCompare/handler/router/baseRouter"
	"DataCompare/handler/router/smRouter"
	"DataCompare/handler/router/userRouter"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	defaultRouter := gin.Default()
	apiGroup := defaultRouter.Group("api/v1")
	baseRouter.InitBaseRouter(apiGroup) // base 通用
	userRouter.InitUserRouter(apiGroup) // user 用户路由
	smRouter.InitSmRouter(apiGroup)     // 短信验证码
	return defaultRouter
}
