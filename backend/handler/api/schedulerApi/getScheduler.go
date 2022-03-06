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
	var schedulerModel taskModel.SchedulerList
	err = db.Where("id = ? ", schedulerId).First(&schedulerModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || schedulerModel.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("获取调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, schedulerModel, "获取调度信息成功")
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
	var dbSchedulerList []taskModel.SchedulerList
	err = db.Find(&dbSchedulerList).Error
	if err != nil {
		global.Logger.Error("获取调度信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, dbSchedulerList, "获取调度信息列表成功")
}
