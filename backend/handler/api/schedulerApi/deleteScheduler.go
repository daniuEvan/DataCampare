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
// DeleteScheduler
// @Description: 删除调度
// @param ctx:
//
func DeleteScheduler(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("删除调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbLinkId := ctx.Param("scheduler_id")
	if dbLinkId == "" {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	var dbScheduler taskModel.SchedulerList
	err = db.Where("id = ? ", dbLinkId).First(&dbScheduler).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || dbScheduler.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("删除调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	err = db.Delete(&dbScheduler).Error
	if err != nil {
		global.Logger.Error("删除调度", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, nil, "调度删除成功")
}
