<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <div v-if="schedulerInfo['SchedulerName']">
        <el-form ref="form" label-width="auto" label-suffix=" : "  size="mini">
          <el-form-item label="调度名称">
            <el-input class="input-item"  :disabled="true"  :placeholder="schedulerInfo['SchedulerName']"></el-input>
          </el-form-item>
          <el-form-item label="任务名称">
            <el-input class="input-item" :disabled="true"  :placeholder="schedulerInfo['TaskName']"></el-input>
          </el-form-item>
          <el-form-item label="调度信息">
            <el-input class="input-item" :disabled="true" :placeholder="schedulerInfo['TaskSchedule']"></el-input>
          </el-form-item>
          <el-form-item label="调度状态">
            <el-input class="input-item" :disabled="true" :placeholder="schedulerInfo['SchedulerStatus']?'已启用':'未启用'"></el-input>
          </el-form-item>
          <el-form-item label="调度并发数">
            <el-input class="input-item" :disabled="true" :placeholder="schedulerInfo['TaskConcurrent']"></el-input>
          </el-form-item>
          <el-form-item label="源表查询sql">
            <el-input class="input-item" type="textarea" :readonly="true" v-model="schedulerInfo['SourceTableQuerySQL']"></el-input>
          </el-form-item>
          <el-form-item label="目标表查询sql">
            <el-input class="input-item" type="textarea" :readonly="true" v-model="schedulerInfo['TargetTableQuerySQL']"></el-input>
          </el-form-item>
          <el-form-item label="配置表查询sql">
            <el-input class="input-item" type="textarea" :readonly="true" v-model="schedulerInfo['ConfigTableQuerySQL']"></el-input>
          </el-form-item>
          <el-form-item label="结果表插入sql">
            <el-input class="input-item" type="textarea"
                      :readonly="true" v-model="schedulerInfo['ResultTableInsertSQL']"></el-input>
          </el-form-item>
          <el-form-item label="结果表初始化sql">
            <el-input class="input-item" type="textarea" :readonly="true" v-model="schedulerInfo['ResultTableInitSQL']"></el-input>
          </el-form-item>

          <el-form-item label="描述信息" v-if="false">
            <el-input class="input-item" type="textarea"  :placeholder="schedulerInfo['Desc']"></el-input>
          </el-form-item>
        </el-form>
      </div>
    </el-col>
  </el-row>

</template>

<script>
export default {
  name: "SchedulerMsg",
  props: {
    schedulerInfo: Object,
    dbLogo: Object,
  },
  // 监听测试
  // watch: {
  //   schedulerInfo: {
  //     handler(newValue, oldValue) {
  //       console.log(newValue);
  //       console.log(oldValue);
  //     },
  //     immediate: true, // 初次监听即执行
  //   },
  // },
  data() {
    return {
      hasValue: false,
    }
  },
  mounted() {
    // console.log(this.schedulerInfo.length);
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

.el-form .el-form-item {
  margin-bottom: 10px;
}

.input-item >>> .el-input-group__prepend {
  /*background-color: #fff;*/
  border: none;
  padding: 0;
  margin: 0;
}

.input-db >>> .el-input__inner {
  padding-left: 4px !important;
}

.input-db-logo {
  width: 15px;
  margin: 4px 0 0 14px;
  /*display: inline;*/
  /*float: left;*/
}

</style>