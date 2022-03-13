/**
 * @date: 2022/3/9
 * @desc: ...
 */

package engine

import (
	"DataCompare/taskEngine/cronLogger"
	"DataCompare/taskEngine/engineType"
	"log"
	"time"
)

func NewCronHandler(schedulerInfo engineType.SchedulerInfo) engineType.CronHandler {
	schedulerName := schedulerInfo["scheduler_name"]
	logger := cronLogger.CronLogger(schedulerName + ".log")
	cronFunc := func() {
		nowDate := time.Now().Format("2006-01-02")
		logger.Info("调度 " + schedulerName + " :开始")
		coreFactory := NewCronFuncFactory(schedulerInfo, logger)
		coreFactory.checkDateString = nowDate
		coreFactory.maxCheckDateString = nowDate + " 23:59:59"
		// 查询配置表
		logger.Info("调度 " + schedulerName + " :查询配置表")
		if err := coreFactory.queryConfigTable(); err != nil {
			logger.Error(err.Error())
			return
		}
		// 初始化结果表
		logger.Info("调度 " + schedulerName + " :初始化结果表")
		if err := coreFactory.initResultTable(); err != nil {
			logger.Error(err.Error())
			return
		}
		// 构造 source 和 target 查询sql
		logger.Info("调度 " + schedulerName + " :构造 source 和 target 查询sql")
		go coreFactory.buildOriginTableQuerySql()
		// 查询源表和目标表 将更新sql写入chan
		logger.Info("调度 " + schedulerName + " :查询源表和目标表")
		go func() {
			if err := coreFactory.querySourceAndTargetTable(); err != nil {
				logger.Error(err.Error())
				return
			}
		}()

		// 更新结果表
		logger.Info("调度 " + schedulerName + " :更新结果表")
		if err := coreFactory.insertResultTable(); err != nil { //
			log.Println(err.Error())
			return
		}
		defer logger.Info("调度 " + schedulerName + " :结束")
	}
	return engineType.CronHandler{
		SchedulerInfo: schedulerInfo,
		CronScheduler: schedulerInfo["task_schedule"],
		CornFunc:      cronFunc,
	}
}
