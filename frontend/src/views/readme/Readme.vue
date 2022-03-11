<template>
  <div class="block">
    <el-timeline>
      <el-timeline-item timestamp="第一步" placement="top">
        <el-card style="width: 90%">
          <p>点击"连接管理"创建数据库连接</p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="第二步" placement="top">
        <el-card style="width: 90%">
          <p>点击"任务管理"创建要执行的作业</p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="第三步" placement="top">
        <el-card style="width: 90%">
          <p>点击"调度管理"创建作业调度,仅支持crontab表达式(* * * * *)</p>
        </el-card>
      </el-timeline-item>
      <el-timeline-item timestamp="第四步" placement="top">
        <el-card style="width: 90%">
          <p>点击"调度监控"查看调度是否启动成功</p>
          <p><i class="el-icon-warning"></i>调度启动成功的不代表作业执行成功</p>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>

</template>

<script>
export default {
  name: "Readme",
  props:["readmeVisible"]
}
</script>

<style scoped>
.el-timeline /deep/ .el-card__body{
  padding-top: 5px;
  padding-bottom: 5px;
}
</style>