/**
 * @date: 2022/3/4
 * @desc: ...
 */

package schedulerApi

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
	"strconv"
)

//
// UpdateScheduler
// @Description: 更新调度信息
// @param ctx:
//
func UpdateScheduler(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("更新调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	// 表单验证
	dbSchedulerForm := taskForm.SchedulerForm{}
	if err := ctx.ShouldBindJSON(&dbSchedulerForm); err != nil {
		global.Logger.Error(err.Error())
		validatorErrorHandler.ValidatorErrorHandler(ctx, err)
		return
	}
	dbSchedulerId := dbSchedulerForm.ID
	if dbSchedulerId == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	var dbSchedulerModel taskModel.SchedulerList
	db.Where("id = ?", dbSchedulerId).First(&dbSchedulerModel)

	if errors.Is(err, gorm.ErrRecordNotFound) || dbSchedulerModel.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("更新调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbSchedulerModel.SchedulerName = dbSchedulerForm.SchedulerName
	dbSchedulerModel.TaskId = dbSchedulerForm.TaskId
	dbSchedulerModel.TaskSchedule = dbSchedulerForm.TaskSchedule
	dbSchedulerModel.ConfigTableQuerySQL = dbSchedulerForm.ConfigTableQuerySQL
	dbSchedulerModel.ResultTableInsertSQL = dbSchedulerForm.ResultTableInsertSQL
	dbSchedulerModel.ResultTableInitSQL = dbSchedulerForm.ResultTableInitSQL
	dbSchedulerModel.SourceTableQuerySQL = dbSchedulerForm.SourceTableQuerySQL
	dbSchedulerModel.TargetTableQuerySQL = dbSchedulerForm.TargetTableQuerySQL
	dbSchedulerModel.SchedulerStatus = dbSchedulerForm.SchedulerStatus
	dbSchedulerModel.TaskConcurrent = dbSchedulerForm.TaskConcurrent
	dbSchedulerModel.Desc = dbSchedulerForm.Desc
	err = db.Save(&dbSchedulerModel).Error
	if err != nil {
		global.Logger.Error("更新调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, dbSchedulerModel, "更新调度成功")
}

//
// EnableScheduler
// @Description: 更新调度信息
// @param ctx:
//
func EnableScheduler(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("调度开关", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	schedulerId, err := strconv.Atoi(ctx.Param("scheduler_id"))
	if err != nil {
		global.Logger.Error("解析调度id", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	schedulerStatus, err := strconv.ParseBool(ctx.Param("scheduler_status"))
	if err != nil {
		global.Logger.Error("解析调度状态", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	var dbSchedulerModel taskModel.SchedulerList
	err = db.Model(&dbSchedulerModel).Where("id = ?", schedulerId).Update("scheduler_status", schedulerStatus).Error
	if err != nil {
		global.Logger.Error("更新调度状态", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, nil, "更新调度状态成功")
}
