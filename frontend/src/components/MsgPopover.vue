<template>
  <el-popover
      placement="top-start"
      :title="popoverMsg['title']"
      width="200"
      trigger="hover"
      :content="popoverMsg['content']">
    <i slot="reference" class="el-icon-chat-line-square"></i>
  </el-popover>
</template>

<script>
export default {
  name: "MsgPopover",
  props: ["popoverMsg"],
}
</script>

<style scoped>

</style>