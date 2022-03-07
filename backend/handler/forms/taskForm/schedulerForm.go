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
	TaskId               uint   `forms:"TaskId" json:"TaskId" binding:"required"`
	TaskSchedule         string `forms:"TaskSchedule" json:"TaskSchedule" binding:"required"`
	ConfigTableQuerySQL  string `forms:"ConfigTableQuerySQL" json:"ConfigTableQuerySQL" binding:"required"`
	ResultTableCreateSQL string `forms:"ResultTableCreateSQL" json:"ResultTableCreateSQL" binding:"required"`
	ResultTableInitSQL   string `forms:"ResultTableInitSQL" json:"ResultTableInitSQL" binding:"required"`
	SourceTableQuerySQL  string `forms:"SourceTableQuerySQL" json:"SourceTableQuerySQL" binding:"required"`
	TargetTableQuerySQL  string `forms:"TargetTableQuerySQL" json:"TargetTableQuerySQL" binding:"required"`
	SchedulerStatus      int    `forms:"SchedulerStatus" json:"SchedulerStatus" binding:"required,gte=0,lte=1"`
	TaskConcurrent       uint   `forms:"TaskConcurrent" json:"TaskConcurrent" binding:"required,gte=1"`
	Desc                 string `forms:"Desc" json:"Desc"`
}
