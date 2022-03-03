/**
 * @date: 2022/3/3
 * @desc: ...
 */

package engine

import "database/sql"

type MiddleDBInfo struct {
	MiddleDBLink *sql.DB
	Owner        string
	Table        string
}

type CompareEngine struct {
	MiddleDB     MiddleDBInfo
	SourceDBLink *sql.DB
	TargetDBLink *sql.DB
}

//func NewCompareEngine(MiddleDBLink, SourceDBLink, TargetDBLink *sql.DB) *CompareEngine {
//
//}
func _() {
	// 1. 查询配置表, 初始化 配置表(按照月 按照天)
	// 2. 查询源表信息
	// 3. 查询目标表信息
	// 4. 写入数据库
	// 5.
}
