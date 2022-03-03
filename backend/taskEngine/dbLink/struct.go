/**
 * @date: 2022/3/3
 * @desc: ...
 */

package dbLink

//
// DataBaseOption
// @Description: 数据库连接信息
//
type DataBaseOption struct {
	DBType     string
	DBHost     string
	DBPort     int
	DBName     string
	DBUsername string
	DBPassword string
}
