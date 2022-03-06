/*
 * @date: 2021/12/15
 * @desc: ...
 */

package userRouter

import (
	"DataCompare/global"
	"DataCompare/handler/api/userApi"
	"DataCompare/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("user").Use(middleware.GinLogger(global.Logger))
	{
		userRouter.POST("/login/", userApi.PasswordLogin)
		userRouter.POST("/login_ldap/", userApi.LdapLogin)
		userRouter.POST("/register/", userApi.Register)
	}
	editUserRouter := routerGroup.Group("user_info").Use(middleware.GinLogger(global.Logger)).Use(middleware.JWTAuth())
	{
		editUserRouter.GET("/user_list/", middleware.AdminFilter(), userApi.GetUserList)
		editUserRouter.GET("/:mobile/", middleware.AdminFilter(), userApi.GetUserInfo)
		editUserRouter.POST("/:mobile/", userApi.EditUserInfo)
		editUserRouter.POST("/change_pwd/", userApi.ChangePwd)
	}
}
