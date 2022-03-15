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
	"DataCompare/handler/model/taskModel"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

//
// GetSchedulerInfo
// @Description: 获取调度信息
// @param ctx:
//
func GetSchedulerInfo(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	schedulerId := ctx.Param("scheduler_id")
	if schedulerId == "" {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	type SchedulerRes struct {
		taskModel.SchedulerList
		TaskName string
	}
	var schedulerRes SchedulerRes
	err = db.Model(&taskModel.SchedulerList{}).Select("compare_scheduler_list.*", "t.task_name").Joins(
		"left join compare_task_list t on  compare_scheduler_list.task_id =  t.id ",
	).Where("compare_scheduler_list.id = ? ", schedulerId).Scan(&schedulerRes).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || schedulerRes.ID == 0 {
		global.Logger.Error("获取调度信息", zap.String("msg", customError.ResourceNotFountError.Error()))
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("获取调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, schedulerRes, "获取调度信息成功")
}

//
// GetSchedulerList
// @Description: 获取调度信息列表
// @param ctx:
//
func GetSchedulerList(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	type SchedulerRes struct {
		taskModel.SchedulerList
		TaskName string
	}
	var schedulerRes []SchedulerRes
	err = db.Model(&taskModel.SchedulerList{}).Select("compare_scheduler_list.*", "t.task_name").Joins(
		"left join compare_task_list t on  compare_scheduler_list.task_id =  t.id ",
	).Scan(&schedulerRes).Error
	if err != nil {
		global.Logger.Error("获取调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, schedulerRes, "获取调度信息列表成功")
}
