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

func DeleteDBLink(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("删除数据库连接", zap.String("msg", err.Error()))
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
		global.Logger.Error("删除数据库连接", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	err = db.Delete(&dbLink).Error
	if err != nil {
		global.Logger.Error("删除数据库连接", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, customError.InternalServerError.Error())
		return
	}
	response.Success(ctx, nil, "删除成功")
}
