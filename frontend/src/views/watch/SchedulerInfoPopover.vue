<template>
  <el-popover
      placement="top-start"
      width="200"
      title="调度信息"
      trigger="hover"
      :content="popoverMsg['content']"
  >
    <ul style="list-style:none;margin-left: -20px" >
      <li><span>运行周期：</span>{{popoverMsg['TaskScheduler']}}</li>
      <li><span>任务并发：</span>{{popoverMsg['TaskConcurrent']}}</li>
      <li><span>调度启用：</span>{{popoverMsg['SchedulerEnable']?'已启用':'未启用'}}</li>
    </ul>
    <i slot="reference" class="el-icon-info" style="color: #b3d8ff"></i>
  </el-popover>
</template>

<script>
export default {
  name: "SchedulerInfoPopover",
  props: ["popoverMsg"],
}
</script>

<style scoped>

</style>