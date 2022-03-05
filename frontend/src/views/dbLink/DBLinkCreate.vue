<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <el-form
          :model="dbLinkCreate"
          :rules="rules"
          ref="dbLinkCreate"
          label-width="120px"
          label-suffix=" : "
          size="mini"
      >
        <el-form-item prop="LinkName" label="连接名称">
          <el-input class="input-item" v-model="dbLinkCreate['LinkName']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBType" label="数据库类型">
          <el-select v-model="dbLinkCreate['DBType']" placeholder="请选择数据库类型">
            <el-option v-for="item in dbTypeArr" :key="item" :label="item" :value="item"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item required prop="DBHost" label="主机地址">
          <el-input class="input-item" v-model="dbLinkCreate['DBHost']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBPort" label="端口">
          <el-input class="input-item" v-model.number="dbLinkCreate['DBPort']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBName" label="数据库">
          <el-input class="input-item" v-model="dbLinkCreate['DBName']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBUsername" label="用户名">
          <el-input class="input-item" v-model="dbLinkCreate['DBUsername']"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="DBPassword">
          <el-input show-password class="input-item" v-model="dbLinkCreate['DBPassword']"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="testDBLink('dbLinkCreate')">测试</el-button>
          <el-button type="primary" @click="submitForm('dbLinkCreate')">保存</el-button>
          <el-button @click="cancel()">取消</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>

</template>

<script>
import request from "@/utils/request";

export default {
  name: "DBLinkCreate",
  props: ["getDBLink", "toDefaultShow"],
  data() {
    return {
      dbTypeArr: [
        "vertica",
        "oracle",
        "mysql",
        "postgres",
      ],
      dbLinkCreate: {},
      rules: {
        LinkName: [{required: true, message: '连接名称不能为空', trigger: "blur"}],
        DBType: [{required: true, message: '数据库类型不能为空', trigger: "blur"}],
        DBHost: [{required: true, message: '主机地址不能为空', trigger: "blur"}],
        DBPort: [{required: true, type: "number", message: '正确输入端口号', trigger: "blur"}],
        DBName: [{required: true, message: '数据库不能为空', trigger: "blur"}],
        DBUsername: [{required: true, message: '用户名不能为空', trigger: "blur"}],
        DBPassword: [
          {required: true, message: '密码不能为空', trigger: "blur"},
          {min: 6, max: 18, message: '长度在 6 到 18 个字符', trigger: "blur"}
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      let _this = this
      this.$refs[formName].validate((valid) => {
        if (valid) {
          request.post("/db_link/", _this.dbLinkCreate)
              .then(function (response) {
                if (response.data.code !== 200) {
                  _this.$message.error({center: true, message: '数据库连接创建失败' + response.data.msg,});
                  return null
                }
                _this.$message.success({center: true, message: '数据库连接创建成功',});
                _this.getDBLink()
                _this.toDefaultShow()
              })
              .catch(function (err) {
                console.log(err);
                _this.$message.error({center: true, message: '数据库连接创建'});
              })
        } else {
          this.$message.error("表单校验不通过")
          return false;
        }
      });
    },
    cancel() {
      this.toDefaultShow()
    },
    testDBLink(formName) {
      let _this = this
      this.$refs[formName].validate((valid) => {
        if (valid) {
          request.post("/db_link/ping", _this.dbLinkCreate)
              .then(function (response) {
                if (response.data.code !== 200) {
                  _this.$message.error({center: true, message: '测试连接失败' + response.data.msg,});
                  return null
                }
                _this.$message.success({center: true, message: '测试连接成功',});
              })
              .catch(function (err) {
                console.log(err);
                _this.$message.error({center: true, message: '测试连接失败'});
              })
        } else {
          this.$message.error("表单校验不通过")
          return false;
        }
      });
    }
  }
}
</script>

<style scoped>

</style>