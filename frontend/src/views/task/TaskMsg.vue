<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <div v-if="taskInfo['TaskName']">
        <el-form ref="form" label-width="auto" label-suffix=" : " :disabled="true" size="mini">
          <el-form-item label="任务名称">
            <el-input class="input-item" :placeholder="taskInfo['TaskName']"></el-input>
          </el-form-item>
          <el-form-item label="源端数据库连接">
            <el-input class="input-item input-db" :placeholder="taskInfo['SourceDBLinkName']">
              <template slot="prepend">
                <img class="input-db-logo" :src="dbLogo[taskInfo['SourceDBType']]" alt="">
              </template>
            </el-input>
          </el-form-item>
          <el-form-item label="目标端数据库连接">
            <el-input class="input-item input-db" :placeholder="taskInfo['TargetDBLinkName']">
              <template slot="prepend">
                <img class="input-db-logo" :src="dbLogo[taskInfo['TargetDBType']]" alt="">
              </template>
            </el-input>
          </el-form-item>
          <el-form-item label="配置表数据库连接">
            <el-input class="input-item input-db" :placeholder="taskInfo['ConfigDBLinkName']">
              <template slot="prepend">
                <img class="input-db-logo" :src="dbLogo[taskInfo['ConfigDBType']]" alt="">
              </template>
            </el-input>
          </el-form-item>
          <el-form-item label="配置表owner">
            <el-input class="input-item" :placeholder="taskInfo['ConfigTableOwner']"></el-input>
          </el-form-item>
          <el-form-item label="配置表table">
            <el-input class="input-item" :placeholder="taskInfo['ConfigTableName']"></el-input>
          </el-form-item>
          <el-form-item label="结果表数据库连接">
            <el-input class="input-item input-db" :placeholder="taskInfo['ResultDBLinkName']">
              <template slot="prepend">
                <img class="input-db-logo" :src="dbLogo[taskInfo['ResultDBType']]" alt="">
              </template>
            </el-input>
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
          <el-form-item label="描述信息" v-show="false">
            <el-input class="input-item" type="textarea" :placeholder="taskInfo['Desc']"></el-input>
          </el-form-item>
        </el-form>
      </div>
    </el-col>
  </el-row>

</template>

<script>
export default {
  name: "TaskMsg",
  props: {
    taskInfo: Object,
    dbLogo: Object,
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
    return {
      hasValue: false
    }
  },
  mounted() {
    // console.log(this.taskInfo.length);
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