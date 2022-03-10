/**
 * @date: 2022/3/4
 * @desc: 测试数据连接接口
 */

package dbLinkApi

import (
	"DataCompare/common/customError"
	"DataCompare/common/response"
	"DataCompare/common/validatorErrorHandler"
	"DataCompare/global"
	"DataCompare/handler/forms/taskForm"
	"DataCompare/taskEngine/dbLinkEngine"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func PingDBLink(ctx *gin.Context) {
	dbLinkForm := taskForm.DBLinkForm{}
	if err := ctx.ShouldBindJSON(&dbLinkForm); err != nil {
		global.Logger.Error(err.Error())
		validatorErrorHandler.ValidatorErrorHandler(ctx, err)
		return
	}
	dbOptions := dbLinkEngine.DataBaseOption{
		DBType:     dbLinkForm.DBType,
		DBHost:     dbLinkForm.DBHost,
		DBPort:     dbLinkForm.DBPort,
		DBName:     dbLinkForm.DBName,
		DBUsername: dbLinkForm.DBUsername,
		DBPassword: dbLinkForm.DBPassword,
	}
	var dbLinker dbLinkEngine.DBLinker
	var err error
	switch dbLinkForm.DBType {
	case "vertica":
		dbLinker, err = dbLinkEngine.NewVerticaLink(dbOptions)
	case "oracle":
		dbLinker, err = dbLinkEngine.NewOracleLink(dbOptions)
	case "mysql":
		dbLinker, err = dbLinkEngine.NewMysqlLink(dbOptions)
	case "postgres":
		dbLinker, err = dbLinkEngine.NewPostgresLink(dbOptions)
	default:
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	if err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.DatabaseConnectError.Error())
		return
	}
	defer func(dbLinker dbLinkEngine.DBLinker) {
		err := dbLinker.Close()
		if err != nil {
			global.Logger.Error("数据库关闭异常", zap.String("msg", err.Error()))
		}
	}(dbLinker)
	response.Success(ctx, nil, "连接成功")
}
