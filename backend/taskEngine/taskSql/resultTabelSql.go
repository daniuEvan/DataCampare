/**
 * @date: 2022/3/3
 * sql
 */

package taskSql

// 结果表SQL
var (
	// ResultTableInitCheckSqlMap 结果表检查是否已经跑过
	ResultTableInitCheckSqlMap = map[string]string{
		"mysql":    `select count(*) from %s.%s WHERE CHECK_TIME = '%s'`,
		"oracle":   `SELECT COUNT(1) FROM %s.%s WHERE ROWNUM = 1 and CHECK_TIME=to_date('%s','yyyy-MM-dd')`,
		"postgres": `select count(*) from %s.%s WHERE CHECK_TIME=to_date('%s','yyyy-MM-dd')`,
		"vertica":  `select count(*) from %s.%s WHERE CHECK_TIME=to_date('%s','yyyy-MM-dd')`,
	}
	// ResultTableInitSqlMap 结果表init sql
	ResultTableInitSqlMap = map[string]string{
		"mysql":    `INSERT INTO %s.%s (OWNER, TABLENAME, SOURCE_NUM, TARGET_NUM,CHECK_TIME) values('%s','%s',0,0, '%s')`,
		"oracle":   `INSERT INTO %s.%s (OWNER, TABLENAME, SOURCE_NUM, TARGET_NUM,CHECK_TIME) values('%s','%s',0,0, to_date('%s','yyyy-MM-dd'))`,
		"postgres": `INSERT INTO %s.%s (OWNER, TABLENAME, SOURCE_NUM, TARGET_NUM,CHECK_TIME) values('%s','%s',0,0, to_date('%s','yyyy-MM-dd'))`,
		"vertica":  `INSERT INTO %s.%s (OWNER, TABLENAME, SOURCE_NUM, TARGET_NUM,CHECK_TIME) values('%s','%s',0,0, to_date('%s','yyyy-MM-dd'))`,
	}

	// ResultTableInsertSqlMap 结果表插入数据sql
	ResultTableInsertSqlMap = map[string]string{
		"mysql":    `UPDATE %s.%s SET  SOURCE_NUM=%s, TARGET_NUM=%s,SOURCE_MAX='%s',TARGET_MAX='%s' where 1=1 and CHECK_TIME='%s' and OWNER = '%s' and TABLENAME='%s'`,
		"oracle":   `UPDATE %s.%s SET  SOURCE_NUM=%s, TARGET_NUM=%s,SOURCE_MAX='%s',TARGET_MAX='%s' where 1=1 and CHECK_TIME=to_date('%s','yyyy-MM-dd') and  OWNER = '%s' and  TABLENAME='%s'`,
		"postgres": `UPDATE %s.%s SET  SOURCE_NUM=%s, TARGET_NUM=%s,SOURCE_MAX='%s',TARGET_MAX='%s' where 1=1 and CHECK_TIME=to_date('%s','yyyy-MM-dd') and  OWNER = '%s' and  TABLENAME='%s'`,
		"vertica":  `UPDATE %s.%s SET  SOURCE_NUM=%s, TARGET_NUM=%s,SOURCE_MAX='%s',TARGET_MAX='%s' where 1=1 and CHECK_TIME=to_date('%s','yyyy-MM-dd') and  OWNER = '%s' and  TABLENAME='%s'`,
	}
)
