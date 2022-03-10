/**
 * @date: 2022/3/8
 * @desc: ...
 */

package taskSql

// 调度表sql
var (
	// SchedulerInfoQuerySQL 调度信息查询
	SchedulerInfoQuerySQL = `SELECT
								s.id sid,
								s.scheduler_name,
								s.task_schedule,
								s.config_table_query_sql,
								s.result_table_insert_sql,
								s.result_table_init_sql,
								s.source_table_query_sql,
								s.target_table_query_sql,
								s.task_concurrent,
								t.id tid,
								t.config_table_owner,
								t.config_table_name,
								t.result_table_owner,
								t.result_table_name,
								t.task_name,
								t.config_db_link_id,
								cl.link_name config_link_name,
								cl.db_type config_db_type,
								cl.db_host config_db_host,
								cl.db_port config_db_port,
								cl.db_name config_db_name,
								cl.db_username config_db_username,
								cl.db_password config_db_password,
								t.source_db_link_id,
								sl.link_name source_link_name,
								sl.db_type source_db_type,
								sl.db_host source_db_host,
								sl.db_port source_db_port,
								sl.db_name source_db_name,
								sl.db_username source_db_username,
								sl.db_password source_db_password,
								t.target_db_link_id,
								tl.link_name target_link_name,
								tl.db_type target_db_type,
								tl.db_host target_db_host,
								tl.db_port target_db_port,
								tl.db_name target_db_name,
								tl.db_username target_db_username,
								tl.db_password target_db_password,
								t.result_db_link_id,
								rl.link_name result_link_name,
								rl.db_type result_db_type,
								rl.db_host result_db_host,
								rl.db_port result_db_port,
								rl.db_name result_db_name,
								rl.db_username result_db_username,
								rl.db_password result_db_password 
							FROM
								data_compare.compare_scheduler_list s
								LEFT JOIN data_compare.compare_task_list t ON s.task_id = t.id
								LEFT JOIN data_compare.compare_db_link cl ON t.config_db_link_id = cl.id
								LEFT JOIN data_compare.compare_db_link sl ON t.source_db_link_id = sl.id
								LEFT JOIN data_compare.compare_db_link tl ON t.target_db_link_id = tl.id
								LEFT JOIN data_compare.compare_db_link rl ON t.result_db_link_id = rl.id 
							WHERE
								1 = 1 
								AND s.scheduler_status IS TRUE 
								AND s.deleted_at IS NULL 
								AND t.deleted_at IS NULL 
								AND cl.deleted_at IS NULL`
)
