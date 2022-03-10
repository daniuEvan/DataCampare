/**
 * @date: 2022/3/7
 * @desc: ...
 */

package engine

import "github.com/robfig/cron/v3"

type SchedulerInterface interface {
	BuildCornHandler() (cronInfoList []CronInfo)
	CronStart()
	CronStop()
	AddFunc(cornInfo CronInfo) (entryID cron.EntryID, err error)
	RemoveWithSchedulerId(schedulerId int)
	RemoveWithEntryId(entryId cron.EntryID)
	Clear()
}
