/**
 * @date: 2022/3/3
 * @desc: ...
 */

package taskForm

//
// SchedulerForm
// @Description: 调度表Form
//
type SchedulerForm struct {
	ID                   uint   `forms:"ID" json:"ID" binding:"-"`
	SchedulerName        string `forms:"SchedulerName" json:"SchedulerName" binding:"required"`
	TaskId               uint   `forms:"TaskId" json:"TaskId" binding:"required"`
	TaskSchedule         string `forms:"TaskSchedule" json:"TaskSchedule" binding:"required"`
	ConfigTableQuerySQL  string `forms:"ConfigTableQuerySQL" json:"ConfigTableQuerySQL"`
	ResultTableCreateSQL string `forms:"ResultTableCreateSQL" json:"ResultTableCreateSQL"`
	ResultTableInitSQL   string `forms:"ResultTableInitSQL" json:"ResultTableInitSQL"`
	SourceTableQuerySQL  string `forms:"SourceTableQuerySQL" json:"SourceTableQuerySQL"`
	TargetTableQuerySQL  string `forms:"TargetTableQuerySQL" json:"TargetTableQuerySQL"`
	SchedulerStatus      bool   `forms:"SchedulerStatus" json:"SchedulerStatus" `
	TaskConcurrent       uint   `forms:"TaskConcurrent" json:"TaskConcurrent" binding:"required,gte=1"`
	Desc                 string `forms:"Desc" json:"Desc"`
}
