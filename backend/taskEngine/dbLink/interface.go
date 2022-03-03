/**
 * @date: 2022/3/3
 * @desc: ...
 */

package dbLink

//
// DBLinker
// @Description: 数据库连接抽象层
//
type DBLinker interface {
	Query(sqlStr string) (map[string]interface{}, error)
	Exec(sqlStr string) (string, error)
	Close() error
}
