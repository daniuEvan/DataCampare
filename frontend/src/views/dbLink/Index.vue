<template>
  <el-container>
    <el-aside width="300px">
      <el-card class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span style="margin-left: 10px;color: #969696">数据库连接</span>
          <el-button @click="clickDBLinkAdd()" style="float: right; padding: 3px" type="text"><i
              class="el-icon-paperclip">新建连接</i></el-button>
        </div>
        <div v-if="dbLinkInfoArr !==null">
          <div v-for="o in dbLinkInfoArr.length" :key="o" class="text item">
            <div class="db-icon" @click="clickDBLink(o)"><img style="width: 18px;padding:2px 5px 0 0"
                                                              :src="dbLogo[dbLinkInfoArr[o - 1]['DBType']]"
                                                              alt="">
            </div>
            <div class="db-name" @click="clickDBLink(o)">{{ dbLinkInfoArr[o - 1]["LinkName"] }}
            </div>
            <div class="db-delete" @click="clickDBLinkDelete(dbLinkInfoArr[o - 1]['ID'])" style="color: red"><i
                class="el-icon-delete"></i></div>
            <div class="db-edit" @click="clickDBLinkEdit(dbLinkInfoArr[o - 1])" style="color: #409EFF"><i
                class="el-icon-edit"></i></div>
          </div>

        </div>
      </el-card>
    </el-aside>
    <el-container>
      <el-header>{{ dbLinkName }}</el-header>
      <el-main v-if="isShow.dbLinkInfoMsgShow">
        <DBLinkMsg :dbLinkInfo="dbLinkInfo"/>
      </el-main>
      <el-main v-if="isShow.dbLinkInfoFormShow">
        <DBLinkEdit :dbLinkEditor="dbLinkEditor" :dbLogo="dbLogo" :toDefaultShow="toDefaultShow" :getDBLink="getDBLink"/>
      </el-main>
      <el-main v-if="isShow.dbLinkCreateInfoFormShow">
        <DBLinkCreate :toDefaultShow="toDefaultShow" :dbLogo="dbLogo" :getDBLink="getDBLink"/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import DBLinkMsg from "@/views/dbLink/DBLinkMsg";
import request from "@/utils/request";
import DBLinkEdit from "@/views/dbLink/DBLinkEdit";
import DBLinkCreate from "@/views/dbLink/DBLinkCreate";
import storageService from "@/service/storageService";

export default {
  name: "DBLinkIndex",
  data() {
    return {
      isShow: {
        dbLinkInfoMsgShow: true,  // 展示数据连接信息
        dbLinkInfoFormShow: false, // 编辑组件展示
        dbLinkCreateInfoFormShow: false, // 提交组件展示
      },
      dbLinkInfoArr: [],
      dbLinkName: "",
      dbLinkInfo: {},
      dbLogo: {
        "mysql": require("@/assets/databaseLogo/mysql.png"),
        "vertica": require("@/assets/databaseLogo/vertica.png"),
        "oracle": require("@/assets/databaseLogo/oracle.png"),
        "postgres": require("@/assets/databaseLogo/postgres.png"),
      },
      dbLinkEditor: {},
    }
  },
  components: {
    DBLinkMsg,
    DBLinkEdit,
    DBLinkCreate,
  },
  methods: {
    toDefaultShow(item) {
      this.dbLinkEditor = {}
      if (item === "edit") {
        this.isShow.dbLinkInfoFormShow = true
        this.isShow.dbLinkInfoMsgShow = false
        this.isShow.dbLinkCreateInfoFormShow = false
        return null
      } else if (item === "create") {
        this.isShow.dbLinkInfoFormShow = false
        this.isShow.dbLinkInfoMsgShow = false
        this.isShow.dbLinkCreateInfoFormShow = true
        return null
      }
      this.isShow.dbLinkInfoFormShow = false
      this.isShow.dbLinkCreateInfoFormShow = false
      this.isShow.dbLinkInfoMsgShow = true
      return null
    },
    clickDBLink(item) {
      this.toDefaultShow()
      // 点击更改值
      this.dbLinkInfo = this.dbLinkInfoArr[item - 1]
      this.dbLinkName = this.dbLinkInfo["LinkName"]
    },
    getDBLink() {
      let _this = this
      request.get('/db_link/list/')
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            _this.dbLinkInfoArr = response.data.data
            if (_this.dbLinkInfoArr.length <= 0) {
              _this.dbLinkInfo = {}
              _this.dbLinkName = ""
              return null
            }
            _this.dbLinkInfo = _this.dbLinkInfoArr[0]
            _this.dbLinkName = _this.dbLinkInfo["LinkName"]
            // 存入浏览器
            storageService.set(storageService.DB_LINK_LIST, JSON.stringify(response.data.data))
          })
          .catch(function (err) {
            _this.$message.error(err)
          })
    },
    clickDBLinkDelete(id) {
      let _this = this
      this.$confirm('此操作将永久删除此链接, 是否继续?', '提示', {
        confirmButtonText: '确定删除',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        request.delete(`/db_link/${id}/`)
            .then(function (response) {
              if (response.data.code !== 200) {
                _this.$message.error({message: '连接删除失败' + response.data.msg,});
                return null
              }
              _this.$message.success({message: '连接删除成功',});
              _this.getDBLink()
            })
            .catch(function (err) {
              console.log(err);
              _this.$message.error({message: '连接删除失败',});
            })
      })

    },
    clickDBLinkEdit(item) {
      this.dbLinkName = item["LinkName"]
      this.toDefaultShow("edit")
      this.dbLinkEditor = JSON.parse(JSON.stringify(item)) // 深拷贝
    },
    clickDBLinkAdd() {
      this.dbLinkName = "新建数据库连接"
      this.toDefaultShow("create")
    },
  },
  mounted() {
    this.getDBLink()
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