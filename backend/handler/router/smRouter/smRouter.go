/**
 * @date: 2022/2/18
 * @desc: ...
 */

package smRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/smApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitSmRouter(routerGroup *gin.RouterGroup) {
	sMRouter := routerGroup.Group("sm").Use(middleware.GinLogger(global.Logger))
	{
		sMRouter.GET("/sm_code/:mobile", smApi.GetSmCode)
	}
}
