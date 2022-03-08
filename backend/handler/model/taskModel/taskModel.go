/**
 * @date: 2022/3/2
 * @desc: 数据比对任务表模型
 */

package taskModel

import (
	"gorm.io/gorm"
)

//
// TaskList
// @Description: 任务列表
//
type TaskList struct {
	gorm.Model
	TaskName         string          `gorm:"type:varchar(100);not null"`
	ConfigDBLinkId   uint            `gorm:"type:int;not null;comment:任务配置表数据库连接id"`
	SourceDBLinkId   uint            `gorm:"type:int;not null;comment:任务源数据库连接id"`
	TargetDBLinkId   uint            `gorm:"type:int;not null;comment:任务目标数据库连接id"`
	ResultDBLinkId   uint            `gorm:"type:int;not null;comment:任务结果表数据库连接id"`
	ResultTableOwner string          `gorm:"type:varchar(50);not null;comment:比对结果表Owner"`
	ResultTableName  string          `gorm:"type:varchar(50);not null;comment:比对结果表name"`
	ConfigTableOwner string          `gorm:"type:varchar(50);not null;comment:中间配置表Owner"`
	ConfigTableName  string          `gorm:"type:varchar(50);not null;comment:中间配置表name"`
	Desc             string          `gorm:"type:text"`
	SchedulerLists   []SchedulerList `gorm:"foreignkey:TaskId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

//
// TaskLog
// @Description: 任务执行日志
//
type TaskLog struct {
	gorm.Model
	TaskId     uint   `gorm:"type:int;not null"`
	TaskStatus uint   `gorm:"type:int;not null;0-执行成功,1-执行失败"`
	Desc       string `gorm:"type:text"`
}
