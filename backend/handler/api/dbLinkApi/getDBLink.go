/**
 * @date: 2022/3/4
 * @desc: ...
 */

package dbLinkApi

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
// GetDBLinkInfo
// @Description: 根据id获取数据库连接信息
// @param ctx:
//
func GetDBLinkInfo(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取数据库连接信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbLinkId := ctx.Param("db_link_id")
	if dbLinkId == "" {
		response.Response(ctx, http.StatusBadRequest, 400, nil, customError.BadRequestError.Error())
		return
	}
	var dbLink taskModel.DBLink
	err = db.Where("id = ? ", dbLinkId).First(&dbLink).Error
	if errors.Is(err, gorm.ErrRecordNotFound) || dbLink.ID == 0 {
		response.Response(ctx, http.StatusNotFound, 422, nil, customError.UnprocessableEntityError.Error())
		return
	} else if err != nil {
		global.Logger.Error("获取数据库连接信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	dbLink.DBPassword = ""
	response.Success(ctx, dbLink, "请求成功")
}

//
// GetDBLinkList
// @Description: 获取全部数据库连接信息
// @param ctx:
//
func GetDBLinkList(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("获取数据库连接信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	var dbLinks []taskModel.DBLink
	err = db.Find(&dbLinks).Error
	if err != nil {
		global.Logger.Error("获取数据库连接信息", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	for i, _ := range dbLinks {
		dbLinks[i].DBPassword = ""
	}
	response.Success(ctx, dbLinks, "请求成功")
}
