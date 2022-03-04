/**
 * @date: 2022/3/4
 * @desc: ...
 */

package dbLinkApi

import (
	"DataCompare/common/customError"
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/forms/taskForm"
	"DataCompare/handler/model/taskModel"
	"DataCompare/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

//
// AddDBLink
// @Description: 新增数据库连接
// @param ctx:
//
func AddDBLink(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("新增数据库连接", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbLinkForm := taskForm.DBLinkForm{}
	if err := ctx.ShouldBindJSON(&dbLinkForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	dbPassword := dbLinkForm.DBPassword
	cipherPassword := utils.EncodeStringToBase64(dbPassword)
	var dbLinkModel taskModel.DBLink
	dbLinkModel.LinkName = dbLinkForm.LinkName
	dbLinkModel.DBType = dbLinkForm.DBType
	dbLinkModel.DBHost = dbLinkForm.DBHost
	dbLinkModel.DBPort = dbLinkForm.DBPort
	dbLinkModel.DBName = dbLinkForm.DBName
	dbLinkModel.DBUsername = dbLinkForm.DBUsername
	dbLinkModel.DBPassword = cipherPassword
	err = db.Create(&dbLinkModel).Error
	if err != nil {
		global.Logger.Error("新增数据库连接", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbLinkModel.DBPassword = ""
	response.Success(ctx, dbLinkModel, "创建成功")

}
