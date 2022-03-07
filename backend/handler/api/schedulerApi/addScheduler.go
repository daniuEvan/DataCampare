/**
 * @date: 2022/3/4
 * @desc: ...
 */

package schedulerApi

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
// AddScheduler
// @Description: 新增调度
// @param ctx:
//
func AddScheduler(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("新增调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbSchedulerForm := taskForm.SchedulerForm{}
	if err := ctx.ShouldBindJSON(&dbSchedulerForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	var dbSchedulerModel taskModel.SchedulerList
	dbSchedulerModel.SchedulerName = dbSchedulerForm.SchedulerName
	dbSchedulerModel.TaskId = dbSchedulerForm.TaskId
	dbSchedulerModel.TaskSchedule = dbSchedulerForm.TaskSchedule
	dbSchedulerModel.ConfigTableQuerySQL = dbSchedulerForm.ConfigTableQuerySQL
	dbSchedulerModel.ResultTableCreateSQL = dbSchedulerForm.ResultTableCreateSQL
	dbSchedulerModel.ResultTableInitSQL = dbSchedulerForm.ResultTableInitSQL
	dbSchedulerModel.SourceTableQuerySQL = dbSchedulerForm.SourceTableQuerySQL
	dbSchedulerModel.TargetTableQuerySQL = dbSchedulerForm.TargetTableQuerySQL
	dbSchedulerModel.SchedulerStatus = dbSchedulerForm.SchedulerStatus
	dbSchedulerModel.TaskConcurrent = dbSchedulerForm.TaskConcurrent
	dbSchedulerModel.Desc = dbSchedulerForm.Desc
	err = db.Create(&dbSchedulerModel).Error
	if err != nil {
		global.Logger.Error("新增调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, dbSchedulerModel, "新增调度成功")

}
