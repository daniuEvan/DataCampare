/**
 * @date: 2022/3/9
 * @desc: ...
 */

package engine

import (
	"DataCompare/taskEngine/engineType"
	"log"
	"time"
)

func NewCronHandler(schedulerInfo engineType.SchedulerInfo) engineType.CronHandler {
	nowDate := time.Now().Format("2006-01-02")
	cronFunc := func() {
		coreFactory := NewCronFuncFactory(schedulerInfo)
		coreFactory.checkDateString = nowDate
		coreFactory.maxCheckDateString = nowDate + " 23:59:59"
		// 查询配置表
		if err := coreFactory.queryConfigTable(); err != nil {
			log.Println(err.Error())
			return
		}
		// 初始化结果表
		if err := coreFactory.initResultTable(); err != nil {
			log.Println(err.Error())
			return
		}
		// 构造 source 和 target 查询sql
		go func() {
			coreFactory.buildOriginTableQuerySql()
		}()
		// 查询源表和目标表 将更新sql写入chan
		if err := coreFactory.querySourceAndTargetTable(); err != nil {
			log.Println(err.Error())
			return
		}
		// 更新结果表
		if err := coreFactory.insertResultTable(); err != nil { //
			log.Println(err.Error())
			return
		}
	}
	return engineType.CronHandler{
		SchedulerInfo: schedulerInfo,
		CronScheduler: schedulerInfo["task_schedule"],
		CornFunc:      cronFunc,
	}
}
