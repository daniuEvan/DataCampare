<template>
  <el-container>
    <el-aside width="300px">
      <el-card class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span style="margin-left: 10px;color: #969696">数据比对调度</span>
          <el-button @click="clickSchedulerAdd()" style="float: right; padding: 3px" type="text"><i
              class="el-icon-paperclip">新建调度</i></el-button>
        </div>
        <div v-if="schedulerInfoArr !==null">
          <div v-for="o in schedulerInfoArr.length" :key="o" class="text item">
            <div class="db-icon" @click="clickScheduler(o)"><img style="width: 18px;padding:2px 5px 0 0"
                                                            :src="schedulerLogo"
                                                            alt="">
            </div>
            <div class="db-name" @click="clickScheduler(o)">{{ schedulerInfoArr[o - 1]["SchedulerName"] }}
            </div>
            <div class="db-delete" @click="clickSchedulerDelete(schedulerInfoArr[o - 1]['ID'])" style="color: red"><i
                class="el-icon-delete"></i></div>
            <div class="db-edit" @click="clickSchedulerEdit(schedulerInfoArr[o - 1])" style="color: #409EFF"><i
                class="el-icon-edit"></i></div>
          </div>
        </div>
      </el-card>
    </el-aside>
    <el-container>
      <el-header>{{ schedulerName }}</el-header>
      <el-main v-if="isShow.schedulerInfoMsgShow">
        <SchedulerMsg :schedulerInfo="schedulerInfo" :dbLogo="dbLogo"/>
      </el-main>
      <el-main v-if="isShow.schedulerInfoFormShow">
        <SchedulerEdit :schedulerEditor="schedulerEditor" :dbLogo="dbLogo" :toDefaultShow="toDefaultShow" :getScheduler="getScheduler"/>
      </el-main>
      <el-main v-if="isShow.schedulerCreateInfoFormShow">
        <SchedulerCreate :toDefaultShow="toDefaultShow" :dbLogo="dbLogo" :getScheduler="getScheduler"/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import SchedulerMsg from "@/views/scheduler/SchedulerMsg";
import request from "@/utils/request";
import SchedulerEdit from "@/views/scheduler/SchedulerEdit";
import SchedulerCreate from "@/views/scheduler/SchedulerCreate";
import storageService from "@/service/storageService";

export default {
  name: "SchedulerIndex",
  data() {
    return {
      isShow: {
        schedulerInfoMsgShow: true,  // 展示数据连接信息
        schedulerInfoFormShow: false, // 编辑组件展示
        schedulerCreateInfoFormShow: false, // 提交组件展示
      },
      schedulerInfoArr: [],
      schedulerName: "",
      schedulerInfo: {},
      dbLogo: {
        "mysql": require("@/assets/databaseLogo/mysql.png"),
        "vertica": require("@/assets/databaseLogo/vertica.png"),
        "oracle": require("@/assets/databaseLogo/oracle.png"),
        "postgres": require("@/assets/databaseLogo/postgres.png"),
      },
      schedulerLogo: require("@/assets/databaseLogo/scheduler.png"),
      schedulerEditor: {},
    }
  },
  components: {
    SchedulerMsg,
    SchedulerEdit,
    SchedulerCreate,
  },
  methods: {
    toDefaultShow(item) {
      this.schedulerEditor = {}
      if (item === "edit") {
        this.isShow.schedulerInfoFormShow = true
        this.isShow.schedulerInfoMsgShow = false
        this.isShow.schedulerCreateInfoFormShow = false
        return null
      } else if (item === "create") {
        this.isShow.schedulerInfoFormShow = false
        this.isShow.schedulerInfoMsgShow = false
        this.isShow.schedulerCreateInfoFormShow = true
        return null
      }
      this.isShow.schedulerInfoFormShow = false
      this.isShow.schedulerCreateInfoFormShow = false
      this.isShow.schedulerInfoMsgShow = true
      return null
    },
    clickScheduler(item) {
      this.toDefaultShow()
      // 点击更改值
      this.schedulerInfo = this.schedulerInfoArr[item - 1]
      this.schedulerName = this.schedulerInfo["SchedulerName"]
    },
    getScheduler() {
      let _this = this
      request.get('/scheduler/list/')
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            _this.schedulerInfoArr = response.data.data
            if (_this.schedulerInfoArr === null) {
              _this.schedulerInfo = {}
              _this.schedulerName = ""
              return null
            }
            _this.schedulerInfo = _this.schedulerInfoArr[0]
            _this.schedulerName = _this.schedulerInfo["SchedulerName"]
            // 存入浏览器
            storageService.set(storageService.SCHEDULER_INFO_LIST, JSON.stringify(response.data.data))
          })
          .catch(function (err) {
            _this.$message.error(err)
          })
    },
    clickSchedulerDelete(id) {
      let _this = this
      this.$confirm('此操作将永久删除此调度(不会删除调度内容), 是否继续?', '提示', {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        request.delete(`/scheduler/${id}/`)
            .then(function (response) {
              if (response.data.code !== 200) {
                _this.$message.error({message: '连接删除失败' + response.data.msg,});
                return null
              }
              _this.$message.success({message: '连接删除成功',});
              _this.getScheduler()
              _this.toDefaultShow()
            })
            .catch(function (err) {
              console.log(err);
              _this.$message.error({message: '连接删除失败',});
            })
      })

    },
    clickSchedulerEdit(item) {
      this.schedulerName = item["SchedulerName"]
      this.toDefaultShow("edit")
      this.schedulerEditor = JSON.parse(JSON.stringify(item))
    },
    clickSchedulerAdd() {
      this.schedulerName = "新建调度"
      this.toDefaultShow("create")
    },
  },
  mounted() {
    this.getScheduler()
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