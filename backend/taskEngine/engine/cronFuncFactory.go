/**
 * @date: 2022/3/9
 * @desc: ...
 */

package engine

import (
	"DataCompare/global"
	"DataCompare/taskEngine/dbLinkEngine"
	"DataCompare/taskEngine/engineType"
	"DataCompare/taskEngine/taskSql"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

type CornFuncFactory struct {
	checkDateString             string                   // 检查日期 格式2006-01-02
	maxCheckDateString          string                   // 检查日期 格式2006-01-02 23:59:59
	schedulerInfo               engineType.SchedulerInfo // 所有调度的配置信息
	configTableList             []map[string]string
	sourceAndTargetQuerySqlChan chan map[string]string // {s:sSql,t:tSql,owner,tablename,bd_column}
	resultTableInsertSqlChan    chan string            // result table insert sql
	backendDBOption             dbLinkEngine.DataBaseOption
}

func NewCornFuncFactory(schedulerInfo engineType.SchedulerInfo) *CornFuncFactory {
	backendDBOption := dbLinkEngine.DataBaseOption{
		DBType:     "mysql",
		DBHost:     global.ServerConfig.DatabaseInfo.MysqlInfo.Host,
		DBPort:     uint(global.ServerConfig.DatabaseInfo.MysqlInfo.Port),
		DBName:     global.ServerConfig.DatabaseInfo.MysqlInfo.DBName,
		DBUsername: global.ServerConfig.DatabaseInfo.MysqlInfo.Username,
		DBPassword: global.ServerConfig.DatabaseInfo.MysqlInfo.Password,
	}
	return &CornFuncFactory{
		schedulerInfo:   schedulerInfo,
		backendDBOption: backendDBOption,
	}
}

//
// BuildCronFunc
// @Description: 返回定时任务要执行的函数
// @param schedulerInfo:
// @return func(): 返回定时任务要执行的函数
//
func (c *CornFuncFactory) BuildCronFunc() engineType.CronHandler {
	c.checkDateString = time.Now().Format("2006-01-02")
	c.maxCheckDateString = c.checkDateString + " 23:59:59"
	coreFunc := func() {
		// 查询配置表
		if err := c.queryConfigTable(); err != nil {
			log.Println(err.Error())
			return
		}
		// 初始化结果表
		if err := c.initResultTable(); err != nil {
			log.Println(err.Error())
			return
		}

		// 查询源表和目标表 将更新sql写入chan
		if err := c.querySourceAndTargetTable(); err != nil {
			log.Println(err.Error())
			return
		}
		// 更新结果表
		if err := c.insertResultTable(); err != nil {
			log.Println(err.Error())
			return
		}
	}
	return engineType.CronHandler{
		SchedulerInfo: c.schedulerInfo,
		CronScheduler: c.schedulerInfo["task_schedule"],
		CornFunc:      coreFunc,
	}
}

//
// queryConfigTable
// @Description: 查询配置表
// @receiver c
// @return error:
//
func (c *CornFuncFactory) queryConfigTable() error {
	configTableQuerySql := c.schedulerInfo["config_table_query_sql"]
	go func() {
		// config query sql 写入数据库
		insertSchedulerConfigTableQuerySql := fmt.Sprintf(
			"update compare_scheduler_list set config_table_query_sql= '%s' where id = %s ",
			strings.ReplaceAll(configTableQuerySql, "'", "''"), c.schedulerInfo["sid"],
		)
		backendDBLinker, err := dbLinkEngine.GetDBLinker(c.backendDBOption)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer backendDBLinker.Close()
		_, err = backendDBLinker.Exec(insertSchedulerConfigTableQuerySql)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}()
	configTableDBPort, err := strconv.Atoi(c.schedulerInfo["config_db_port"])
	if err != nil {
		log.Println(err.Error())
		return err
	}
	configDBOption := dbLinkEngine.DataBaseOption{
		DBType:     c.schedulerInfo["config_db_type"],
		DBHost:     c.schedulerInfo["config_db_host"],
		DBPort:     uint(configTableDBPort),
		DBName:     c.schedulerInfo["config_db_name"],
		DBUsername: c.schedulerInfo["config_db_username"],
		DBPassword: c.schedulerInfo["config_db_password"],
	}
	dbLinker, err := dbLinkEngine.GetDBLinker(configDBOption)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	defer dbLinker.Close()
	queryRes, err := dbLinker.Query(configTableQuerySql)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	byteSliceRes, err := json.Marshal(queryRes)
	gResult := gjson.ParseBytes(byteSliceRes)
	infoKeys := gResult.Get("columns").Array()
	infoValues := gResult.Get("values").Array()
	for _, values := range infoValues {
		configInfo := make(map[string]string)
		for j, value := range values.Array() {
			configInfo[infoKeys[j].String()] = value.String()
		}
		c.configTableList = append(c.configTableList, configInfo)
	}
	return nil
}

//
// initResultTable
// @Description: 初始化结果表
// @receiver c
// @return error:
//
func (c *CornFuncFactory) initResultTable() error {
	// config init sql 写入数据库
	go func() {
		insertSchedulerConfigTableInitSql := fmt.Sprintf(
			"update compare_scheduler_list set result_table_init_sql= '%s' where id = %s ",
			strings.ReplaceAll(c.schedulerInfo["result_table_init_sql"], "'", "''"), c.schedulerInfo["sid"],
		)
		backendDBLinker, err := dbLinkEngine.GetDBLinker(c.backendDBOption)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer backendDBLinker.Close()
		_, err = backendDBLinker.Exec(insertSchedulerConfigTableInitSql)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}()
	resultTableInitCheckSql := fmt.Sprintf(c.schedulerInfo["result_table_init_check_sql"], c.checkDateString)
	resultTableDBPort, err := strconv.Atoi(c.schedulerInfo["result_db_port"])
	if err != nil {
		return err
	}
	resultBDOption := dbLinkEngine.DataBaseOption{
		DBType:     c.schedulerInfo["result_db_type"],
		DBHost:     c.schedulerInfo["result_db_host"],
		DBPort:     uint(resultTableDBPort),
		DBName:     c.schedulerInfo["result_db_name"],
		DBUsername: c.schedulerInfo["result_db_username"],
		DBPassword: c.schedulerInfo["result_db_password"],
	}
	dbLinker, err := dbLinkEngine.GetDBLinker(resultBDOption)
	defer dbLinker.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	queryRes, err := dbLinker.Query(resultTableInitCheckSql)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	byteSliceRes, err := json.Marshal(queryRes)
	gResult := gjson.ParseBytes(byteSliceRes)
	infoValues := gResult.Get("values").Array()
	resultQueryNum := infoValues[0].Array()[0].Int()
	fmt.Println("resultQueryNum---:   ", resultQueryNum)
	fmt.Println("infoValues:   ", infoValues)
	if resultQueryNum > 0 {
		return nil
	}
	// 初始化结果表
	resultTableInitSqlChan := make(chan string, 10)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(c.configTableList))
	go func() { // 拼接sql
		for _, item := range c.configTableList {
			resultTableInitSql := fmt.Sprintf(c.schedulerInfo["result_table_init_sql"], item["owner"], item["tablename"], c.checkDateString)
			resultTableInitSqlChan <- resultTableInitSql
		}
	}()
	// 拼接初始化结果表sql (3个协程读取sql 插入到结果表)
	for i := 0; i < 3; i++ {
		go func() {
			for resultTableInitSql := range resultTableInitSqlChan {
				_, err := dbLinker.Exec(resultTableInitSql)
				if err != nil {
					log.Println(err.Error())
				}
				waitGroup.Done()
			}
		}()
	}
	waitGroup.Wait()
	defer close(resultTableInitSqlChan)
	return nil

}

//
// querySourceAndTargetTable
// @Description: 查询源表和结果表
// @receiver c
// @return error:
//
func (c *CornFuncFactory) querySourceAndTargetTable() error {
	c.sourceAndTargetQuerySqlChan = make(chan map[string]string, 100)
	c.resultTableInsertSqlChan = make(chan string, 100)
	// 构造 source 和target 查询sql
	go func() {
		for _, item := range c.configTableList {
			owner := item["owner"]
			tablename := item["tablename"]
			bdColumn := item["bd_column"]
			sourceTableDBType := c.schedulerInfo["source_db_type"]
			targetTableDBType := c.schedulerInfo["target_db_type"]
			var sourceQuerySql string
			var targetQuerySql string
			// 判断是否存在 bd_column 列 拼接 源表和目标表查询sql
			if len(bdColumn) > 0 {
				sourceQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableHasBdQuerySqlMap[sourceTableDBType], bdColumn, owner, tablename, bdColumn, c.maxCheckDateString)
				targetQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableHasBdQuerySqlMap[targetTableDBType], bdColumn, owner, tablename, bdColumn, c.maxCheckDateString)
			} else {
				sourceQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableNoBdQuerySqlMap[sourceTableDBType], owner, tablename)
				targetQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableNoBdQuerySqlMap[targetTableDBType], owner, tablename)
			}
			// 将sql写入channel
			c.sourceAndTargetQuerySqlChan <- map[string]string{
				"sourceQuerySql": sourceQuerySql,
				"targetQuerySql": targetQuerySql,
				"owner":          owner,
				"tablename":      tablename,
				"bd_column":      bdColumn,
			}

		}
		// 将 source 和target 查询sql 写入完毕之后关闭sourceAndTargetQuerySqlChan
		defer close(c.sourceAndTargetQuerySqlChan)
	}()
	// 并发数
	taskConcurrent, err := strconv.Atoi(c.schedulerInfo["task_concurrent"])
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// 获取sql 并查询
	var waitGroup sync.WaitGroup
	waitGroup.Add(taskConcurrent)
	// 结果表的更新sql写入到后台库
	go func() {
		// result insert sql字符串 写入数据库
		insertSchedulerResultTableInsertSql := fmt.Sprintf(
			"update compare_scheduler_list set result_table_insert_sql= '%s' where id = %s ",
			strings.ReplaceAll(c.schedulerInfo["result_table_insert_sql"], "'", "''"), c.schedulerInfo["sid"],
		)
		backendDBLinker, err := dbLinkEngine.GetDBLinker(c.backendDBOption)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer backendDBLinker.Close()
		_, err = backendDBLinker.Exec(insertSchedulerResultTableInsertSql)
		if err != nil {
			log.Println(err.Error())
			return
		}
	}()
	// 构造结果表insert sql
	for i := 0; i < taskConcurrent; i++ {
		go func() {
			sourceDBPort, err := strconv.Atoi(c.schedulerInfo["source_db_port"])
			if err != nil {
				log.Println(err.Error())
				return
			}
			targetDBPort, err := strconv.Atoi(c.schedulerInfo["target_db_port"])
			if err != nil {
				log.Println(err.Error())
				return
			}
			sourceDBOptions := dbLinkEngine.DataBaseOption{
				DBType:     c.schedulerInfo["source_db_type"],
				DBHost:     c.schedulerInfo["source_db_host"],
				DBPort:     uint(sourceDBPort),
				DBName:     c.schedulerInfo["source_db_name"],
				DBUsername: c.schedulerInfo["source_db_username"],
				DBPassword: c.schedulerInfo["source_db_password"],
			}
			targetDBOptions := dbLinkEngine.DataBaseOption{
				DBType:     c.schedulerInfo["target_db_type"],
				DBHost:     c.schedulerInfo["target_db_host"],
				DBPort:     uint(targetDBPort),
				DBName:     c.schedulerInfo["target_db_name"],
				DBUsername: c.schedulerInfo["target_db_username"],
				DBPassword: c.schedulerInfo["target_db_password"],
			}
			sourceDBlinker, err := dbLinkEngine.GetDBLinker(sourceDBOptions)
			if err != nil {
				log.Println(err.Error())
				return
			}
			targetDBlinker, err := dbLinkEngine.GetDBLinker(targetDBOptions)
			if err != nil {
				log.Println(err.Error())
				return
			}
			defer func() {
				sourceDBlinker.Close()
				targetDBlinker.Close()
			}()
			for itemMap := range c.sourceAndTargetQuerySqlChan {
				sourceQuerySql := itemMap["sourceQuerySql"]
				targetQuerySql := itemMap["targetQuerySql"]
				owner := itemMap["owner"]
				tablename := itemMap["tablename"]
				// 查询源端和目标端数据
				sourceQueryRes, err := sourceDBlinker.Query(sourceQuerySql)
				if err != nil {
					log.Println(err.Error())
					return
				}
				sourceQueryResBytes, err := json.Marshal(sourceQueryRes)
				if err != nil {
					log.Println(err.Error())
					return
				}
				sourceValues := gjson.ParseBytes(sourceQueryResBytes).Get("values").Array()[0].Array()
				sourceCount, sourceMax := sourceValues[0].String(), sourceValues[1].String()
				targetQueryRes, err := targetDBlinker.Query(targetQuerySql)
				if err != nil {
					log.Println(err.Error())
					return
				}
				targetQueryResBytes, err := json.Marshal(targetQueryRes)
				if err != nil {
					log.Println(err.Error())
					return
				}
				targetValues := gjson.ParseBytes(targetQueryResBytes).Get("values").Array()[0].Array()
				targetCount, targetMax := targetValues[0].String(), targetValues[1].String()
				// 拼接 结果表更新sql
				resultInsertSql := fmt.Sprintf(c.schedulerInfo["result_table_insert_sql"], sourceCount, targetCount, sourceMax, targetMax, c.checkDateString, owner, tablename)
				c.resultTableInsertSqlChan <- resultInsertSql
			}
			// for range 会在 sourceAndTargetQuerySqlChan 关闭且读取完毕 之后跳出
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	// 结果表插入sql 拼接完毕之后关闭 ResultTableInsertSqlChan
	close(c.resultTableInsertSqlChan)
	return nil
}

//
// insertResultTable
// @Description: 读取chan sql 更新结果表数据
// @receiver c
// @return error:
//
func (c CornFuncFactory) insertResultTable() error {
	taskConcurrent, err := strconv.Atoi(c.schedulerInfo["task_concurrent"])
	if err != nil {
		log.Println(err.Error())
		return err
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(taskConcurrent)
	for i := 0; i < taskConcurrent; i++ {
		resultTableDBPort, err := strconv.Atoi(c.schedulerInfo["result_db_port"])
		if err != nil {
			log.Println(err.Error())
			return err
		}
		resultBDOption := dbLinkEngine.DataBaseOption{
			DBType:     c.schedulerInfo["result_db_type"],
			DBHost:     c.schedulerInfo["result_db_host"],
			DBPort:     uint(resultTableDBPort),
			DBName:     c.schedulerInfo["result_db_name"],
			DBUsername: c.schedulerInfo["result_db_username"],
			DBPassword: c.schedulerInfo["result_db_password"],
		}
		resultDBLinker, err := dbLinkEngine.GetDBLinker(resultBDOption)
		defer resultDBLinker.Close()
		if err != nil {
			log.Println(err.Error())
			return err
		}
		go func() {
			for insertSql := range c.resultTableInsertSqlChan {
				_, err := resultDBLinker.Exec(insertSql)
				if err != nil {
					log.Println(err.Error())
					continue
				}
			}
			// for range 会在 sourceAndTargetQuerySqlChan 关闭且读取完毕 之后跳出
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	return nil
}
