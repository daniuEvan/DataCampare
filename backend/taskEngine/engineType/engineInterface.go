/**
 * @date: 2022/3/7
 * @desc: ...
 */

package engineType

import (
	"github.com/robfig/cron/v3"
)

type SchedulerInterface interface {
	GetAllScheduler() map[int]SchedulerStartStatus
	BuildCronHandlers() (cronHandlerList []CronHandler)
	BuildCronHandler(schedulerId int) (cronHandler CronHandler, err error)
	CronStart()
	CronStop()
	AddCronFunc(cronHandler CronHandler) (entryID cron.EntryID, err error)
	RemoveCronFuncWithSchedulerId(schedulerId int)
	RemoveCronFuncWithEntryId(entryId cron.EntryID)
	ClearCronFunc()
}
