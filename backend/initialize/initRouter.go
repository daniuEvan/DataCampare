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
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	var defaultRouter *gin.Engine
	if utils.IsDebugEnv() {
		defaultRouter = gin.Default()
		pprof.Register(defaultRouter)
	} else {
		gin.SetMode(gin.ReleaseMode)
		defaultRouter = gin.New()
	}
	defaultRouter.Use(middleware.Cors()) // 跨域
	apiGroup := defaultRouter.Group("api/v1")
	baseRouter.InitBaseRouter(apiGroup)      // base 通用
	userRouter.InitUserRouter(apiGroup)      // user 用户路由
	smRouter.InitSmRouter(apiGroup)          // 短信验证码
	taskRouter.InitDBLinkRouter(apiGroup)    // 数据比对 数据库连接
	taskRouter.InitTaskRouter(apiGroup)      // 数据比对 任务
	taskRouter.InitSchedulerRouter(apiGroup) // 数据比对 调度
	taskRouter.InitResultRouter(apiGroup)    // 数据比对 调度
	return defaultRouter
}
