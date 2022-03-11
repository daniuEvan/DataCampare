/**
 * @date: 2022/3/6
 * @desc: ...
 */

package taskRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/schedulerApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitSchedulerRouter(routerGroup *gin.RouterGroup) {
	schedulerRouter := routerGroup.Group("scheduler").Use(middleware.GinLogger(global.Logger))
	{
		schedulerRouter.GET("/:scheduler_id/", schedulerApi.GetSchedulerInfo)
		schedulerRouter.GET("/list/", schedulerApi.GetSchedulerList)
		schedulerRouter.GET("/watch_scheduler/", schedulerApi.WatchScheduler) // 监控scheduler
		schedulerRouter.POST("/", schedulerApi.AddScheduler, middleware.ReInitCron())
		schedulerRouter.PUT("/", schedulerApi.UpdateScheduler, middleware.ReInitCron())
		schedulerRouter.DELETE("/:scheduler_id/", schedulerApi.DeleteScheduler, middleware.ReInitCron())
	}
}
