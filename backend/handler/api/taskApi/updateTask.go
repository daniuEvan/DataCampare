/**
 * @date: 2022/3/4
 * @desc: ...
 */

package taskApi

import (
	"DataCompare/common/customError"
	"DataCompare/common/response"
	"DataCompare/common/validatorErrorHandler"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/forms/taskForm"
	"DataCompare/handler/model/taskModel"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

//
// UpdateTask
// @Description: 更新任务信息
// @param ctx:
//
func UpdateTask(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("更新任务信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	// 表单验证
	dbTaskForm := taskForm.TaskForm{}
	if err := ctx.ShouldBindJSON(&dbTaskForm); err != nil {
		global.Logger.Error(err.Error())
		validatorErrorHandler.ValidatorErrorHandler(ctx, err)
		return
	}
	dbTaskId := dbTaskForm.ID
	if dbTaskId == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	var dbTaskModel taskModel.TaskList
	db.Where("id = ?", dbTaskId).First(&dbTaskModel)

	if errors.Is(err, gorm.ErrRecordNotFound) || dbTaskModel.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("更新任务信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbTaskModel.TaskName = dbTaskForm.TaskName
	dbTaskModel.ConfigDBLinkId = dbTaskForm.ConfigDBLinkId
	dbTaskModel.SourceDBLinkId = dbTaskForm.SourceDBLinkId
	dbTaskModel.TargetDBLinkId = dbTaskForm.TargetDBLinkId
	dbTaskModel.ResultDBLinkId = dbTaskForm.ResultDBLinkId
	dbTaskModel.ResultTableOwner = dbTaskForm.ResultTableOwner
	dbTaskModel.ResultTableName = dbTaskForm.ResultTableName
	dbTaskModel.ConfigTableOwner = dbTaskForm.ConfigTableOwner
	dbTaskModel.ConfigTableName = dbTaskForm.ConfigTableName
	dbTaskModel.Desc = dbTaskForm.Desc
	err = db.Save(&dbTaskModel).Error
	if err != nil {
		global.Logger.Error("更新任务信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, dbTaskModel, "更新成功")
}
