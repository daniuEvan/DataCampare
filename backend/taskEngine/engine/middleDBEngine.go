/**
 * @date: 2022/3/3
 * @desc: 中间配置表
 */

package engine

import (
	"DataCompare/taskEngine/taskSql"
	"database/sql"
	"fmt"
)

type MiddleDBEngine struct {
	MiddleDBLink *sql.DB
	Owner        string
	TableName    string
	QuerySql     string
}

func NewMiddleDBEngine(dbOption DataBaseOption, owner string, tableName string) *MiddleDBEngine {
	//middleDBLink:=
	querySql := fmt.Sprintf(taskSql.QueryMiddleTable, owner, tableName)
	return &MiddleDBEngine{
		MiddleDBLink: middleDBLink,
		Owner:        owner,
		TableName:    tableName,
		QuerySql:     querySql,
	}
}

func (m *MiddleDBEngine) GetMiddleTableData() {
	rows, err := m.MiddleDBLink.Exec(m.QuerySql)
	defer rows.Close()
	if err != nil {
		return
	}
}
