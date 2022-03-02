/*
 * @date: 2021/12/15
 * @desc: ...
 */

package userApi

import (
	"DataCompare/common/response"
	"DataCompare/database"
	"DataCompare/global"
	"DataCompare/handler/forms/userForm"
	"DataCompare/handler/model/userModel"
	"DataCompare/middleware"
	"DataCompare/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

//
// Register
// @Description: 用户注册
// @param ctx:
//
func Register(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("用户注册", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	registerLoginForm := userForm.RegisterForm{}
	if err := ctx.ShouldBindJSON(&registerLoginForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	mobile := registerLoginForm.Mobile
	username := registerLoginForm.Username
	if username == "admin" {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "admin不可用,请更换用户名")
		return
	}
	password := registerLoginForm.Password
	smCode := registerLoginForm.VerifyCode
	addr := registerLoginForm.Addr
	// 校验验证码
	smService := utils.NewSmService()
	ok, err := smService.VerifySmCode(mobile, smCode)
	if err != nil {
		global.Logger.Error("用户注册", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "验证码校验服务异常")
		return
	}
	if !ok {
		response.Response(ctx, http.StatusInternalServerError, 400, nil, "验证码错误")
		return
	}

	// 验证用户是否存在
	if IsMobilExist(db, mobile) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经存在")
		return
	}
	// 用户不存在创建用户
	hashPassword, err := utils.GetHashPwd(password)
	if err != nil {
		global.Logger.Error("用户注册", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "服务器异常")
		return
	}
	user := userModel.User{
		Mobile:   mobile,
		Username: username,
		Addr:     addr,
		Password: string(hashPassword),
	}
	db.Create(&user)
	//	发放token
	effectiveTime := global.ServerConfig.AuthInfo.JWTInfo.EffectiveTime
	token, err := buildToken(user)
	if err != nil {
		global.Logger.Error("用户注册", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "服务器异常")
		return
	}
	userInfo := gin.H{
		"id":             user.ID,
		"username":       user.Username,
		"mobile":         user.Mobile,
		"addr":           user.Addr,
		"token":          token,
		"expired_at":     time.Now().Unix() + int64(effectiveTime),
		"effective_time": effectiveTime,
	}
	ctx.Set("userInfo", userInfo)
	response.Success(ctx, userInfo, "注册成功")
}

//
// PasswordLogin
// @Description: 密码登录
// @param ctx:
//
func PasswordLogin(ctx *gin.Context) {
	db, err := database.GetMysqlDB(ctx)
	if err != nil {
		global.Logger.Error("密码登录", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "服务异常")
		return
	}
	// 表单验证
	passwordLoginForm := userForm.LoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}

	mobile := passwordLoginForm.Mobile
	//username := passwordLoginForm.Username
	password := passwordLoginForm.Password
	var user userModel.User
	db.Where("mobile = ?", mobile).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//   验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		global.Logger.Error("密码登录", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 401, nil, "密码错误")
		return
	}
	effectiveTime := global.ServerConfig.AuthInfo.JWTInfo.EffectiveTime
	token, err := buildToken(user)
	if err != nil {
		global.Logger.Error("密码登录", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusUnprocessableEntity, 500, nil, "服务器异常")
		return
	}
	userInfo := gin.H{
		"id":             user.ID,
		"username":       user.Username,
		"mobile":         user.Mobile,
		"addr":           user.Addr,
		"token":          token,
		"expired_at":     time.Now().Unix() + int64(effectiveTime),
		"effective_time": effectiveTime,
	}
	ctx.Set("userInfo", userInfo)
	response.Success(ctx, userInfo, "登录成功")
}

//
// LdapLogin
// @Description:  ldap 登录
// @param ctx:
//
func LdapLogin(ctx *gin.Context) {
	// 表单验证
	passwordLoginForm := userForm.LoginForm{}
	if err := ctx.ShouldBindJSON(&passwordLoginForm); err != nil {
		global.Logger.Error(err.Error())
		utils.ValidatorErrorHandler(ctx, err)
		return
	}
	// user信息
	_ = passwordLoginForm.Mobile
	username := passwordLoginForm.Mobile
	password := passwordLoginForm.Password
	// ldap base msg
	ldapInfo := global.ServerConfig.AuthInfo.LADPInfo
	//ldap host msg
	ldapHost := ldapInfo.LdapHost
	ldapPort := ldapInfo.LdapPort
	// 拨号5s超时
	ldap.DefaultTimeout = 5 * time.Second
	ldapHandler, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapHost, ldapPort))
	if err != nil {
		global.Logger.Error("LDAP服务连接失败", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, "LDAP服务连接失败")
		ctx.Abort()
		return
	}
	// ldap bind msg
	bindUsername := ldapInfo.BindDN
	bindPassword := ldapInfo.BindPassword
	_, err = ldapHandler.SimpleBind(&ldap.SimpleBindRequest{
		Username: bindUsername,
		Password: bindPassword,
	})
	if err != nil {
		global.Logger.Error("LDAP服务绑定检查用户失败", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, "LDAP服务绑定检查用户失败")
		ctx.Abort()
		return
	}
	ldapBaseDN := ldapInfo.BaseDN
	ldapSearchProperty := ldapInfo.SearchProperty
	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		ldapBaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=*)(%s=%s))", ldapSearchProperty, username),
		[]string{"dn", "sn", "gidNumber", "uidNumber"},
		nil,
	)
	sr, err := ldapHandler.Search(searchRequest)
	if err != nil {
		global.Logger.Error("LDAP 查找用户失败 ", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, "LDAP 查找用户失败 ")
		ctx.Abort()
		return
	}
	if len(sr.Entries) != 1 {
		global.Logger.Error("LDAP 查找用户失败 ", zap.String("msg", fmt.Sprintf("查找到了%d个用户", len(sr.Entries))))
		response.Response(ctx, http.StatusBadRequest, 400, nil, "LDAP 查找用户失败 ")
		ctx.Abort()
		return
	}
	userFullDN := sr.Entries[0].DN // 完整的dn
	// Bind as the user to verify their password
	err = ldapHandler.Bind(userFullDN, password)
	if err != nil {
		global.Logger.Error("用户密码错误", zap.String("msg", err.Error()))
		response.Response(ctx, http.StatusBadRequest, 400, nil, "用户密码错误")
		ctx.Abort()
		return
	}
	var user userModel.User
	user.ID = 0 // todo 修改为自己的参数
	user.Username = username
	user.Mobile = username
	// token有效时间s
	effectiveTime := global.ServerConfig.AuthInfo.JWTInfo.EffectiveTime
	token, _ := buildToken(user)
	ctx.Set("userId", user.ID)
	response.Success(ctx, gin.H{
		"id":             user.ID,
		"username":       user.Username,
		"mobile":         user.Mobile,
		"token":          token,
		"expired_at":     time.Now().Unix() + int64(effectiveTime),
		"effective_time": effectiveTime,
	}, "登录成功")
}

//
// IsMobilExist
// @Description: 验证用户是否存在
// @param db:
// @param mobile:
// @return bool:
//
func IsMobilExist(db *gorm.DB, mobile string) bool {
	var user userModel.User
	db.Where("mobile=?", mobile).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

//
// buildToken
// @Description: 发放token
// @param user:
// @return token:
// @return err:
//
func buildToken(user userModel.User) (token string, err error) {
	//	发放token
	effectiveTime := global.ServerConfig.AuthInfo.JWTInfo.EffectiveTime
	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		ID:            uint(user.ID),
		Username:      user.Username,
		Mobile:        user.Mobile,
		EffectiveTime: effectiveTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                        // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(effectiveTime), // 签名过期(s单位)
			Issuer:    "ginCli",
		},
	}
	token, err = j.CreateToken(claims)
	if err != nil {
		global.Logger.Error("生成token 失败", zap.String("msg", err.Error()))
		return "", err
	}
	return token, nil
}
