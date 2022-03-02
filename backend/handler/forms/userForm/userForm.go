/*
 * @date: 2021/12/15
 * @desc: ...
 */

package userForm

//
// LoginForm
// @Description: 登录form
//
type LoginForm struct {
	Mobile   string `forms:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `forms:"password" json:"password" binding:"required,min=3,max=20"`
}

//
// RegisterForm
// @Description: 注册form
//
type RegisterForm struct {
	Mobile     string `forms:"mobile" json:"mobile" binding:"required,mobile"`
	Username   string `forms:"username" json:"username" binding:"required,min=1,max=20"`
	Password   string `forms:"password" json:"password" binding:"required,min=3,max=20"`
	Addr       string `forms:"addr" json:"addr" binding:"required"`
	VerifyCode string `forms:"code" json:"code" binding:"required,min=4,max=6"`
}

//
// UserInfoForm
// @Description: 修改用户信息
//
type UserInfoForm struct {
	Id       int    `forms:"id" json:"id" binding:"required"`
	Username string `forms:"username" json:"username" binding:"required"`
	Addr     string `forms:"addr" json:"addr" binding:"required"`
}

//
// UserPwdInfoForm
// @Description: 修改密码
//
type UserPwdInfoForm struct {
	Id         int    `forms:"id" json:"id" binding:"required"`
	Mobile     string `forms:"mobile" json:"mobile" binding:"required,mobile"`
	Password   string `forms:"password" json:"password" binding:"required,min=3,max=20"`
	VerifyCode string `forms:"code" json:"code" binding:"required,min=4,max=6"`
}
