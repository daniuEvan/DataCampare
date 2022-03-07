<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <el-form ref="form" label-width="auto"           label-suffix=" : " :disabled="true" size="mini">
        <el-form-item label="任务名称">
          <el-input class="input-item" :placeholder="taskInfo['TaskName']"></el-input>
        </el-form-item>
        <el-form-item label="源端数据库连接">
          <el-input class="input-item" :placeholder="taskInfo['SourceDbLinkName']"></el-input>
        </el-form-item>
        <el-form-item label="目标端数据库连接">
          <el-input class="input-item" :placeholder="taskInfo['TargetDbLinkName']"></el-input>
        </el-form-item>
        <el-form-item label="配置表数据库连接">
          <el-input class="input-item" :placeholder="taskInfo['ConfigDbLinkName']"></el-input>
        </el-form-item>
        <el-form-item label="配置表owner">
          <el-input class="input-item" :placeholder="taskInfo['ConfigTableOwner']"></el-input>
        </el-form-item>
        <el-form-item label="配置表table">
          <el-input class="input-item" :placeholder="taskInfo['ConfigTableName']"></el-input>
        </el-form-item>
        <el-form-item label="结果表数据库连接">
          <el-input class="input-item" :placeholder="taskInfo['ResultDbLinkName']"></el-input>
        </el-form-item>
        <el-form-item label="结果表owner">
          <el-input class="input-item" :placeholder="taskInfo['ResultTableOwner']"></el-input>
        </el-form-item>
        <el-form-item label="结果表table">
          <el-input class="input-item" :placeholder="taskInfo['ResultTableName']"></el-input>
        </el-form-item>
        <el-form-item label="创建日期">
          <el-input class="input-item" :placeholder="taskInfo['CreatedAt']"></el-input>
        </el-form-item>
<!--        <el-form-item label="修改日期">-->
<!--          <el-input class="input-item" :placeholder="taskInfo['UpdatedAt']"></el-input>-->
<!--        </el-form-item>-->
      </el-form>
    </el-col>
  </el-row>

</template>

<script>
export default {
  name: "TaskMsg",
  props: {
    taskInfo: Object
  },
  // 监听测试
  // watch: {
  //   taskInfo: {
  //     handler(newValue, oldValue) {
  //       console.log(newValue);
  //       console.log(oldValue);
  //     },
  //     immediate: true, // 初次监听即执行
  //   },
  // },
  data() {
    return {}
  },
  mounted() {
    // console.log(this.taskInfo);
  }
}
</script>

<style scoped>
.input-item >>> .el-input__inner {
  -webkit-appearance: none;
  border-radius: 4px;
  border: none;
  cursor: auto;
  color: #969696;
  font-size: 14px;
}
.input-item >>> .el-input__inner:-moz-placeholder,
textarea:-moz-placeholder {
  color: #5b5b5b;
}

.input-item >>> .el-input__inner:-ms-input-placeholder,
textarea:-ms-input-placeholder {
  color: #5b5b5b;
}

.input-item >>> .el-input__inner::-webkit-input-placeholder,
textarea::-webkit-input-placeholder {
  color: #5b5b5b;
}

.el-form .el-form-item{
  margin-bottom: 10px;
}

</style>