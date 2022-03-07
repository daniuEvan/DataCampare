/**
 * @date: 2022/3/5
 * @desc: 任务调度模型
 */

package taskModel

import "gorm.io/gorm"

//
// SchedulerList
// @Description:
//
type SchedulerList struct {
	gorm.Model
	TaskId               uint   `gorm:"type:int;not null;comment:任务id"`
	TaskSchedule         string `gorm:"type:varchar(20);comment:调度表达式"`
	ConfigTableQuerySQL  string `gorm:"type:text;not null;comment:配置库查询sql"`
	ResultTableCreateSQL string `gorm:"type:text;not null;comment:结果表创建sql"`
	ResultTableInitSQL   string `gorm:"type:text;not null;comment:结果表初始化sql"`
	SourceTableQuerySQL  string `gorm:"type:text;not null;comment:源端表查询sql"`
	TargetTableQuerySQL  string `gorm:"type:text;not null;comment:目标表查询sql"`
	TaskConcurrent       uint   `gorm:"type:int;default:1;comment:任务并发"`
	SchedulerStatus      int    `gorm:"type:int;not null;comment:任务状态0-启动,1-禁用"`
	Desc                 string `gorm:"type:text"`
}
