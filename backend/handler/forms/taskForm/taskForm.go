/**
 * @date: 2022/3/3
 * @desc: ...
 */

package taskForm

//
// TaskForm
// @Description: 任务表Form
//
type TaskForm struct {
	ID               uint   `forms:"ID" json:"ID" binding:"-"`
	TaskName         string `forms:"TaskName" json:"TaskName" binding:"required"`
	ConfigDBLinkId   uint   `forms:"ConfigDBLinkId" json:"ConfigDBLinkId" binding:"required"`
	SourceDBLinkId   uint   `forms:"SourceDBLinkId" json:"SourceDBLinkId" binding:"required"`
	TargetDBLinkId   uint   `forms:"TargetDBLinkId" json:"TargetDBLinkId" binding:"required"`
	ResultDBLinkId   uint   `forms:"ResultDBLinkId" json:"ResultDBLinkId" binding:"required"`
	ResultTableOwner string `forms:"ResultTableOwner" json:"ResultTableOwner" binding:"required"`
	ResultTableName  string `forms:"ResultTableName" json:"ResultTableName" binding:"required"`
	ConfigTableOwner string `forms:"ConfigTableOwner" json:"ConfigTableOwner" binding:"required"`
	ConfigTableName  string `forms:"ConfigTableName" json:"ConfigTableName" binding:"required"`
	TaskConcurrent   uint   `forms:"TaskConcurrent" json:"TaskConcurrent" binding:"required,gte=1"`
	Desc             string `forms:"Desc" json:"Desc"`
}
