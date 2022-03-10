/**
 * @date: 2022/3/8
 * @desc: ...
 */

package engine

//
// CronInfo
// @Description: 调度任务信息
//
type CronInfo struct {
	SchedulerInfo SchedulerInfo // 用户写入库里的配置信息 调度任务的详细信息 可查看 字段详情查看taskSql/schedulerSql
	CronScheduler string        //  调度表达式 * * * * *
	CornFunc      func()        // 调度函数
}

//
// SchedulerStartStatus
// @Description: 调度任务启动状态
//
type SchedulerStartStatus struct {
	Status   bool   // 调度状态
	ErrorMsg string // 调度错误信息
}
