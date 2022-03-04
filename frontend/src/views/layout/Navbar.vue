<template>
  <div>
    <el-menu
        :default-active="activeIndex2"
        class="el-menu-demo"
        mode="horizontal"
        @select="handleSelect"
        background-color="#0166CF"
        text-color="#fff"
        active-text-color="#94C9FF">
      <!--      <el-menu-item index="1">调度中心</el-menu-item>-->
      <!--      <el-menu-item index="1">任务中心</el-menu-item>-->
      <!--      <el-menu-item index="1">任务中心</el-menu-item>-->
      <el-menu-item index="3">
          连接管理
      </el-menu-item>
      <el-menu-item index="6">
          连接管理3
      </el-menu-item>

      <el-menu-item index="5"><a href="https://element.eleme.cn/#/zh-CN/component/menu" target="_blank">ele-ui</a>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "Navbar",
  data() {
    return {
      activeIndex: '1',
      activeIndex2: '1'
    };
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    }
  }
}
</script>

<style scoped>
.router-link-active{
  text-decoration: none;
}
a {
  text-decoration: none;
}
</style>