/**
 * @date: 2022/2/21
 * @desc: ...
 */

package utils

import (
	"DataCompare/common/response"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CustomClaims struct {
	ID            uint
	Username      string
	Mobile        string
	EffectiveTime int // 有效时间s
	jwt.StandardClaims
}

//
// GetCurrentUserID
// @Description: 获取当前登录用户id
// @param ctx:
// @return userId:
// @return ok:
//
func GetCurrentUserID(ctx *gin.Context) (userId uint, ok bool) {
	userIdInterface, exists := ctx.Get("userId")
	if !exists {
		return 0, false
	}
	userId, ok = userIdInterface.(uint)
	if !ok {
		return 0, false
	}
	return
}

//
// GetCurrentUserInfo
// @Description: 获取当前用户信息
// @param ctx:
// @return userId:
// @return ok:
//
func GetCurrentUserInfo(ctx *gin.Context) (userInfo CustomClaims, ok bool) {
	userInfoInterface, exists := ctx.Get("claims")
	if !exists {
		return CustomClaims{}, false
	}
	userInfoJson, err := json.Marshal(userInfoInterface)
	if err != nil {
		response.Failed(ctx, gin.H{}, "获取用户信息失败")
		return
	}
	err = json.Unmarshal(userInfoJson, &userInfo)
	if err != nil {
		response.Failed(ctx, gin.H{}, "获取用户信息失败")
		return
	}
	return userInfo, true
}
