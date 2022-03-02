/*
 * @date: 2021/12/16
 * @desc: ...
 */

package userModel

import (
	"gorm.io/gorm"
)

//
// User
// @Description: 普通用户表
//
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(25);not null"`
	Mobile   string `gorm:"type:varchar(11);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
	Addr     string `gorm:"type:varchar(200)"`
}
