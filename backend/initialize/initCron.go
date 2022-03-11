/**
 * @date: 2022/3/10
 * @desc: ...
 */

package initialize

import (
	"DataCompare/global"
	"DataCompare/taskEngine/dbLinkEngine"
	"DataCompare/taskEngine/engine"
	"go.uber.org/zap"
)

func InitCron() {
	backendDBInfo := global.ServerConfig.DatabaseInfo.MysqlInfo
	dbOptions := dbLinkEngine.DataBaseOption{
		DBType:     "mysql",
		DBHost:     backendDBInfo.Host,
		DBPort:     uint(backendDBInfo.Port),
		DBName:     backendDBInfo.DBName,
		DBUsername: backendDBInfo.Username,
		DBPassword: backendDBInfo.Password,
	}
	schedulerHandler, err := engine.NewScheduler(dbOptions)
	if err != nil {
		global.Logger.Error(err.Error(), zap.String("初始化调度", err.Error()))
		return
	}
	global.SchedulerHandler = schedulerHandler
	global.SchedulerHandler.CronStart()
	cronHandlerList := global.SchedulerHandler.BuildCronHandlers()
	for _, cronHandler := range cronHandlerList {
		_, err = global.SchedulerHandler.AddCronFunc(cronHandler)
		if err != nil {
			global.Logger.Error(err.Error(), zap.String("初始化调度", err.Error()))
		}
	}
	// 单独初始化
	//cronHandler1, err := global.SchedulerHandler.BuildCronHandler(5)
	//if err != nil {
	//	global.Logger.Error(err.Error(), zap.String("初始化调度", err.Error()))
	//	log.Println(err)
	//	return
	//}
	//_, err = global.SchedulerHandler.AddCronFunc(cronHandler1)
}
