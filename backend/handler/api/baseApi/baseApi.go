/*
 * @date: 2021/12/15
 * @desc: ...
 */

package baseApi

import (
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"net/http"
)

func Ping(ctx *gin.Context) {
	conn, err := database.GetRedisConn()

	if err != nil {
		global.Logger.Error("获取redis连接异常", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "获取redis连接异常")
	}
	defer conn.Close()
	res, err := redis.Bool(conn.Do("sAdd", "test", 123))
	ctx.JSON(http.StatusOK, gin.H{
		"msg": res,
	})
}
