/**
 * @date: 2022/2/18
 * @desc: ...
 */

package userApi

import (
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/forms/userForm"
	"DataCompare/handler/model/userModel"
	"DataCompare/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//
// EditUserInfo
// @Description: 编辑用户信息
// @param ctx:
//
func EditUserInfo(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("编辑用户信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	// 表单验证
	userInfoForm := userForm.UserInfoForm{}
	if err := ctx.ShouldBindJSON(&userInfoForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	userId, ok := utils.GetCurrentUserID(ctx)
	if !ok {
		response.Response(ctx, http.StatusUnauthorized, 401, nil, "未登录")
		return
	}
	var user userModel.User
	db.Where("id = ?", userId).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	user.Username = userInfoForm.Username
	user.Addr = userInfoForm.Addr
	db.Save(&user)
	response.Success(ctx, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"mobile":   user.Mobile,
		"addr":     user.Addr,
	}, "更新成功")

}

//
// GetUserInfo
// @Description: 获取用户信息
// @param ctx:
//
func GetUserInfo(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取用户信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	userMobile := ctx.Param("mobile")
	var user userModel.User
	db.Where("mobile=?", userMobile).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	response.Success(ctx, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"mobile":   user.Mobile,
		"addr":     user.Addr,
	}, "")

}

//
// ChangePwd
// @Description: 修改密码
// @param ctx:
//
func ChangePwd(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("修改密码", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	// 表单验证
	userPwdInfoForm := userForm.UserPwdInfoForm{}
	if err := ctx.ShouldBindJSON(&userPwdInfoForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	mobile := userPwdInfoForm.Mobile
	userId, ok := utils.GetCurrentUserID(ctx)
	if !ok {
		response.Response(ctx, http.StatusUnauthorized, 401, nil, "未登录或者登录异常")
		return
	}
	smCode := userPwdInfoForm.VerifyCode
	newPwd := userPwdInfoForm.Password
	var user userModel.User
	db.Where("id = ? and mobile = ?", userId, mobile).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	// 校验验证码
	smService := utils.NewSmService()
	ok, err = smService.VerifySmCode(mobile, smCode)
	if err != nil {
		global.Logger.Error("修改密码", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "验证码校验服务异常")
		return
	}
	if !ok {
		response.Response(ctx, http.StatusInternalServerError, 400, nil, "验证码错误")
		return
	}
	hashPassword, err := utils.GetHashPwd(newPwd)
	if err != nil {
		global.Logger.Error("修改密码", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "密码加密错误")
		return
	}
	user.Password = hashPassword
	db.Save(&user)
	response.Success(ctx, gin.H{}, "密码修改成功")
}

//
// GetUserList
// @Description: 获取用户列表
// @param ctx:
//
func GetUserList(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取用户列表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	var userList []userModel.User
	result := db.Find(&userList)
	if result.Error != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "获取用户列表异常")
		return
	}

	resList := make([]interface{}, 0)
	for _, user := range userList {
		userMap := utils.StructToMapViaReflect(&user)
		delete(userMap, "Password")
		resList = append(resList, userMap)
	}
	resResponse := map[string]interface{}{
		"dataList": resList,
		"count":    result.RowsAffected,
	}

	response.Response(ctx, http.StatusOK, 200, resResponse, "res")
}
