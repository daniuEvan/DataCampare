<template>
  <el-container>
    <el-aside width="20%">
      <el-card class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span style="margin-left: 10px;color: #969696">数据库连接</span>
          <el-button style="float: right; padding: 3px" type="text"><i class="el-icon-paperclip"></i>添加</el-button>
        </div>
        <div v-for="o in dbLinkInfoArr.length" :key="o" class="text item">
          <div class="db-icon"><img style="width: 16px;margin:1px 3px 0 0" :src="dbLogo[dbLinkInfoArr[o - 1]['DBType']]"
                                    alt="">
          </div>
          <div class="db-name" @click="clickDBLink(o)">{{ dbLinkInfoArr[o - 1]["LinkName"] }}
          </div>
          <div class="db-delete" @click="clickDBLinkDelete(dbLinkInfoArr[o - 1]['ID'])" style="color: red"><i
              class="el-icon-delete"></i></div>
          <div class="db-edit" @click="clickDBLinkEdit(dbLinkInfoArr[o - 1])" style="color: #409EFF"><i
              class="el-icon-edit"></i></div>
        </div>
      </el-card>
    </el-aside>
    <el-container>
      <el-header>{{ dbLinkName }}</el-header>
      <el-main v-show="dbLinkInfoMsgShow">
        <DBLinkMsg :dbLinkInfo="dbLinkInfo"/>
      </el-main>
      <el-main v-show="dbLinkInfoFormShow">
        <DBLinkForm :dbLinkInfo="dbLinkEditor"/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import DBLinkMsg from "@/views/dbLink/DBLinkMsg";
import request from "@/utils/request";
import DBLinkForm from "@/views/dbLink/DBLinkForm";

export default {
  name: "index",
  data() {
    return {
      dbLinkInfoMsgShow: true,  // 展示数据连接信息
      dbLinkInfoFormShow: false, // 编辑组件
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
    DBLinkForm,
  },
  methods: {
    toDefaultShow(ok) {
      if (!ok) {
        this.dbLinkInfoFormShow = true
        this.dbLinkInfoMsgShow = false
        return null
      }
      this.dbLinkInfoFormShow = false
      this.dbLinkInfoMsgShow = true
      return null
    },
    clickDBLink(item) {
      this.toDefaultShow(true)
      // 点击更改值
      this.dbLinkInfo = this.dbLinkInfoArr[item - 1]
      this.dbLinkName = this.dbLinkInfo["LinkName"]
    },
    getDBLink() {
      let _this = this
      request.get('/db_link/list')
          .then(function (response) {
            if (response.data.code !== 200) {
              alert(response.data.msg)
              return null
            }
            _this.dbLinkInfoArr = response.data.data
            _this.dbLinkInfo = _this.dbLinkInfoArr[0]
            _this.dbLinkName = _this.dbLinkInfo["LinkName"]

          })
          .catch(function (err) {
            console.log(err);
          })

    },
    clickDBLinkDelete(id) {
      let _this = this
      request.delete(`/db_link/${id}`)
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
    },
    clickDBLinkEdit(item) {
      this.dbLinkName = item["LinkName"]
      this.toDefaultShow()
      this.dbLinkEditor = item

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
  line-height: 15px;
  /*font-size: 15px;*/
  color: #333333;
  border-bottom: 1px solid #e5e5e5;
}

.db-icon, .db-name {
  display: inline-block;
  vertical-align: middle;
  font-size: 13px;
  color: #494848;
  cursor: pointer;
}

.db-delete, .db-edit {
  display: inline-block;
  vertical-align: middle;
  float: right;
  cursor: pointer;
  font-size: 10px;
  margin: 3px;
}

</style>