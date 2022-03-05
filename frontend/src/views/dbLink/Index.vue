<template>
  <el-container>
    <el-aside width="20%">
      <el-card class="box-card" shadow="never">
        <div slot="header" class="clearfix">
          <span style="margin-left: 10px;color: #969696">数据库连接</span>
          <el-button style="float: right; padding: 3px" type="text"><i class="el-icon-paperclip"></i>添加</el-button>
        </div>
        <div v-for="o in dbLinkInfoArr.length" :key="o" class="text item">
          <div class="db-icon"><img style="width: 16px;margin:1px 3px 0 0" :src="dbLogo[dbLinkInfoArr[o - 1]['DBType']]" alt="">
          </div>
          <div class="db-name" @click="clickDBLink(o)">{{ dbLinkInfoArr[o - 1]["LinkName"] }}
          </div>
          <div class="db-delete"  @click="clickDBLinkDelete(dbLinkInfoArr[o - 1]['ID'])" style="color: red"><i class="el-icon-delete"></i></div>
          <div class="db-edit" @click="clickDBLinkEdit(dbLinkInfoArr[o - 1]['ID'])" style="color: #409EFF"><i class="el-icon-edit"></i></div>
        </div>
      </el-card>
    </el-aside>
    <el-container>
      <el-header>{{ dbLinkInfo['LinkName'] }}</el-header>
      <el-main>
        <DBLinkMsg :dbLinkInfo="dbLinkInfo"/>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import DBLinkMsg from "@/views/dbLink/DBLinkMsg";
import request from "@/utils/request";

export default {
  name: "index",
  data() {
    return {
      dbLinkInfoArr: [],
      dbLinkInfo: {},
      dbLogo: {
        "mysql": require("@/assets/databaseLogo/mysql.png"),
        "vertica": require("@/assets/databaseLogo/vertica.png"),
        "oracle": require("@/assets/databaseLogo/oracle.png"),
        "postgres": require("@/assets/databaseLogo/postgres.png"),
      },
    }
  },
  components: {
    DBLinkMsg
  },
  methods: {
    clickDBLink(item) {
      // 点击更改值
      this.dbLinkInfo = this.dbLinkInfoArr[item - 1]
    },
    clickDBLinkDelete(id){
      request.delete('/db_link/')
          .then(function (response) {
            if (response.data.code !== 200) {
              alert(response.data.msg)
              return null
            }
            _this.dbLinkInfoArr = response.data.data
            _this.dbLinkInfo = _this.dbLinkInfoArr[0]
          })
          .catch(function (err) {
            console.log(err);
          })
    },
    clickDBLinkEdit(id){

    },
  },
  mounted() {
    let _this = this
    request.get('/db_link/list')
        .then(function (response) {
          if (response.data.code !== 200) {
            alert(response.data.msg)
            return null
          }
          _this.dbLinkInfoArr = response.data.data
          _this.dbLinkInfo = _this.dbLinkInfoArr[0]
        })
        .catch(function (err) {
          console.log(err);
        })
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