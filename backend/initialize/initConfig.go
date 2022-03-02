/*
 * @date: 2021/12/15
 * @desc: ...
 */

package initialize

import (
	"DataCompare/global"
	"DataCompare/utils"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfigFromYaml() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync() // flushes buffer, if any
	debugEnv := utils.IsDebugEnv()
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("./%s-dev.yaml", configFilePrefix)
	if !debugEnv {
		configFileName = fmt.Sprintf("./%s-pro.yaml", configFilePrefix)
		logger, _ = zap.NewProduction()
	}
	sugar := logger.Sugar()
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		sugar.Errorw("配置初始化失败", "msg", err.Error())
		panic(err.Error())
	}
	if err := v.Unmarshal(global.ServerConfig); err != nil {
		sugar.Errorw("配置初始化失败", "msg", err.Error())
		panic(err.Error())
	}
	sugar.Infof("配置文件初始化完成:%v", global.ServerConfig)

	// 监测配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		sugar.Info("配置文件重新初始化:")
		if err := v.ReadInConfig(); err != nil {
			sugar.Errorw("配置重新初始化失败", "msg", err.Error())
			panic(err.Error())
		}
		if err := v.Unmarshal(global.ServerConfig); err != nil {
			sugar.Errorw("配置重新初始化失败", "msg", err.Error())
			panic(err.Error())
		}
		sugar.Infof("配置文件重新初始化完成:\n%v", global.ServerConfig)
	})

}
