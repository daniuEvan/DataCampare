/**
 * @date: 2022/3/9
 * @desc: ...
 */

package engine

import (
	"DataCompare/taskEngine/dbLinkEngine"
	"DataCompare/taskEngine/taskSql"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"sync"
)

type CornFuncFactory struct {
	CheckDateString             string        // 检查日期 格式2006-01-02
	MaxCheckDateString          string        // 检查日期 格式2006-01-02 23:59:59
	SchedulerInfo               SchedulerInfo // 所有调度的配置信息
	ConfigTableList             []map[string]string
	SourceAndTargetQuerySqlChan chan map[string]string // {s:sSql,t:tSql,owner,tablename,bd_column}
	ResultTableInsertSqlChan    chan string            // result table insert sql
}

func NewCornFuncFactory(schedulerInfo SchedulerInfo, checkDateString string) *CornFuncFactory {
	return &CornFuncFactory{
		SchedulerInfo:      schedulerInfo,
		CheckDateString:    checkDateString,
		MaxCheckDateString: checkDateString + " 23:59:59",
	}
}

//
// BuildCronFunc
// @Description: 返回定时任务要执行的函数
// @param schedulerInfo:
// @return func(): 返回定时任务要执行的函数
//
func (c *CornFuncFactory) BuildCronFunc() CronInfo {
	coreFunc := func() {
		// 查询配置表
		if err := c.queryConfigTable(); err != nil {
			log.Fatalln(err.Error())
			return
		}
		// 初始化结果表
		if err := c.initResultTable(); err != nil {
			log.Fatalln(err.Error())
			return
		}

		// 查询源表和目标表 将更新sql写入chan
		if err := c.querySourceAndTargetTable(); err != nil {
			log.Fatalln(err.Error())
			return
		}
		// 更新结果表
		if err := c.insertResultTable(); err != nil {
			log.Fatalln(err.Error())
			return
		}
	}
	return CronInfo{
		SchedulerInfo: c.SchedulerInfo,
		CronScheduler: c.SchedulerInfo["task_schedule"],
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
	configTableQuerySql := c.SchedulerInfo["config_table_query_sql"]
	configTableDBPort, err := strconv.Atoi(c.SchedulerInfo["config_db_port"])
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	configBDOption := dbLinkEngine.DataBaseOption{
		DBType:     c.SchedulerInfo["config_db_type"],
		DBHost:     c.SchedulerInfo["config_db_host"],
		DBPort:     uint(configTableDBPort),
		DBName:     c.SchedulerInfo["config_db_name"],
		DBUsername: c.SchedulerInfo["config_db_username"],
		DBPassword: c.SchedulerInfo["config_db_password"],
	}
	dbLinker, err := dbLinkEngine.GetDBLinker(configBDOption)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	defer dbLinker.Close()
	queryRes, err := dbLinker.Query(configTableQuerySql)
	if err != nil {
		log.Fatalln(err.Error())
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
		c.ConfigTableList = append(c.ConfigTableList, configInfo)
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
	resultTableInitCheckSql := c.SchedulerInfo["result_table_init_check_sql"]
	resultTableDBPort, err := strconv.Atoi(c.SchedulerInfo["result_db_port"])
	if err != nil {
		return err
	}
	resultBDOption := dbLinkEngine.DataBaseOption{
		DBType:     c.SchedulerInfo["result_db_type"],
		DBHost:     c.SchedulerInfo["result_db_host"],
		DBPort:     uint(resultTableDBPort),
		DBName:     c.SchedulerInfo["result_db_name"],
		DBUsername: c.SchedulerInfo["result_db_username"],
		DBPassword: c.SchedulerInfo["result_db_password"],
	}
	dbLinker, err := dbLinkEngine.GetDBLinker(resultBDOption)
	defer dbLinker.Close()
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	queryRes, err := dbLinker.Query(resultTableInitCheckSql)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	byteSliceRes, err := json.Marshal(queryRes)
	gResult := gjson.ParseBytes(byteSliceRes)
	infoValues := gResult.Get("values").Array()
	resultQueryNum := infoValues[0].Array()[0].Int()
	fmt.Println("resultQueryNum---:   ", resultQueryNum)
	if resultQueryNum > 0 {
		return nil
	}
	// 初始化结果表
	resultTableInitSqlChan := make(chan string, 10)
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(c.ConfigTableList))
	go func() { // 拼接sql
		for _, item := range c.ConfigTableList {
			resultTableInitSql := fmt.Sprintf(c.SchedulerInfo["result_table_init_sql"], item["owner"], item["tablename"])
			resultTableInitSqlChan <- resultTableInitSql
		}
	}()
	// 拼接初始化结果表sql (3个协程读取sql 插入到结果表)
	for i := 0; i < 3; i++ {
		go func() {
			for resultTableInitSql := range resultTableInitSqlChan {
				_, err := dbLinker.Exec(resultTableInitSql)
				if err != nil {
					log.Fatalln(err.Error())
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
	c.SourceAndTargetQuerySqlChan = make(chan map[string]string, 100)
	c.ResultTableInsertSqlChan = make(chan string, 100)
	// 构造 source 和target 查询sql
	go func() {
		for _, item := range c.ConfigTableList {
			owner := item["owner"]
			tablename := item["tablename"]
			bdColumn := item["bd_column"]
			sourceTableDBType := c.SchedulerInfo["source_db_type"]
			targetTableDBType := c.SchedulerInfo["target_db_type"]
			var sourceQuerySql string
			var targetQuerySql string
			// 判断是否存在 bd_column 列 拼接 源表和目标表查询sql
			if len(bdColumn) > 0 {
				sourceQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableHasBdQuerySqlMap[sourceTableDBType], bdColumn, owner, tablename, bdColumn, c.MaxCheckDateString)
				targetQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableHasBdQuerySqlMap[targetTableDBType], bdColumn, owner, tablename, bdColumn, c.MaxCheckDateString)
			} else {
				sourceQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableNoBdQuerySqlMap[sourceTableDBType], owner, tablename)
				targetQuerySql = fmt.Sprintf(taskSql.SourceAndTargetTableNoBdQuerySqlMap[targetTableDBType], owner, tablename)
			}
			// 将sql写入channel
			c.SourceAndTargetQuerySqlChan <- map[string]string{
				"sourceQuerySql": sourceQuerySql,
				"targetQuerySql": targetQuerySql,
				"owner":          owner,
				"tablename":      tablename,
				"bd_column":      bdColumn,
			}

		}
		// 将 source 和target 查询sql 写入完毕之后关闭SourceAndTargetQuerySqlChan
		defer close(c.SourceAndTargetQuerySqlChan)
	}()
	// 并发数
	taskConcurrent, err := strconv.Atoi(c.SchedulerInfo["task_concurrent"])
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	// 获取sql 并查询
	var waitGroup sync.WaitGroup
	waitGroup.Add(taskConcurrent)
	for i := 0; i < taskConcurrent; i++ {
		go func() {
			sourceDBPort, err := strconv.Atoi(c.SchedulerInfo["source_db_port"])
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			targetDBPort, err := strconv.Atoi(c.SchedulerInfo["target_db_port"])
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			sourceDBOptions := dbLinkEngine.DataBaseOption{
				DBType:     c.SchedulerInfo["source_db_type"],
				DBHost:     c.SchedulerInfo["source_db_host"],
				DBPort:     uint(sourceDBPort),
				DBName:     c.SchedulerInfo["source_db_name"],
				DBUsername: c.SchedulerInfo["source_db_username"],
				DBPassword: c.SchedulerInfo["source_db_password"],
			}
			targetDBOptions := dbLinkEngine.DataBaseOption{
				DBType:     c.SchedulerInfo["target_db_type"],
				DBHost:     c.SchedulerInfo["target_db_host"],
				DBPort:     uint(targetDBPort),
				DBName:     c.SchedulerInfo["target_db_name"],
				DBUsername: c.SchedulerInfo["target_db_username"],
				DBPassword: c.SchedulerInfo["target_db_password"],
			}
			sourceDBlinker, err := dbLinkEngine.GetDBLinker(sourceDBOptions)
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			targetDBlinker, err := dbLinkEngine.GetDBLinker(targetDBOptions)
			if err != nil {
				log.Fatalln(err.Error())
				return
			}
			defer func() {
				sourceDBlinker.Close()
				targetDBlinker.Close()
			}()
			for itemMap := range c.SourceAndTargetQuerySqlChan {
				sourceQuerySql := itemMap["sourceQuerySql"]
				targetQuerySql := itemMap["targetQuerySql"]
				owner := itemMap["owner"]
				tablename := itemMap["tablename"]
				// 查询源端和目标端数据
				sourceQueryRes, err := sourceDBlinker.Query(sourceQuerySql)
				if err != nil {
					log.Fatalln(err.Error())
					return
				}
				sourceQueryResBytes, err := json.Marshal(sourceQueryRes)
				if err != nil {
					log.Fatalln(err.Error())
					return
				}
				sourceValues := gjson.ParseBytes(sourceQueryResBytes).Get("values").Array()[0].Array()
				sourceCount, sourceMax := sourceValues[0].String(), sourceValues[1].String()
				targetQueryRes, err := targetDBlinker.Query(targetQuerySql)
				if err != nil {
					log.Fatalln(err.Error())
					return
				}
				targetQueryResBytes, err := json.Marshal(targetQueryRes)
				if err != nil {
					log.Fatalln(err.Error())
					return
				}
				targetValues := gjson.ParseBytes(targetQueryResBytes).Get("values").Array()[0].Array()
				targetCount, targetMax := targetValues[0].String(), targetValues[1].String()
				// 拼接 结果表更新sql
				resultInsertSql := fmt.Sprintf(c.SchedulerInfo["result_table_insert_sql"], sourceCount, targetCount, sourceMax, targetMax, owner, tablename)
				c.ResultTableInsertSqlChan <- resultInsertSql
			}
			// for range 会在 SourceAndTargetQuerySqlChan 关闭且读取完毕 之后跳出
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	// 结果表插入sql 拼接完毕之后关闭 ResultTableInsertSqlChan
	close(c.ResultTableInsertSqlChan)
	return nil
}

//
// insertResultTable
// @Description: 读取chan sql 更新结果表数据
// @receiver c
// @return error:
//
func (c CornFuncFactory) insertResultTable() error {
	taskConcurrent, err := strconv.Atoi(c.SchedulerInfo["task_concurrent"])
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	var waitGroup sync.WaitGroup
	waitGroup.Add(taskConcurrent)
	for i := 0; i < taskConcurrent; i++ {
		resultTableDBPort, err := strconv.Atoi(c.SchedulerInfo["result_db_port"])
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		resultBDOption := dbLinkEngine.DataBaseOption{
			DBType:     c.SchedulerInfo["result_db_type"],
			DBHost:     c.SchedulerInfo["result_db_host"],
			DBPort:     uint(resultTableDBPort),
			DBName:     c.SchedulerInfo["result_db_name"],
			DBUsername: c.SchedulerInfo["result_db_username"],
			DBPassword: c.SchedulerInfo["result_db_password"],
		}
		resultDBLinker, err := dbLinkEngine.GetDBLinker(resultBDOption)
		defer resultDBLinker.Close()
		if err != nil {
			log.Fatalln(err.Error())
			return err
		}
		go func() {
			for insertSql := range c.ResultTableInsertSqlChan {
				_, err := resultDBLinker.Exec(insertSql)
				if err != nil {
					log.Fatalln(err.Error())
					continue
				}
			}
			// for range 会在 SourceAndTargetQuerySqlChan 关闭且读取完毕 之后跳出
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	return nil
}
