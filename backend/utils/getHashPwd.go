/**
 * @date: 2022/2/19
 * @desc: ...
 */

package utils

import "golang.org/x/crypto/bcrypt"

//
// GetHashPwd
// @Description: 获取加密的pwd
// @param pwd:
// @return hashPwd:
//
func GetHashPwd(pwd string) (hashPwd string, err error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost) // 密码加密
	if err != nil {
		return "", err
	}
	hashPwd = string(hashPassword)
	return hashPwd, nil
}
