/**
 * @date: 2022/3/15
 * @desc: ...
 */

package taskApi

import (
	"DataCompare/common/customError"
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/customSQL"
	"DataCompare/handler/model/taskModel"
	"DataCompare/taskEngine/dbLinkEngine"
	"DataCompare/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

func GetResultTableInfo(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	nowDate := time.Now().Format("2006-1-2")
	pageNum, err := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	startDate := ctx.DefaultQuery("startDate", nowDate)
	endDate := ctx.DefaultQuery("endDate", nowDate)
	taskId, err := strconv.Atoi(ctx.DefaultQuery("taskId", "0"))
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	if taskId == 0 {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "taskId 参数未携带")
		return
	}
	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	// 获取task info
	var taskListModel taskModel.TaskList
	err = db.Where("id = ?", taskId).First(&taskListModel).Error
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	}
	resultDBOwner := taskListModel.ResultTableOwner
	resultDBName := taskListModel.ResultTableName

	// 获取数据库连接
	dbLinkId := taskListModel.ResultDBLinkId
	var dbLinkModel taskModel.DBLink
	err = db.Where("id = ?", dbLinkId).First(&dbLinkModel).Error
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	}
	resultTableDBOption := dbLinkEngine.DataBaseOption{
		DBType:     dbLinkModel.DBType,
		DBHost:     dbLinkModel.DBHost,
		DBPort:     dbLinkModel.DBPort,
		DBName:     dbLinkModel.DBName,
		DBUsername: dbLinkModel.DBUsername,
		DBPassword: utils.DecodeBase64ToString(dbLinkModel.DBPassword),
	}
	resultDBlinker, err := dbLinkEngine.GetDBLinker(resultTableDBOption)
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	defer resultDBlinker.Close()
	dbType := resultTableDBOption.DBType
	requestTableQuerySQL := customSQL.ResultTableSQLMap[dbType]
	var countSQL string
	switch dbType {
	case "oracle":
		requestTableQuerySQL = fmt.Sprintf(
			requestTableQuerySQL,
			resultDBOwner, resultDBName,
			startDate, endDate,
			pageSize*pageNum, (pageNum-1)*pageSize,
		)
		countSQL = fmt.Sprintf("select count(*) count from %s.%s where check_time >= to_date('%s', 'yy-MM-dd') and check_time <= to_date('%s', 'yy-MM-dd')", resultDBOwner, resultDBName, startDate, endDate)
	default:
		requestTableQuerySQL = fmt.Sprintf(
			requestTableQuerySQL,
			resultDBOwner, resultDBName,
			startDate, endDate,
			pageSize, (pageNum-1)*pageSize,
		)
		countSQL = fmt.Sprintf("select count(*) count from %s.%s where check_time >= '%s' and check_time <= '%s'", resultDBOwner, resultDBName, startDate, endDate)
	}
	fmt.Println(requestTableQuerySQL)
	tempQueryRes, err := resultDBlinker.Query(requestTableQuerySQL)
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	queryRes, err := utils.ParseQueryResult(tempQueryRes)
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	tempCountQueryRes, err := resultDBlinker.Query(countSQL)
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	countRes, err := utils.ParseQueryResult(tempCountQueryRes)
	if err != nil {
		global.Logger.Error("查询结果表", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	for _, item := range queryRes {
		item["check_time"] = item["check_time"][0:10]
	}
	dataResponse := map[string]interface{}{
		"dataArray": queryRes,
		"count":     countRes[0]["count"],
	}
	response.Success(ctx, dataResponse, "请求成功")
}
