/**
 * @date: 2022/3/11
 * @desc: ...
 */

package schedulerApi

import (
	"DataCompare/common/customError"
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/model/taskModel"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

//
// WatchScheduler
// @Description: 返回任务状态列表
// @param ctx:
//
func WatchScheduler(ctx *gin.Context) {
	type responseStruct struct {
		SchedulerId     uint
		SchedulerName   string
		SchedulerStatus bool
		SchedulerEnable bool
		TaskConcurrent  uint
		TaskScheduler   string // 任务表达式
		ErrorMsg        string
	}
	responseMap := map[string][]responseStruct{
		"schedulerAllItems":     make([]responseStruct, 0),
		"schedulerSuccessItems": make([]responseStruct, 0),
		"schedulerFailItems":    make([]responseStruct, 0),
		"schedulerBanItems":     make([]responseStruct, 0),
	}
	// 已启动调度
	schedulerHandler := global.SchedulerHandler
	schedulerStartStatusMap := schedulerHandler.GetAllScheduler()
	for schedulerId, schedulerStartStatus := range schedulerStartStatusMap {
		schedulerInfo := schedulerStartStatus.SchedulerInfo
		schedulerName := schedulerInfo["scheduler_name"]
		taskConcurrent, err := strconv.Atoi(schedulerInfo["task_concurrent"])
		if err != nil {
			global.Logger.Error("监控调度接口 taskConcurrent 异常", zap.String("msg", err.Error()))
			taskConcurrent = 0
		}
		taskScheduler := schedulerInfo["task_schedule"]
		cronErrorMsg := schedulerStartStatus.ErrorMsg
		schedulerInfoStruct := responseStruct{
			SchedulerId:     uint(schedulerId),
			SchedulerName:   schedulerName,
			SchedulerEnable: true,
			SchedulerStatus: schedulerStartStatus.Status,
			TaskConcurrent:  uint(taskConcurrent),
			TaskScheduler:   taskScheduler,
			ErrorMsg:        cronErrorMsg,
		}
		responseMap["schedulerAllItems"] = append(responseMap["schedulerAllItems"], schedulerInfoStruct)
		if schedulerStartStatus.Status {
			responseMap["schedulerSuccessItems"] = append(responseMap["schedulerSuccessItems"], schedulerInfoStruct)
		} else {
			responseMap["schedulerFailItems"] = append(responseMap["schedulerFailItems"], schedulerInfoStruct)
		}
	}
	// 未启动调度
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("监控调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	var schedulerModels []taskModel.SchedulerList
	// 查询所有未启用调度
	err = db.Where("scheduler_status is not true").Find(&schedulerModels).Error
	if err != nil {
		global.Logger.Error("监控调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	for _, schedulerModel := range schedulerModels {
		schedulerInfoStruct := responseStruct{
			SchedulerName:   schedulerModel.SchedulerName,
			SchedulerId:     schedulerModel.ID,
			SchedulerEnable: false,
			TaskConcurrent:  schedulerModel.TaskConcurrent,
			TaskScheduler:   schedulerModel.TaskSchedule,
		}
		responseMap["schedulerBanItems"] = append(responseMap["schedulerBanItems"], schedulerInfoStruct)
		responseMap["schedulerAllItems"] = append(responseMap["schedulerAllItems"], schedulerInfoStruct)
	}
	response.Success(ctx, responseMap, "获取调度状态信息成功")
}
