<template>
  <el-row :gutter="24" class="watch-row">
    <el-col :span="16" :offset="4" class="watch-col">
      <el-card class="box-card" shadow="never">
        <div class="text item" style="font-size: 13px;color: #5b5b5b">
          <span style="margin-right: 10px">选择日期: </span>
          <el-date-picker
              v-model="selectedDate"
              type="daterange"
              size="small"
              range-separator="至"
              start-placeholder="开始日期"
              end-placeholder="结束日期">
          </el-date-picker>
          <span style="margin:0 10px"> 选择任务: </span>
          <el-select v-model="selectedTaskId" size="small" placeholder="请选择任务">
            <el-option v-for="item in allTask"
                       :key="item.ID"
                       :label="item.TaskName"
                       :value="item.ID">
            </el-option>
          </el-select>
          <el-button @click="query()" style="margin:0 30px" type="primary" size="mini"> 查询</el-button>
        </div>
      </el-card>
      <TablePage :tableData="tableData"></TablePage>
      <el-pagination
          v-if="dataCount"
          background
          layout="prev, pager, next"
          small
          :current-page="selectedPage"
          @current-change="pageChange"
          :total="dataCount"
      >
      </el-pagination>

    </el-col>
  </el-row>


</template>

<script>

import TablePage from "@/views/resultTableMsg/TablePage";
import storageService from "@/service/storageService";
import request from "@/utils/request";

export default {
  name: "Index",
  components: {
    TablePage
  },
  data() {
    return {
      selectedDate: [new Date(), new Date()],
      selectedTaskId: "",
      selectedPage: 1,
      tableData: [],
      dataCount: 0,
      allTask: [],
      pageSize: 10,
    }
  },
  methods: {
    query() {
      this.selectedPage = 1
      this.getTableData()
    },
    getAllTask() {
      let tasks = storageService.get(storageService.TASK_INFO_LIST)
      if (tasks.length > 0) {
        this.allTask = JSON.parse(tasks)
        this.selectedTaskId = this.allTask[0].ID
      }
    },
    pageChange(currentPage) {
      this.selectedPage = currentPage
      this.getTableData()
    },
    parseDate(dataArray) {
      let startDate, endDate = ""
      let startY = dataArray[0].getFullYear()
      let startM = dataArray[0].getMonth() + 1
      let startD = dataArray[0].getDate()
      let endY = dataArray[1].getFullYear()
      let endM = dataArray[1].getMonth() + 1
      let endD = dataArray[1].getDate()
      startDate = `${startY}-${startM}-${startD}`
      endDate = `${endY}-${endM}-${endD}`
      return [startDate, endDate]
    },
    getTableData() {
      let _this = this
      let dateArray = this.parseDate(this.selectedDate)
      let startDate = dateArray[0]
      let endDate = dateArray[1]
      request.get(
          `/result/result_table/?taskId=${this.selectedTaskId}&pageNum=${this.selectedPage}&pageSize=${this.pageSize}&startDate=${startDate}&endDate=${endDate}`
      )
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            _this.tableData = response.data.data["dataArray"]
            _this.dataCount = Number(response.data.data["count"])
          })
          .catch(function (err) {
            _this.$message.error(err)
          })
    },
  },
  mounted() {
    this.getAllTask()
    this.query()
  }
}
</script>

<style scoped>
.watch-col, .watch-row {
  /*background-color: #e5e5e5;*/
  /*width: 100%;*/
  height: 100%;
}

.watch-col {
  background-color: #fff;
}

.box-card {
  border: none;
}

.el-pagination {
  margin-top: 20px;
}


#page-body /deep/ .el-input__inner {
  background-color: #2c3e50;
  height: 100px;
}


</style>