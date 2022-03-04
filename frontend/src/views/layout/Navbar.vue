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
    <el-menu-item index="taskCenter"><i class="el-icon-folder-opened my-icon"></i>任务中心</el-menu-item>
    <el-menu-item index="taskManager"><i class="el-icon-edit my-icon"></i>任务管理</el-menu-item>
    <el-menu-item index="schedulerManager"><i class="el-icon-date my-icon"></i>调度管理</el-menu-item>
    <el-menu-item index="dbLinkManager"><i class="el-icon-coin my-icon"></i>连接管理</el-menu-item>
    <el-menu-item index="5"><a href="https://element.eleme.cn/#/zh-CN/component/menu" target="_blank">ele-ui</a>
    </el-menu-item>
  </el-menu>
</template>

<script>
export default {
  name: "Navbar",
  data() {
    return {
      activeIndex: 'taskCenter',
    };
  },
  methods: {
    handleSelect(key) {
      this.activeIndex = key
      this.$router.push({name: key})
    }
  },
  mounted: function () {
    let _this = this;
    setTimeout(function() {
      if (_this.$route.name !== null){
        _this.activeIndex = _this.$route.name
      }else{
        _this.activeIndex = "taskCenter"
      }
    }, 100);

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

</style>