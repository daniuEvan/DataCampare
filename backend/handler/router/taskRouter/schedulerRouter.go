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
		schedulerRouter.POST("/", schedulerApi.AddScheduler)
		schedulerRouter.PUT("/", schedulerApi.UpdateScheduler)
		schedulerRouter.DELETE("/:scheduler_id/", schedulerApi.DeleteScheduler)
	}
}
