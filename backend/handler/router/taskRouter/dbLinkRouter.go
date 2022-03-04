/**
 * @date: 2022/3/4
 * @desc: ...
 */

package taskRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/dbLinkApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitDBLinkRouter(routerGroup *gin.RouterGroup) {
	dbLinkRouter := routerGroup.Group("db_link").Use(middleware.GinLogger(global.Logger))
	{
		dbLinkRouter.GET("/:db_link_id", dbLinkApi.GetDBLinkInfo)
		dbLinkRouter.GET("/list", dbLinkApi.GetDBLinkList)
		dbLinkRouter.POST("/", dbLinkApi.AddDBLink)
		dbLinkRouter.POST("/ping", dbLinkApi.PingDBLink)
		dbLinkRouter.PUT("/", dbLinkApi.UpdateDBLink)
		dbLinkRouter.DELETE("/:db_link_id", dbLinkApi.DeleteDBLink)
	}
}
