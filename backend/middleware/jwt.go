package middleware

import (
	"DataCompare/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CustomClaims struct {
	ID            uint
	Username      string
	Mobile        string
	EffectiveTime int // 有效时间s
	jwt.StandardClaims
}

// JWTAuth 校验
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取头部信息 x-token 登录时回返回token信息
		headTokenKey := global.ServerConfig.AuthInfo.JWTInfo.TokenKey
		token := ctx.Request.Header.Get(headTokenKey)
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "请登录",
			})
			ctx.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				if err == TokenExpired {
					ctx.JSON(http.StatusUnauthorized, map[string]string{
						"msg": "授权已过期",
					})

					ctx.Abort()
					return
				}
			}
			ctx.JSON(http.StatusUnauthorized, map[string]string{
				"msg": "未登录",
			})
			ctx.Abort()
			return
		}
		// 如果 `( 过期时间点 - 当前时间点) / < 1/3 * 有效时间`  重新发放token
		effectiveTime := claims.EffectiveTime
		expiresTime := claims.ExpiresAt
		nowUnix := time.Now().Unix()
		if expiresTime-nowUnix < int64(effectiveTime/3) {
			j2 := NewJWT()
			newToken, err := j2.RefreshToken(token)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, map[string]string{
					"msg": "自动更新token失败",
				})
				ctx.Abort()
				return
			}
			ctx.Header(headTokenKey, newToken)
			claims, _ = j.ParseToken(newToken)
		}
		// 更新ctx中的信息
		ctx.Set("claims", claims)
		ctx.Set("userId", claims.ID)
		ctx.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New(" Token is expired ")
	TokenNotValidYet = errors.New(" Token not active yet ")
	TokenMalformed   = errors.New(" That's not even a token ")
	TokenInvalid     = errors.New(" Couldn't handle this token: ")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.AuthInfo.JWTInfo.SigningKey), // 可以设置过期时间
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.SigningKey, nil
		},
	)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}

}

// RefreshToken 更新token过期时间 ( 获取新的token )
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	//jwt.TimeFunc = func() time.Time {
	//	return time.Unix(0, 0)
	//}
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomClaims{},
		func(token *jwt.Token) (i interface{}, e error) {
			return j.SigningKey, nil
		},
	)
	fmt.Println(time.Now().Format(time.StampNano))
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		effectiveTime := time.Duration(global.ServerConfig.AuthInfo.JWTInfo.EffectiveTime) * time.Second
		claims.StandardClaims.ExpiresAt = time.Now().Unix() + int64(effectiveTime)
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
