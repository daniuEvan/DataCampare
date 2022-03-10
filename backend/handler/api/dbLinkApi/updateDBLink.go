/**
 * @date: 2022/3/4
 * @desc: ...
 */

package dbLinkApi

import (
	"DataCompare/common/customError"
	"DataCompare/common/response"
	"DataCompare/common/validatorErrorHandler"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/forms/taskForm"
	"DataCompare/handler/model/taskModel"
	"DataCompare/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func UpdateDBLink(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("更新数据库连接", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	// 表单验证
	dbLinkForm := taskForm.DBLinkForm{}
	if err := ctx.ShouldBindJSON(&dbLinkForm); err != nil {
		global.Logger.Error(err.Error())
		validatorErrorHandler.ValidatorErrorHandler(ctx, err)
		return
	}
	dbLinkId := dbLinkForm.ID
	if dbLinkId == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	var dbLinkModel taskModel.DBLink
	db.Where("id = ?", dbLinkId).First(&dbLinkModel)
	if dbLinkModel.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, customError.UnprocessableEntityError.Error())
		return
	}
	dbPassword := dbLinkForm.DBPassword
	cipherPassword := utils.EncodeStringToBase64(dbPassword)
	dbLinkModel.LinkName = dbLinkForm.LinkName
	dbLinkModel.DBType = dbLinkForm.DBType
	dbLinkModel.DBHost = dbLinkForm.DBHost
	dbLinkModel.DBPort = dbLinkForm.DBPort
	dbLinkModel.DBName = dbLinkForm.DBName
	dbLinkModel.DBUsername = dbLinkForm.DBUsername
	dbLinkModel.DBPassword = cipherPassword
	err = db.Save(&dbLinkModel).Error
	if err != nil {
		global.Logger.Error("更新数据库连接", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbLinkModel.DBPassword = ""
	response.Success(ctx, dbLinkModel, "更新成功")
}
