<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <el-form
          :model="dbLinkInfo"
          :rules="rules"
          ref="dbLinkInfo"
          label-width="auto"
          label-suffix=":"
          size="mini"
      >
        <el-form-item required prop="LinkName" label="连接名称">
          <el-input class="input-item" v-model="dbLinkInfo['LinkName']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBType" label="数据库类型">
          <el-select v-model="dbLinkInfo['DBType']" placeholder="请选择数据库类型">
            <el-option v-for="item in dbTypeArr" :key="item" :label="item"
                       :value="item"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item required prop="DBHost" label="主机地址">
          <el-input class="input-item" v-model="dbLinkInfo['DBHost']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBPort" label="端口">
          <el-input class="input-item" v-model="dbLinkInfo['DBPort']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBName" label="数据库">
          <el-input class="input-item" v-model="dbLinkInfo['DBName']"></el-input>
        </el-form-item>
        <el-form-item required prop="DBUsername" label="用户名">
          <el-input class="input-item" v-model="dbLinkInfo['DBUsername']"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="DBPassword">
          <el-input type="password" class="input-item" v-model="dbLinkInfo['DBPassword']"></el-input>
        </el-form-item>
        <el-form-item>
          <!--          <el-button type="primary" @click="testForm('ruleForm')">测试连接</el-button>-->
          <el-button type="primary" @click="submitForm('dbLinkInfo')">保存</el-button>
          <el-button @click="cancel()">取消</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>

</template>

<script>
import request from "@/utils/request";

export default {
  name: "DBLinkForm",
  props: ["dbLinkInfo", "dbLinkInfoMsgShow", "dbLinkInfoFormShow"],
  data() {
    return {
      dbTypeArr: [
        "vertica",
        "oracle",
        "mysql",
        "postgres",
      ],
      rules: {
        mustRequired: [
          {required: true, message: '不能为空', trigger: "blur"}
        ],
        DBPassword: [
          {required: true, message: '密码不能为空', trigger: "blur"},
          {min: 6, max: 18, message: '长度在 6 到 18 个字符', trigger: "blur"}
        ]
      }
    };
  },
  methods: {
    toDefaultShow(ok) {
      if (!ok) {
        this.dbLinkInfoFormShow = true
        this.dbLinkInfoMsgShow = false
        return null
      }
      this.dbLinkInfoFormShow = false
      this.dbLinkInfoMsgShow = true
      return null
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let _this = this
          request.put("/db_link/", _this.dbLinkInfo)
              .then(function (response) {
                if (response.data.code !== 200) {
                  _this.$message.error({message: '更新失败' + response.data.msg,});
                  return null
                }
                _this.$message.success({message: '更新成功',});
                _this.toDefaultShow(true)
                // _this.toDefaultShow()
              })
              .catch(function (err) {
                console.log(err);
                _this.$message.error({message: '更新失败'});
              })
        } else {
          this.$message.error("表单校验不通过")
          return false;
        }
      });
    },
    cancel() {
      this.toDefaultShow()
    }
  }
}
</script>

<style scoped>

</style>