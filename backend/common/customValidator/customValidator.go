/*
 * @date: 2021/12/15
 * @desc: ...
 */

package customValidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidateMobile 手机号码校验器
func ValidateMobile(f1 validator.FieldLevel) bool {
	mobile := f1.Field().String()
	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
