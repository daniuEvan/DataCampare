/**
 * @date: 2022/3/4
 * @desc:
 */

package utils

import (
	"encoding/base64"
	"reflect"
	"unsafe"
)

const BASE64Table = "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

//
// EncodeStringToBase64
// @Description: 加密
// @param str:
// @return base64Str:  密文字符串
// @return err:
//
func EncodeStringToBase64(str string) (base64Str string) {
	content := *(*[]byte)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&str))))
	coder := base64.NewEncoding(BASE64Table)
	return coder.EncodeToString(content)
}

//
// DecodeBase64ToString
// @Description: 解密
// @param cipherStr:
// @return resStr:  明文字符串
// @return err:
//
func DecodeBase64ToString(base64Str string) (resStr string) {
	coder := base64.NewEncoding(BASE64Table)
	result, _ := coder.DecodeString(base64Str)
	resStr = *(*string)(unsafe.Pointer(&result))
	return resStr
}
