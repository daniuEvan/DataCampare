/*
 * @date: 2021/12/15
 * @desc: ...
 */

package global

import (
	"DataCompare/config"
	"DataCompare/taskEngine/engine"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
)

var (
	ServerConfig     = &config.ServerConfig{} // 全局配置
	Logger           *zap.Logger              // 全局logger
	Trans            ut.Translator            // 错误表单校验
	SchedulerHandler *engine.Scheduler        // 调度
)

// 中英文转义
const (
	ZH = "zh"
	EN = "en"
)
