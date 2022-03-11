/**
 * @date: 2022/3/6
 * @desc: ...
 */

package taskRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/taskApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitTaskRouter(routerGroup *gin.RouterGroup) {
	taskRouter := routerGroup.Group("task").Use(middleware.GinLogger(global.Logger))
	{
		taskRouter.GET("/:task_id/", taskApi.GetTaskInfo)
		taskRouter.GET("/list/", taskApi.GetTaskList)
		taskRouter.POST("/", taskApi.AddTask)
		taskRouter.PUT("/", taskApi.UpdateTask, middleware.ReInitCron())
		taskRouter.DELETE("/:task_id/", taskApi.DeleteTask, middleware.ReInitCron())
	}
}
