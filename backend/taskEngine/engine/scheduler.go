/**
 * @date: 2022/3/7
 * @desc: 调度函数构造器
 */

package engine

import (
	"DataCompare/taskEngine/dbLinkEngine"
	"DataCompare/taskEngine/taskSql"
	"DataCompare/utils"
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/tidwall/gjson"
	"log"
	"strconv"
	"time"
)

//
// SchedulerInfo
// @Description: 调度详细信息
//
type SchedulerInfo map[string]string // 用户写入库里的配置信息 调度任务的详细信息 可查看 字段详情查看taskSql/schedulerSql

type Scheduler struct {
	CheckDateString           string // 检查日期 2006-01-02
	SchedulerInfoList         []SchedulerInfo
	StartedSchedulerMap       map[int]SchedulerInfo        // 已启动的调度 { 调度id: { 任务字段:详情 } }
	SchedulerWithCronGraphMap map[int]cron.EntryID         // { 调度id: cron id}
	Cron                      *cron.Cron                   // cron 调度器
	SchedulerStartStatus      map[int]SchedulerStartStatus // 调度状态 { 调度id: SchedulerStartStatus }
}

func NewScheduler(backendDBOptions dbLinkEngine.DataBaseOption) (*Scheduler, error) {
	nowDate := time.Now().Format("2006-01-02")
	link, err := dbLinkEngine.NewMysqlLink(backendDBOptions)
	if err != nil {
		log.Fatalln(err.Error())
		return nil, err
	}
	queryRes, err := link.Query(taskSql.SchedulerInfoQuerySQL)
	byteSliceRes, err := json.Marshal(queryRes)
	gResult := gjson.ParseBytes(byteSliceRes)
	infoKeys := gResult.Get("columns").Array()
	infoValues := gResult.Get("values").Array()
	schedulerInfoSlice := make([]SchedulerInfo, 0)
	for _, values := range infoValues {
		schedulerInfo := SchedulerInfo{}
		for j, value := range values.Array() {
			schedulerInfo[infoKeys[j].String()] = value.String()
		}
		schedulerInfoSlice = append(schedulerInfoSlice, schedulerInfo)
	}
	if err != nil {
		return nil, err
	}
	// 初加工config 表和result 表sql, 解密数据库密码
	for _, info := range schedulerInfoSlice {
		// 初加工config 表和result 表sql
		configTableDBType := info["config_db_type"]
		configTableQuerySql := fmt.Sprintf(taskSql.ConfigTableQuerySqlMap[configTableDBType], info["config_table_owner"], info["config_table_name"])
		info["config_table_query_sql"] = configTableQuerySql
		resultTableDBType := info["result_db_type"]
		resultTableInsertSql := fmt.Sprintf(taskSql.ResultTableInsertSqlMap[resultTableDBType], info["result_table_owner"], info["result_table_name"], "%s", "%s", "%s", "%s", nowDate, "%s", "%s")
		info["result_table_insert_sql"] = resultTableInsertSql
		resultTableInitSql := fmt.Sprintf(taskSql.ResultTableInitSqlMap[resultTableDBType], info["result_table_owner"], info["result_table_name"], "%s", "%s", nowDate)
		info["result_table_init_sql"] = resultTableInitSql
		resultTableInitCheckSql := fmt.Sprintf(taskSql.ResultTableInitCheckSqlMap[resultTableDBType], info["result_table_owner"], info["result_table_name"], nowDate)
		info["result_table_init_check_sql"] = resultTableInitCheckSql

		// 解密数据库密码
		info["config_db_password"] = utils.DecodeBase64ToString(info["config_db_password"])
		info["source_db_password"] = utils.DecodeBase64ToString(info["source_db_password"])
		info["target_db_password"] = utils.DecodeBase64ToString(info["target_db_password"])
		info["result_db_password"] = utils.DecodeBase64ToString(info["result_db_password"])
	}
	return &Scheduler{
		SchedulerInfoList:         schedulerInfoSlice,
		CheckDateString:           nowDate,
		Cron:                      cron.New(),
		StartedSchedulerMap:       make(map[int]SchedulerInfo),
		SchedulerWithCronGraphMap: make(map[int]cron.EntryID),
		SchedulerStartStatus:      make(map[int]SchedulerStartStatus),
	}, nil
}

//
// BuildCornHandler
// @Description: 构建任务列表
// @return handlerSlice:
// @return err:
//
func (s *Scheduler) BuildCornHandler() (cronInfoList []CronInfo) {
	for _, schedulerInfo := range s.SchedulerInfoList {
		cronHandler := NewCornFuncFactory(schedulerInfo, s.CheckDateString)
		cronInfo := cronHandler.BuildCronFunc() // 构造定时任务
		cronInfoList = append(cronInfoList, cronInfo)
	}
	return cronInfoList
}

//
// CronStart
// @Description: 启动cron
// @receiver s
//
func (s *Scheduler) CronStart() {
	s.Cron.Start()
}

//
// CronStop
// @Description: 停止cron
// @receiver s
//
func (s *Scheduler) CronStop() {
	s.Cron.Stop()
}

//
// AddFunc
// @Description: 添加任务
// @receiver s
//
func (s *Scheduler) AddFunc(cornInfo CronInfo) (entryID cron.EntryID, err error) {
	cornFunc := cornInfo.CornFunc
	cronScheduler := cornInfo.CronScheduler
	schedulerInfo := cornInfo.SchedulerInfo
	schedulerId, err := strconv.Atoi(schedulerInfo["sid"])
	if err != nil {
		log.Fatalln(err.Error())
		return 0, err
	}
	entryID, err = s.Cron.AddFunc(cronScheduler, cornFunc)
	if err != nil {
		log.Fatalln(err.Error())
		return 0, err
	}
	// 添加到
	s.StartedSchedulerMap[schedulerId] = schedulerInfo
	s.SchedulerWithCronGraphMap[schedulerId] = entryID
	return
}

//
// RemoveWithSchedulerId
// @Description: schedulerId移除任务
// @receiver s
//
func (s *Scheduler) RemoveWithSchedulerId(schedulerId int) {
	entryID := s.SchedulerWithCronGraphMap[schedulerId]
	s.Cron.Remove(entryID)
	delete(s.SchedulerWithCronGraphMap, schedulerId)
	delete(s.StartedSchedulerMap, schedulerId)
}

//
// RemoveWithEntryId
// @Description: 通过entryID移除任务
// @receiver s
//
func (s *Scheduler) RemoveWithEntryId(entryId cron.EntryID) {
	s.Cron.Remove(entryId)
	var schedulerId int
	for key, value := range s.SchedulerWithCronGraphMap {
		if value == entryId {
			schedulerId = key
		}
	}
	delete(s.SchedulerWithCronGraphMap, schedulerId)
	delete(s.StartedSchedulerMap, schedulerId)
}

//
// Clear
// @Description: 移除全部任务
// @receiver s
//
func (s *Scheduler) Clear() {
	for _, entryID := range s.SchedulerWithCronGraphMap {
		s.Cron.Remove(entryID)
	}
	s.StartedSchedulerMap = map[int]SchedulerInfo{}
	s.SchedulerWithCronGraphMap = map[int]cron.EntryID{}
}
