/*
 * @date: 2021/12/16
 * @desc: ...
 */

package model

import (
	"DataCompare/handler/model/taskModel"
)

// ModelsArr 所有model示例
var ModelsArr = []interface{}{
	//&userModel.User{},
	&taskModel.TaskList{},
	&taskModel.TaskLog{},
	&taskModel.SchedulerList{},
	&taskModel.DBLink{},
}
