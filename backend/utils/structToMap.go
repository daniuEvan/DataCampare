/**
 * @date: 2022/2/19
 * @desc: ...
 */

package utils

import (
	"fmt"
	"reflect"
)

//
// StructToMapViaReflect
// @Description: StructToMap
// @param st: st的类型必须是一个 struct
// @return res:
//
func StructToMapViaReflect(st interface{}) map[string]interface{} {
	resMap := make(map[string]interface{})
	elem := reflect.ValueOf(st).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		resMap[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	fmt.Println(resMap)
	return resMap
}
