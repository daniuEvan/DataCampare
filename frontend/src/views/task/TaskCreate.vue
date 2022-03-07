<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <el-form
          :model="taskCreate"
          :rules="rules"
          ref="taskCreate"
          label-width="160px"
          label-suffix=" : "
          size="mini"
      >
        <el-form-item prop="TaskName" label="任务名称">
          <el-input class="input-item" v-model="taskCreate['TaskName']"></el-input>
        </el-form-item>
        <el-form-item prop="SourceDBLinkId" label="源端数据库连接">
          <el-select v-model="taskCreate['SourceDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'source'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="TargetDBLinkId" label="目标端数据库连接">
          <el-select v-model="taskCreate['TargetDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'target'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="ConfigDBLinkId" label="配置表数据库连接">
          <el-select v-model="taskCreate['ConfigDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'config'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="ConfigTableOwner" label="配置表owner">
          <el-input class="input-item" v-model="taskCreate['ConfigTableOwner']"></el-input>
        </el-form-item>
        <el-form-item prop="ConfigTableName" label="配置表table">
          <el-input class="input-item" v-model="taskCreate['ConfigTableName']"></el-input>
        </el-form-item>
        <el-form-item prop="ResultDBLinkId" label="结果表数据库连接">
          <el-select v-model="taskCreate['ResultDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'result'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="ResultTableOwner" label="结果表owner">
          <el-input class="input-item" v-model="taskCreate['ResultTableOwner']"></el-input>
        </el-form-item>
        <el-form-item prop="ResultTableName" label="结果表table">
          <el-input class="input-item" v-model="taskCreate['ResultTableName']"></el-input>
        </el-form-item>
        <el-form-item label="描述信息" v-show="false">
          <el-input class="input-item" type="textarea" :placeholder="taskCreate['Desc']"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('taskCreate')">保存</el-button>
          <el-button @click="cancel()">取消</el-button>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>

</template>

<script>
import request from "@/utils/request";
import storageService from "@/service/storageService";

export default {
  name: "TaskCreate",
  props: ["getTask", "toDefaultShow", "dbLogo"],
  data() {
    return {
      dbTypeArr: [
        "vertica",
        "oracle",
        "mysql",
        "postgres",
      ],
      taskCreate: {},
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
      },
      dbLinkInfo: [],
    };
  },
  methods: {
    submitForm(formName) {
      let _this = this
      this.$refs[formName].validate((valid) => {
        if (valid) {
          request.post("/task/", _this.taskCreate)
              .then(function (response) {
                if (response.data.code !== 200) {
                  _this.$message.error({center: true, message: '任务创建失败' + response.data.msg,});
                  return null
                }
                _this.$message.success({center: true, message: '任务创建成功',});
                _this.getTask()
                _this.toDefaultShow()
              })
              .catch(function (err) {
                console.log(err);
                _this.$message.error({center: true, message: '任务创建失败'});
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
    getDBLinkInfo() {
      this.dbLinkInfo = JSON.parse(storageService.get(storageService.DB_LINK_LIST))
    },
  },
  mounted() {
    this.getDBLinkInfo()
  }
}
</script>

<style scoped>

</style>