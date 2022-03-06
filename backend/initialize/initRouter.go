/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"DataCompare/handler/router/baseRouter"
	"DataCompare/handler/router/smRouter"
	"DataCompare/handler/router/taskRouter"
	"DataCompare/handler/router/userRouter"
	"DataCompare/middleware"
	"DataCompare/utils"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	defaultRouter := gin.New()
	if utils.IsDebugEnv() {
		defaultRouter = gin.Default()
	}
	defaultRouter.Use(middleware.Cors()) // 跨域
	apiGroup := defaultRouter.Group("api/v1")
	baseRouter.InitBaseRouter(apiGroup)   // base 通用
	userRouter.InitUserRouter(apiGroup)   // user 用户路由
	smRouter.InitSmRouter(apiGroup)       // 短信验证码
	taskRouter.InitDBLinkRouter(apiGroup) // 数据库连接
	taskRouter.InitTaskRouter(apiGroup)   // 数据比对任务
	return defaultRouter
}
