/**
 * @date: 2022/3/2
 * @desc: ...
 */

package taskModel

import "gorm.io/gorm"

//
// DBLink
// @Description: 数据库连接信息
//
type DBLink struct {
	gorm.Model
	LinkName   string     `gorm:"type:varchar(100);not null"`
	DBType     string     `gorm:"type:varchar(25);not null;comment:oracle,vertica,mysql,postgres"`
	DBHost     string     `gorm:"type:varchar(100);not null"`
	DBPort     uint       `gorm:"type:int;not null"`
	DBName     string     `gorm:"type:varchar(25);not null"`
	DBUsername string     `gorm:"type:varchar(25);not null"`
	DBPassword string     `gorm:"type:varchar(25);not null"`
	TaskLists1 []TaskList `gorm:"foreignKey:ConfigDBLinkId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskLists2 []TaskList `gorm:"foreignKey:SourceDBLinkId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskLists3 []TaskList `gorm:"foreignKey:TargetDBLinkId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	TaskLists4 []TaskList `gorm:"foreignKey:ResultDBLinkId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
