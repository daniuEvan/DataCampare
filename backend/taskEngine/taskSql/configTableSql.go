/**
 * @date: 2022/3/8
 * @desc: ...
 */

package taskSql

//  配置表SQL
var (
	// ConfigTableQuerySqlMap 配置表查询SQL
	ConfigTableQuerySqlMap = map[string]string{
		"mysql":    `select DISTINCT owner,tablename,bd_column from %s.%s`,
		"oracle":   `select DISTINCT owner,tablename,bd_column from %s.%s`,
		"postgres": `select DISTINCT owner,tablename,bd_column from %s.%s`,
		"vertica":  `select DISTINCT owner,tablename,bd_column from %s.%s`,
	}
)
