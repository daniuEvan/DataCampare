/**
 * @date: 2022/3/3
 * @desc: ...
 */

package dbLinkEngine

//
// DataBaseOption
// @Description: 数据库连接信息
//
type DataBaseOption struct {
	DBType     string
	DBHost     string
	DBPort     uint
	DBName     string
	DBUsername string
	DBPassword string
}
