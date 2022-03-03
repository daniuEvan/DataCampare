/**
 * @date: 2022/3/2
 * @desc: ...
 */

package dbLink

//
// DatabaseLinker
// @Description: 数据库连接接口
//
type DatabaseLinker interface {
	GetConnect() error
	Close() error
	Ping() error          // 测试连接
	Execute(string) error // 执行sql
}
