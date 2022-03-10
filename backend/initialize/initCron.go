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
	"strconv"
)

func InitCron() {
	backendDBInfo := global.ServerConfig.DatabaseInfo.MysqlInfo
	dbOptions := dbLinkEngine.DataBaseOption{
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
	cronInfoList := global.SchedulerHandler.BuildCornHandler()
	for _, cronInfo := range cronInfoList {
		schedulerId, err := strconv.Atoi(cronInfo.SchedulerInfo["sid"])
		if err != nil {
			global.Logger.Error(err.Error(), zap.String("初始化调度", err.Error()))
			return
		}
		_, err = global.SchedulerHandler.AddFunc(cronInfo)
		if err != nil {
			global.Logger.Error(err.Error(), zap.String("初始化调度", err.Error()))
			global.SchedulerHandler.SchedulerStartStatus = map[int]engine.SchedulerStartStatus{
				schedulerId: engine.SchedulerStartStatus{Status: false, ErrorMsg: err.Error()},
			}
			continue
		}
		global.SchedulerHandler.SchedulerStartStatus = map[int]engine.SchedulerStartStatus{
			schedulerId: engine.SchedulerStartStatus{Status: true},
		}
	}
	global.SchedulerHandler.Cron.Start()
}
