/**
 * @date: 2022/3/15
 * @desc: ...
 */

package taskRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/taskApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitResultRouter(routerGroup *gin.RouterGroup) {
	resultRouter := routerGroup.Group("result").Use(middleware.GinLogger(global.Logger))
	{
		resultRouter.GET("/result_table/", taskApi.GetResultTableInfo)
	}
}
