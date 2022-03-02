/*
 * @date: 2021/12/15
 * @desc: ...
 */

package baseRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/baseApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitBaseRouter(routerGroup *gin.RouterGroup) {
	baseRouter := routerGroup.Group("base").Use(middleware.JWTAuth())
	{
		baseRouter.GET("/ping", middleware.AdminFilter(), middleware.GinLogger(global.Logger), baseApi.Ping)
	}
}
