/**
 * @date: 2022/3/3
 * @desc: ...
 */

package taskForm

import "time"

//
// TaskForm
// @Description: 任务表Form
//
type TaskForm struct {
	TaskName         string    `forms:"TaskName" json:"TaskName" binding:"required"`
	RunTime          time.Time `forms:"RunTime" json:"RunTime"`
	TaskSchedule     string    `forms:"TaskSchedule" json:"TaskSchedule"`
	SourceDBLinkId   uint      `forms:"SourceDBLinkId" json:"SourceDBLinkId" binding:"required"`
	TargetDBLinkId   uint      `forms:"TargetDBLinkId" json:"TargetDBLinkId" binding:"required"`
	ResultTableOwner string    `forms:"ResultTableOwner" json:"ResultTableOwner" binding:"required"`
	ResultTableName  string    `forms:"ResultTableName" json:"ResultTableName" binding:"required"`
	ConfigTableOwner string    `forms:"ConfigTableOwner" json:"ConfigTableOwner" binding:"required"`
	ConfigTableName  string    `forms:"ConfigTableName" json:"ConfigTableName" binding:"required"`
	TaskConcurrent   uint      `forms:"TaskConcurrent" json:"TaskConcurrent" binding:"required;gt=1"`
}
