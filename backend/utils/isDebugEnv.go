/*
 * @date: 2021/12/15
 * @desc: ...
 */

package utils

import (
	"github.com/spf13/viper"
)

// IsDebugEnv 获取 DEBUG_ENV 环境变量
func IsDebugEnv() bool {
	debug := false
	viper.AutomaticEnv()
	env := viper.GetString("DEBUG_ENV")
	if env == "true" {
		debug = true
	}
	return debug
}
