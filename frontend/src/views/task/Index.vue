<template>
  <el-container>
    <el-aside width="300px">
      <el-card class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span style="margin-left: 10px;color: #969696">数据比对任务</span>
          <el-button @click="clickTaskAdd()" style="float: right; padding: 3px" type="text"><i
              class="el-icon-paperclip">新建任务</i></el-button>
        </div>
        <div v-if="taskInfoArr !==null">
          <div v-for="o in taskInfoArr.length" :key="o" class="text item">
            <div class="db-icon" @click="clickTask(o)"><img style="width: 18px;padding:2px 5px 0 0"
                                                            :src="taskLogo"
                                                            alt="">
            </div>
            <div class="db-name" @click="clickTask(o)">{{ taskInfoArr[o - 1]["TaskName"] }}
            </div>
            <div class="db-delete" @click="clickTaskDelete(taskInfoArr[o - 1]['ID'])" style="color: red"><i
                class="el-icon-delete"></i></div>
            <div class="db-edit" @click="clickTaskEdit(taskInfoArr[o - 1])" style="color: #409EFF"><i
                class="el-icon-edit"></i></div>
          </div>
        </div>
      </el-card>
    </el-aside>
    <el-container>
      <el-header>{{ taskName }}</el-header>
      <el-main v-if="isShow.taskInfoMsgShow">
        <TaskMsg :taskInfo="taskInfo" :dbLogo="dbLogo"/>
      </el-main>
      <el-main v-if="isShow.taskInfoFormShow">
        <TaskEdit :taskEditor="taskEditor" :dbLogo="dbLogo" :toDefaultShow="toDefaultShow" :getTask="getTask"/>
      </el-main>
      <el-main v-if="isShow.taskCreateInfoFormShow">
        <TaskCreate :toDefaultShow="toDefaultShow" :dbLogo="dbLogo" :getTask="getTask"/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import TaskMsg from "@/views/task/TaskMsg";
import request from "@/utils/request";
import TaskEdit from "@/views/task/TaskEdit";
import TaskCreate from "@/views/task/TaskCreate";
import storageService from "@/service/storageService";

export default {
  name: "TaskIndex",
  data() {
    return {
      isShow: {
        taskInfoMsgShow: true,  // 展示数据连接信息
        taskInfoFormShow: false, // 编辑组件展示
        taskCreateInfoFormShow: false, // 提交组件展示
      },
      taskInfoArr: [],
      taskName: "",
      taskInfo: {},
      dbLogo: {
        "mysql": require("@/assets/databaseLogo/mysql.png"),
        "vertica": require("@/assets/databaseLogo/vertica.png"),
        "oracle": require("@/assets/databaseLogo/oracle.png"),
        "postgres": require("@/assets/databaseLogo/postgres.png"),
      },
      taskLogo: require("@/assets/databaseLogo/task.png"),
      taskEditor: {},
    }
  },
  components: {
    TaskMsg,
    TaskEdit,
    TaskCreate,
  },
  methods: {
    toDefaultShow(item) {
      this.taskEditor = {}
      if (item === "edit") {
        this.isShow.taskInfoFormShow = true
        this.isShow.taskInfoMsgShow = false
        this.isShow.taskCreateInfoFormShow = false
        return null
      } else if (item === "create") {
        this.isShow.taskInfoFormShow = false
        this.isShow.taskInfoMsgShow = false
        this.isShow.taskCreateInfoFormShow = true
        return null
      }
      this.isShow.taskInfoFormShow = false
      this.isShow.taskCreateInfoFormShow = false
      this.isShow.taskInfoMsgShow = true
      return null
    },
    clickTask(item) {
      this.toDefaultShow()
      // 点击更改值
      this.taskInfo = this.taskInfoArr[item - 1]
      this.taskName = this.taskInfo["TaskName"]
    },
    getTask() {
      let _this = this
      request.get('/task/list/')
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            _this.taskInfoArr = response.data.data
            if (_this.taskInfoArr === null) {
              _this.taskInfo = {}
              _this.taskName = ""
              return null
            }
            _this.taskInfo = _this.taskInfoArr[0]
            _this.taskName = _this.taskInfo["TaskName"]
            // 存入浏览器
            storageService.set(storageService.TASK_INFO_LIST, JSON.stringify(response.data.data))
          })
          .catch(function (err) {
            _this.$message.error(err)
          })
    },
    clickTaskDelete(id) {
      let _this = this
      this.$confirm('此操作将永久删除此任务(不会删除调度内容), 是否继续?', '提示', {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        request.delete(`/task/${id}/`)
            .then(function (response) {
              if (response.data.code !== 200) {
                _this.$message.error({message: '连接删除失败' + response.data.msg,});
                return null
              }
              _this.$message.success({message: '连接删除成功',});
              _this.getTask()
              _this.toDefaultShow()
            })
            .catch(function (err) {
              console.log(err);
              _this.$message.error({message: '连接删除失败',});
            })
      })

    },
    clickTaskEdit(item) {
      this.taskName = item["TaskName"]
      this.toDefaultShow("edit")
      this.taskEditor = JSON.parse(JSON.stringify(item))
    },
    clickTaskAdd() {
      this.taskName = "新建任务"
      this.toDefaultShow("create")
    },
  },
  mounted() {
    this.getTask()
  }
}
</script>

<style scoped>
.el-header {
  background-color: #f3f8fd;
  color: #333;
  font-weight: bold;
  line-height: 40px;
  text-align: center;
  height: 40px !important;
}

.el-aside {
  background-color: #E9EEF3;
  color: #333;
  /*line-height: 200px;*/
}

.el-main {
  background-color: #f3f8fd;
  color: #333;
  padding-top: 50px;
  /*text-align: center;*/
  line-height: 100%;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container {
  height: 100%;
}

.el-card /deep/ .el-card__header {
  padding: 7px !important;
  border-bottom: 1px solid #d3d3d3;
}

.el-card /deep/ .el-card__body {
  padding: 10px 20px !important;
}

.box-card {
  background-color: #E9EEF3;
}

.item {
  padding: 5px;
  line-height: 20px;
  /*font-size: 15px;*/
  color: #333333;
  border-bottom: 1px solid #e5e5e5;
}

.db-icon, .db-name {
  display: inline-block;
  vertical-align: middle;
  font-size: 15px;
  color: #494848;
  cursor: pointer;

}

.db-name {
  padding-bottom: 2px;
}

.db-delete, .db-edit {
  display: inline-block;
  vertical-align: middle;
  float: right;
  cursor: pointer;
  font-size: 15px;
  margin: 3px;
}

</style>