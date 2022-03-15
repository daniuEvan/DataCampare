/**
 * @date: 2022/3/15
 * @desc: 解析SQL查询结果
 */

package utils

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"strings"
)

func ParseQueryResult(queryRes map[string]interface{}) (parseRes []map[string]string, err error) {
	byteSliceRes, err := json.Marshal(queryRes)
	if err != nil {
		return nil, err
	}
	gResult := gjson.ParseBytes(byteSliceRes)
	infoKeys := gResult.Get("columns").Array()
	infoValues := gResult.Get("values").Array()
	for _, values := range infoValues {
		tempMap := make(map[string]string)
		for j, value := range values.Array() {
			tempMap[strings.ToLower(infoKeys[j].String())] = value.String()
		}
		parseRes = append(parseRes, tempMap)
	}
	return
}
