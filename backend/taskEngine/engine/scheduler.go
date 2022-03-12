/**
 * @date: 2022/3/7
 * @desc: 调度函数构造器
 */

package engine

import (
	"DataCompare/global"
	"DataCompare/taskEngine/dbLinkEngine"
	"DataCompare/taskEngine/engineType"
	"DataCompare/taskEngine/taskSql"
	"DataCompare/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/tidwall/gjson"
	"strconv"
)

type Scheduler struct {
	schedulerInfoList         []engineType.SchedulerInfo
	startedSchedulerMap       map[int]engineType.SchedulerInfo        // 已启动的调度 { 调度id: { 任务字段:详情 } }
	schedulerWithCronGraphMap map[int]cron.EntryID                    // { 调度id: cron id}
	cron                      *cron.Cron                              // cron 调度器
	schedulerStartStatus      map[int]engineType.SchedulerStartStatus // 调度状态 { 调度id: SchedulerStartStatus }
	backendDBOptions          dbLinkEngine.DataBaseOption
}

func NewScheduler(backendDBOptions dbLinkEngine.DataBaseOption) (*Scheduler, error) {
	link, err := dbLinkEngine.NewMysqlLink(backendDBOptions)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, err
	}
	queryRes, err := link.Query(taskSql.SchedulerInfoQuerySQL)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, err
	}
	byteSliceRes, err := json.Marshal(queryRes)
	if err != nil {
		global.Logger.Error(err.Error())
		return nil, err
	}
	gResult := gjson.ParseBytes(byteSliceRes)
	infoKeys := gResult.Get("columns").Array()
	infoValues := gResult.Get("values").Array()
	schedulerInfoList := make([]engineType.SchedulerInfo, 0)
	for _, values := range infoValues {
		schedulerInfo := engineType.SchedulerInfo{}
		for j, value := range values.Array() {
			schedulerInfo[infoKeys[j].String()] = value.String()
		}
		schedulerInfoList = append(schedulerInfoList, schedulerInfo)
	}
	// 初加工config 表和result 表sql, 解密数据库密码
	for _, info := range schedulerInfoList {
		// 初加工config 表和result 表sql
		configTableDBType := info["config_db_type"]
		configTableQuerySql := fmt.Sprintf(taskSql.ConfigTableQuerySqlMap[configTableDBType], info["config_table_owner"], info["config_table_name"])
		info["config_table_query_sql"] = configTableQuerySql
		resultTableDBType := info["result_db_type"]
		resultTableInsertSql := fmt.Sprintf(taskSql.ResultTableInsertSqlMap[resultTableDBType], info["result_table_owner"], info["result_table_name"], "%s", "%s", "%s", "%s", "%s", "%s", "%s")
		info["result_table_insert_sql"] = resultTableInsertSql
		resultTableInitSql := fmt.Sprintf(taskSql.ResultTableInitSqlMap[resultTableDBType], info["result_table_owner"], info["result_table_name"], "%s", "%s", "%s")
		info["result_table_init_sql"] = resultTableInitSql
		resultTableInitCheckSql := fmt.Sprintf(taskSql.ResultTableInitCheckSqlMap[resultTableDBType], info["result_table_owner"], info["result_table_name"], "%s")
		info["result_table_init_check_sql"] = resultTableInitCheckSql

		// 解密数据库密码
		info["config_db_password"] = utils.DecodeBase64ToString(info["config_db_password"])
		info["source_db_password"] = utils.DecodeBase64ToString(info["source_db_password"])
		info["target_db_password"] = utils.DecodeBase64ToString(info["target_db_password"])
		info["result_db_password"] = utils.DecodeBase64ToString(info["result_db_password"])
	}
	return &Scheduler{
		schedulerInfoList:         schedulerInfoList,
		cron:                      cron.New(cron.WithSeconds()),
		startedSchedulerMap:       make(map[int]engineType.SchedulerInfo),
		schedulerWithCronGraphMap: make(map[int]cron.EntryID),
		schedulerStartStatus:      make(map[int]engineType.SchedulerStartStatus),
		backendDBOptions:          backendDBOptions,
	}, nil
}

//
// GetAllScheduler
// @Description: 获取全部的调度状态和调度信息
// @receiver s
// @return map[int]engineType.SchedulerStartStatus:
//
func (s *Scheduler) GetAllScheduler() map[int]engineType.SchedulerStartStatus {
	return s.schedulerStartStatus
}

//
// BuildCronHandlers
// @Description: 批量构建任务列表
// @return handlerSlice:
// @return err:
//
func (s *Scheduler) BuildCronHandlers() (cronHandlerList []engineType.CronHandler) {
	for _, schedulerInfo := range s.schedulerInfoList {
		cronHandler := NewCronHandler(schedulerInfo) // 构造定时任务
		cronHandlerList = append(cronHandlerList, cronHandler)
	}
	return cronHandlerList
}

//
// BuildCronHandler
// @Description: 单独构建任务
// @param schedulerId:
// @return cronHandler:
// @return err:
//
func (s *Scheduler) BuildCronHandler(schedulerId int) (cronHandler engineType.CronHandler, err error) {
	dbLinker, err := dbLinkEngine.NewMysqlLink(s.backendDBOptions)
	if err != nil {
		global.Logger.Error(err.Error())
		return
	}
	defer dbLinker.Close()
	queryRes, err := dbLinker.Query(taskSql.SchedulerInfoQuerySQL + fmt.Sprintf(" and s.id = %d", schedulerId))
	if err != nil {
		global.Logger.Error(err.Error())
		return cronHandler, err
	}
	byteSliceRes, err := json.Marshal(queryRes)
	if err != nil {
		global.Logger.Error(err.Error())
		return cronHandler, err
	}
	gResult := gjson.ParseBytes(byteSliceRes)
	infoKeys := gResult.Get("columns").Array()
	infoValues := gResult.Get("values").Array()
	if len(infoValues) <= 0 {
		return cronHandler, errors.New("未查询到该调度id")
	}
	schedulerInfo := engineType.SchedulerInfo{}
	for j, value := range infoValues[0].Array() {
		schedulerInfo[infoKeys[j].String()] = value.String()
	}
	// 初加工config 表和result 表sql
	configTableDBType := schedulerInfo["config_db_type"]
	configTableQuerySql := fmt.Sprintf(taskSql.ConfigTableQuerySqlMap[configTableDBType], schedulerInfo["config_table_owner"], schedulerInfo["config_table_name"])
	schedulerInfo["config_table_query_sql"] = configTableQuerySql
	resultTableDBType := schedulerInfo["result_db_type"]
	resultTableInsertSql := fmt.Sprintf(taskSql.ResultTableInsertSqlMap[resultTableDBType], schedulerInfo["result_table_owner"], schedulerInfo["result_table_name"], "%s", "%s", "%s", "%s", "%s", "%s", "%s")
	schedulerInfo["result_table_insert_sql"] = resultTableInsertSql
	resultTableInitSql := fmt.Sprintf(taskSql.ResultTableInitSqlMap[resultTableDBType], schedulerInfo["result_table_owner"], schedulerInfo["result_table_name"], "%s", "%s", "%s")
	schedulerInfo["result_table_init_sql"] = resultTableInitSql
	resultTableInitCheckSql := fmt.Sprintf(taskSql.ResultTableInitCheckSqlMap[resultTableDBType], schedulerInfo["result_table_owner"], schedulerInfo["result_table_name"], "%s")
	schedulerInfo["result_table_init_check_sql"] = resultTableInitCheckSql
	// 解密数据库密码
	schedulerInfo["config_db_password"] = utils.DecodeBase64ToString(schedulerInfo["config_db_password"])
	schedulerInfo["source_db_password"] = utils.DecodeBase64ToString(schedulerInfo["source_db_password"])
	schedulerInfo["target_db_password"] = utils.DecodeBase64ToString(schedulerInfo["target_db_password"])
	schedulerInfo["result_db_password"] = utils.DecodeBase64ToString(schedulerInfo["result_db_password"])
	s.schedulerInfoList = append(s.schedulerInfoList, schedulerInfo)
	//coreFactory := NewCronFuncFactory(schedulerInfo)
	cronHandler = NewCronHandler(schedulerInfo) // 构造定时任务
	return cronHandler, nil
}

//
// CronStart
// @Description: 启动cron
// @receiver s
//
func (s *Scheduler) CronStart() {
	s.cron.Start()
}

//
// CronStop
// @Description: 停止cron
// @receiver s
//
func (s *Scheduler) CronStop() {
	s.cron.Stop()
}

//
// AddCronFunc
// @Description: 添加任务
// @receiver s
//
func (s *Scheduler) AddCronFunc(cornHandler engineType.CronHandler) (entryID cron.EntryID, err error) {
	cornFunc := cornHandler.CornFunc
	cronScheduler := cornHandler.CronScheduler
	schedulerInfo := cornHandler.SchedulerInfo
	schedulerId, err := strconv.Atoi(schedulerInfo["sid"])
	if err != nil {
		global.Logger.Error(err.Error())
		// scheduler status -> false
		s.schedulerStartStatus[schedulerId] = engineType.SchedulerStartStatus{Status: false, SchedulerInfo: schedulerInfo, ErrorMsg: err.Error()}
		return 0, err
	}
	entryID, err = s.cron.AddFunc(cronScheduler, cornFunc)
	if err != nil {
		// scheduler status -> false
		global.Logger.Error(err.Error())
		s.schedulerStartStatus[schedulerId] = engineType.SchedulerStartStatus{Status: false, SchedulerInfo: schedulerInfo, ErrorMsg: err.Error()}
		return 0, err
	}
	// 添加到
	s.startedSchedulerMap[schedulerId] = schedulerInfo
	s.schedulerWithCronGraphMap[schedulerId] = entryID
	// scheduler status -> True
	s.schedulerStartStatus[schedulerId] = engineType.SchedulerStartStatus{Status: true, SchedulerInfo: schedulerInfo}
	return
}

//
// RemoveCronFuncWithSchedulerId
// @Description: schedulerId移除任务
// @receiver s
//
func (s *Scheduler) RemoveCronFuncWithSchedulerId(schedulerId int) {
	entryID := s.schedulerWithCronGraphMap[schedulerId]
	s.cron.Remove(entryID)
	delete(s.schedulerWithCronGraphMap, schedulerId)
	delete(s.startedSchedulerMap, schedulerId)
	delete(s.schedulerStartStatus, schedulerId)
}

//
// RemoveCronFuncWithEntryId
// @Description: 通过entryID移除任务
// @receiver s
//
func (s *Scheduler) RemoveCronFuncWithEntryId(entryId cron.EntryID) {
	s.cron.Remove(entryId)
	var schedulerId int
	for key, value := range s.schedulerWithCronGraphMap {
		if value == entryId {
			schedulerId = key
		}
	}
	delete(s.schedulerWithCronGraphMap, schedulerId)
	delete(s.startedSchedulerMap, schedulerId)
	delete(s.schedulerStartStatus, schedulerId)
}

//
// ClearCronFunc
// @Description: 移除全部任务
// @receiver s
//
func (s *Scheduler) ClearCronFunc() {
	for _, entryID := range s.schedulerWithCronGraphMap {
		s.cron.Remove(entryID)
	}
	s.startedSchedulerMap = map[int]engineType.SchedulerInfo{}
	s.schedulerWithCronGraphMap = map[int]cron.EntryID{}
	s.schedulerStartStatus = map[int]engineType.SchedulerStartStatus{}
}
