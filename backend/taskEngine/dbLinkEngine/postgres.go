/**
 * @date: 2022/3/3
 * @desc: ...
 */

package dbLinkEngine

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

//
// PostgresLink
// @Description: Postgres 连接
//
type PostgresLink struct {
	DBInfo DataBaseOption
	Conn   *sql.DB
}

func NewPostgresLink(dbInfo DataBaseOption) (*PostgresLink, error) {
	// "用户名:密码@[连接方式](主机名:端口号)/数据库名"
	conn, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=disable",
			dbInfo.DBUsername,
			dbInfo.DBPassword,
			dbInfo.DBHost,
			dbInfo.DBPort,
			dbInfo.DBName,
		),
	)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresLink{
		DBInfo: dbInfo,
		Conn:   conn,
	}, err
}

//
// Query
// @Description: 简单查询
// @param sqlStr: 查询sql
// @return map[string]interface{}
//
func (v *PostgresLink) Query(sqlStr string) (map[string]interface{}, error) {
	rows, err := v.Conn.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	queryValueList := make([]interface{}, 0)
	for rows.Next() {
		scanArgs := make([]interface{}, len(columns)) //行数据
		for i := 0; i < len(columns); i++ {
			var values string
			scanArgs[i] = &values
		}
		err = rows.Scan(scanArgs...)
		if err != nil {
			//return nil, err
		}
		queryValueList = append(queryValueList, scanArgs)
	}
	res := map[string]interface{}{
		"columns": columns,
		"values":  queryValueList,
	}
	return res, nil
}

//
// Exec
// @Description: 简单执行sql
// @param sqlStr:
// @return sql.Result:  result.RowsAffected()
// @return error:
//
func (v *PostgresLink) Exec(sqlStr string) (sql.Result, error) {
	result, err := v.Conn.Exec(sqlStr)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Close 关闭数据库连接
func (v *PostgresLink) Close() error {
	err := v.Conn.Close()
	if err != nil {
		return err
	}
	return nil
}
