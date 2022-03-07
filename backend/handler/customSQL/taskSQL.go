/**
 * @date: 2022/3/7
 * @desc: 任务相关sql
 */

package customSQL

var (
	// TaskInfoSQL 任务信息查询
	TaskInfoSQL = `SELECT
					t.*,
					d1.link_name AS source_db_link_name,
					d1.db_type AS source_db_type,
					d2.link_name AS target_db_link_name ,
					d2.db_type AS target_db_type ,
					d3.link_name AS config_db_link_name,
					d3.db_type AS config_db_type,
					d4.link_name AS result_db_link_name ,
					d4.db_type AS result_db_type 
				FROM
						compare_task_list t
						LEFT JOIN compare_db_link d1 ON t.source_db_link_id = d1.id
						LEFT JOIN compare_db_link d2 ON t.target_db_link_id = d2.id
						LEFT JOIN compare_db_link d3 ON t.config_db_link_id = d3.id
						LEFT JOIN compare_db_link d4 ON t.result_db_link_id = d4.id
				WHERE  1 = 1 
						AND t.deleted_at IS NULL 
						AND d1.deleted_at IS NULL 
						AND d2.deleted_at IS NULL 
						AND d3.deleted_at IS NULL 
						AND d4.deleted_at IS NULL
						and t.id = ?`
	// TaskInfoListSQL 全部任务信息查询
	TaskInfoListSQL = `SELECT
					t.*,
					d1.link_name AS source_db_link_name,
					d1.db_type AS source_db_type,
					d2.link_name AS target_db_link_name ,
					d2.db_type AS target_db_type ,
					d3.link_name AS config_db_link_name,
					d3.db_type AS config_db_type,
					d4.link_name AS result_db_link_name ,
					d4.db_type AS result_db_type 
				FROM
						compare_task_list t
						LEFT JOIN compare_db_link d1 ON t.source_db_link_id = d1.id
						LEFT JOIN compare_db_link d2 ON t.target_db_link_id = d2.id
						LEFT JOIN compare_db_link d3 ON t.config_db_link_id = d3.id
						LEFT JOIN compare_db_link d4 ON t.result_db_link_id = d4.id
				WHERE  1 = 1 
					AND t.deleted_at IS NULL 
					AND d1.deleted_at IS NULL 
					AND d2.deleted_at IS NULL 
					AND d3.deleted_at IS NULL 
					AND d4.deleted_at IS NULL`
)
