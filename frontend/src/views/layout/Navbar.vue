<template>
  <el-menu
      id="page-menu"
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      @select="handleSelect"
      background-color="#0166CF"
      text-color="#fff"
      active-text-color="#94C9FF">
    <el-menu-item index="schedulerWatch"><i class="el-icon-view my-icon"></i>调度监控</el-menu-item>
    <el-menu-item index="dbLinkManager"><i class="el-icon-coin my-icon"></i>连接管理</el-menu-item>
    <el-menu-item index="taskManager"><i class="el-icon-guide my-icon"></i>任务管理</el-menu-item>
    <el-menu-item index="schedulerManager"><i class="el-icon-date my-icon"></i>调度管理</el-menu-item>
<!--    <el-menu-item index="5"><a href="https://element.eleme.cn/#/zh-CN/component/menu" target="_blank">ele-ui</a>-->
<!--    </el-menu-item>-->
    <el-submenu index="resultQuery">
      <template slot="title"><i class="el-icon-date my-icon"></i>结果查询</template>
      <el-menu-item index="resultQuery">选项1</el-menu-item>
      <el-menu-item index="2-4-2">选项2</el-menu-item>
      <el-menu-item index="2-4-3">选项3</el-menu-item>
    </el-submenu>
    <div class="readme" @click="readmeVisible = true" >
      <el-tooltip class="item" effect="dark" content="使用说明" placement="left">
        <i class="el-icon-question my-icon"></i>
      </el-tooltip>
    </div>
    <el-dialog title="使用说明" :visible.sync="readmeVisible">
      <Readme></Readme>
    </el-dialog>
  </el-menu>
</template>

<script>
import request from "@/utils/request";
import storageService from "@/service/storageService";
import Readme from "@/views/readme/Readme";

export default {
  name: "Navbar",
  data() {
    return {
      activeIndex: '',
      readmeVisible:false,
    };
  },
  components:{
    Readme
  },
  methods: {
    handleSelect(key,keyPath) {
      console.log(key);
      console.log(keyPath);
      this.activeIndex = key
      this.$router.push({name: key})
    },
    storeDBLinkInfo() {
      let _this = this;
      request.get('/db_link/list/')
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            // 存入浏览器
            storageService.set(storageService.DB_LINK_LIST, JSON.stringify(response.data.data))
          })
          .catch(function (err) {
            _this.$message.error(err)
          })

    },
    storeTaskInfo() {
      let _this = this;
      request.get('/task/list/')
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            // 存入浏览器
            storageService.set(storageService.TASK_INFO_LIST, JSON.stringify(response.data.data))
          })
          .catch(function (err) {
            _this.$message.error(err)
          })

    }
  },
  mounted: function () {
    let _this = this;
    setTimeout(function () {
      if (_this.$route.name !== null) {
        _this.activeIndex = _this.$route.name
      }
    }, 100);
    this.storeDBLinkInfo()
    this.storeTaskInfo()
  }
}
</script>

<style scoped>
.router-link-active {
  text-decoration: none;
}

a {
  text-decoration: none;
}

.my-icon {
  color: #ffffff;
  margin-top: -2px;
}

.readme {
  float: right;
  margin-right: 20px;
  height: 60px;
  line-height: 60px;
  color: #e5e5e5;
  cursor: pointer;
}

</style>