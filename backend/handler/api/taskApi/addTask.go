/**
 * @date: 2022/3/4
 * @desc: ...
 */

package taskApi

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
// AddTask
// @Description: 新增任务
// @param ctx:
//
func AddTask(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("新增任务", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbTaskForm := taskForm.TaskForm{}
	if err := ctx.ShouldBindJSON(&dbTaskForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	var dbTaskModel taskModel.TaskList
	dbTaskModel.TaskName = dbTaskForm.TaskName
	dbTaskModel.SourceDBLinkId = dbTaskForm.SourceDBLinkId
	dbTaskModel.TargetDBLinkId = dbTaskForm.TargetDBLinkId
	dbTaskModel.ResultTableOwner = dbTaskForm.ResultTableOwner
	dbTaskModel.ResultTableName = dbTaskForm.ResultTableName
	dbTaskModel.ConfigTableOwner = dbTaskForm.ConfigTableOwner
	dbTaskModel.ConfigTableName = dbTaskForm.ConfigTableName
	dbTaskModel.TaskConcurrent = dbTaskForm.TaskConcurrent
	dbTaskModel.Desc = dbTaskForm.Desc
	err = db.Create(&dbTaskModel).Error
	if err != nil {
		global.Logger.Error("新增任务", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, dbTaskModel, "创建任务成功")

}
