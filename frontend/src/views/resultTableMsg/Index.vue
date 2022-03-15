<template>
  <el-row :gutter="24" class="watch-row">
    <el-col :span="16" :offset="4" class="watch-col" >
      <el-card class="box-card" shadow="never" >
        <div  class="text item" style="font-size: 13px;color: #5b5b5b">
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
          <el-select  v-model="selectedTaskId"  size="small"  placeholder="请选择任务">
            <el-option v-for="item in allTask"
                       :key="item.ID"
                       :label="item.TaskName"
                       :value="item.ID">
            </el-option>
          </el-select>
          <el-button @click="query()" style="margin:0 30px" type="primary" size="mini"> 查询 </el-button>
        </div>
      </el-card>
      <TablePage :tableData="tableData"></TablePage>
      <el-pagination
          background
          layout="prev, pager, next"
          small
          @current-change="pageChange"
          :total="20"
      >
      </el-pagination>

    </el-col>
  </el-row>


</template>

<script>

import TablePage from "@/views/resultTableMsg/TablePage";
import storageService from "@/service/storageService";

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
      tableData: [
        {
          date: '2016-05-02',
          name: '王小虎',
          address: '金沙江路 1518 弄'
        }, {
          date: '2016-05-04',
          name: '王小虎',
          address: '金沙江路 1517 弄'
        }, {
          date: '2016-05-01',
          name: '王小虎',
          address: '金沙江路 1519 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }, {
          date: '2016-05-03',
          name: '王小虎',
          address: '金沙江路 1516 弄'
        }
      ],
      allTask:[],
    }
  },
  methods: {
    query(){
        console.log(this.selectedTaskId)
        console.log(this.selectedPage)
        console.log(this.selectedDate)
    },
    getAllTask(){
      let tasks = storageService.get(storageService.TASK_INFO_LIST)
      if(tasks.length>0){
        this.allTask = JSON.parse(tasks)
        this.selectedTaskId = this.allTask[0].ID
      }
    },
    pageChange(currentPage){
      this.selectedPage = currentPage
    }
  },
  mounted() {
    this.getAllTask()

  }
}
</script>

<style scoped>
.watch-col, .watch-row {
  /*background-color: #e5e5e5;*/
  /*width: 100%;*/
  height: 100%;
}
.watch-col{
  background-color: #fff;
}
.box-card{
  border: none;
}

.el-pagination{
  margin-top: 20px;
}


#page-body /deep/ .el-input__inner{
  background-color: #2c3e50;
  height: 100px;
}


</style>