/**
 * @date: 2022/3/2
 * @desc: 数据比对任务表模型
 */

package taskModel

import (
	"gorm.io/gorm"
	"time"
)

//
// TaskList
// @Description: 任务列表
//
type TaskList struct {
	gorm.Model
	TaskId            uint   `gorm:"type:int;primarykey"`
	TaskName          string `gorm:"type:varchar(100);not null"`
	RunTime           time.Time
	SourceDBLinkId    uint   `gorm:"type:int;not null;comment:任务源数据库连接id"`
	TargetDBLinkId    uint   `gorm:"type:int;not null;comment:任务目标数据库连接id"`
	ResultTableSchema string `gorm:"type:varchar(50);not null;comment:比对结果表schema"`
	ResultTableName   string `gorm:"type:varchar(50);not null;comment:比对结果表name"`
	ConfigTableSchema string `gorm:"type:varchar(50);not null;comment:中间配置表schema"`
	ConfigTableName   string `gorm:"type:varchar(50);not null;comment:中间配置表name"`
	TaskConcurrent    uint   `gorm:"type:int;not null;comment:任务并发"`
	TaskSchedule      string `gorm:"type:varchar(50);comment:任务计划"` // todo 具体设计
}

//
// TaskLog
// @Description: 任务执行日志
//
type TaskLog struct {
	gorm.Model
	TaskId     uint `gorm:"type:int;not null"`
	TaskStatus uint `gorm:"type:int;not null;0-执行成功,1-执行失败"`
}
