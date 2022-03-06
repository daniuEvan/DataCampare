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
	"DataCompare/handler/model/taskModel"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

//
// GetTaskInfo
// @Description: 根据id获取任务
// @param ctx:
//
func GetTaskInfo(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取任务信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	taskId := ctx.Param("task_id")
	if taskId == "" {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	querySQl := `SELECT
					t.*,
					d1.link_name AS source_db_link_name,
					d1.db_type AS source_db_type,
					d2.link_name AS target_db_link_name ,
					d2.db_type AS target_db_type 
				FROM
					compare_task_list t
					LEFT JOIN compare_db_link d1 ON t.source_db_link_id = d1.id
					LEFT JOIN compare_db_link d2 ON t.target_db_link_id = d2.id
				where t.id = ?`

	type dbTaskInfo struct {
		taskModel.TaskList
		SourceDbLinkName string
		SourceDbType     string
		TargetDbLinkName string
		TargetDbType     string
	}
	var taskInfo dbTaskInfo
	err = db.Raw(querySQl, taskId).Scan(&taskInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || taskInfo.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("获取任务信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, taskInfo, "请求成功")
}

//
// GetTaskList
// @Description: 获取全部任务列表
// @param ctx:
//
func GetTaskList(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取全部任务列表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	querySQl := `SELECT
					t.*,
					d1.link_name AS source_db_link_name,
					d1.db_type AS source_db_type,
					d2.link_name AS target_db_link_name ,
					d2.db_type AS target_db_type 
				FROM
					compare_task_list t
					LEFT JOIN compare_db_link d1 ON t.source_db_link_id = d1.id
					LEFT JOIN compare_db_link d2 ON t.target_db_link_id = d2.id`
	type dbTaskInfo struct {
		taskModel.TaskList
		SourceDbLinkName string
		SourceDbType     string
		TargetDbLinkName string
		TargetDbType     string
	}
	var dbTaskInfoList []dbTaskInfo
	err = db.Raw(querySQl).Scan(&dbTaskInfoList).Error
	if err != nil {
		global.Logger.Error("获取全部任务列表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, dbTaskInfoList, "获取全部任务列表成功")
}
