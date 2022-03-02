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
	LinkId     uint   `gorm:"type:int;primarykey"`
	LinkName   string `gorm:"type:varchar(100);not null"`
	DBType     string `gorm:"type:varchar(25);not null;comment:0-oracle,1-vertica,2-mysql,3-postgres"`
	DBHost     string `gorm:"type:varchar(100);not null"`
	DBPort     uint   `gorm:"type:int;not null"`
	DBName     string `gorm:"type:varchar(25);not null"`
	DBUsername string `gorm:"type:varchar(25);not null"`
	DBPassword string `gorm:"type:varchar(25);not null"`
	LinkType   uint   `gorm:"type:int;not null;comment:0-源端数据库,1-目标端数据库"`
}
