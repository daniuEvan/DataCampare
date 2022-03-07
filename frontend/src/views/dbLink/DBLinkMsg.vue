<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <div v-if="dbLinkInfo['LinkName']">
        <el-form ref="form" label-width="auto"           label-suffix=" : " :disabled="true" size="mini">
          <el-form-item label="连接名称">
            <el-input class="input-item" :placeholder="dbLinkInfo['LinkName']"></el-input>
          </el-form-item>
          <el-form-item label="数据库类型">
            <el-input class="input-item" :placeholder="dbLinkInfo['DBType']"></el-input>
          </el-form-item>
          <el-form-item label="主机地址">
            <el-input class="input-item" :placeholder="dbLinkInfo['DBHost']"></el-input>
          </el-form-item>
          <el-form-item label="端口">
            <el-input class="input-item" :placeholder="dbLinkInfo['DBPort']"></el-input>
          </el-form-item>
          <el-form-item label="数据库">
            <el-input class="input-item" :placeholder="dbLinkInfo['DBName']"></el-input>
          </el-form-item>
          <el-form-item label="用户名">
            <el-input class="input-item" :placeholder="dbLinkInfo['DBUsername']"></el-input>
          </el-form-item>
          <el-form-item label="创建日期">
            <el-input class="input-item" :placeholder="dbLinkInfo['CreatedAt']"></el-input>
          </el-form-item>
          <!--        <el-form-item label="修改日期">-->
          <!--          <el-input class="input-item" :placeholder="dbLinkInfo['UpdatedAt']"></el-input>-->
          <!--        </el-form-item>-->
        </el-form>

      </div>
    </el-col>
  </el-row>

</template>

<script>
export default {
  name: "DBLinkMsg",
  props: {
    dbLinkInfo: Object
  },
  // 监听测试
  // watch: {
  //   dbLinkInfo: {
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
    // console.log(this.dbLinkInfo);
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