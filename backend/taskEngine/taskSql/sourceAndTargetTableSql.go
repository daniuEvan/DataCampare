/**
 * @date: 2022/3/8
 * @desc: 源表sql
 */

package taskSql

var (
	// SourceAndTargetTableHasBdQuerySqlMap 源表查询sql(存在bd_column)
	SourceAndTargetTableHasBdQuerySqlMap = map[string]string{
		"mysql":    `select count(*) count,max(%s) max from %s.%s where %s <= %s`,
		"oracle":   `select /*+ parallel(16)*/count(*) count,max(%s) max from %s.%s where %s <= to_date('%s','YYYYMMDD HH24:MI:SS')`,
		"postgres": `select count(*) count,max(%s) max from %s.%s where %s <= to_timestamp('%s','YYYYMMDD HH24:MI:SS')`,
		"vertica":  `select count(*) count,max(%s) max from %s.%s where %s <= to_timestamp('%s','YYYYMMDD HH24:MI:SS')`,
	}
	// SourceAndTargetTableNoBdQuerySqlMap 源表查询sql(不存在bd_column)
	SourceAndTargetTableNoBdQuerySqlMap = map[string]string{
		"mysql":    `select count(*) count,'-' max from %s.%s `,
		"oracle":   `select /*+ parallel(16)*/count(*) count,'-' max from %s.%s `,
		"postgres": `select count(*) count,'-'  max from %s.%s `,
		"vertica":  `select count(*) count,'-'  max from %s.%s `,
	}
)
