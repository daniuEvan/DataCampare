<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <el-form
          :model=" taskEditor"
          :rules="rules"
          ref="taskEditor"
          label-width="160px"
          label-suffix=" : " size="mini"
      >
        <el-form-item prop="TaskName" label="任务名称">
          <el-input class="input-item" v-model="taskEditor['TaskName']"></el-input>
        </el-form-item>
        <el-form-item prop="SourceDBLinkId" label="源端数据库连接">
          <el-select v-model="taskEditor['SourceDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'source'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="TargetDBLinkId" label="目标端数据库连接">
          <el-select v-model="taskEditor['TargetDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'target'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="ConfigDBLinkId" label="配置表数据库连接">
          <el-select v-model="taskEditor['ConfigDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'config'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="ConfigTableOwner" label="配置表owner">
          <el-input class="input-item" v-model="taskEditor['ConfigTableOwner']"></el-input>
        </el-form-item>
        <el-form-item prop="ConfigTableName" label="配置表table">
          <el-input class="input-item" v-model="taskEditor['ConfigTableName']"></el-input>
        </el-form-item>
        <el-form-item prop="ResultDBLinkId" label="结果表数据库连接">
          <el-select v-model="taskEditor['ResultDBLinkId']" placeholder="请选择数据库连接">
            <el-option v-for="item in dbLinkInfo" :key="item['LinkName']+'result'" :label="item['LinkName']"
                       :value="item['ID']">
              <span style="float: left;margin-right: 2px"><img style="width: 10px" class="input-db-logo"
                                                               :src="dbLogo[item['DBType']]" alt=""></span>
              <span style="float: left; color: #8492a6; font-size: 13px">{{ item['LinkName'] }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="ResultTableOwner" label="结果表owner">
          <el-input class="input-item" v-model="taskEditor['ResultTableOwner']"></el-input>
        </el-form-item>
        <el-form-item prop="ResultTableName" label="结果表table">
          <el-input class="input-item" v-model="taskEditor['ResultTableName']"></el-input>
        </el-form-item>
        <el-form-item label="描述信息" v-show="false">
          <el-input class="input-item" type="textarea" :placeholder="taskEditor['Desc']"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('taskEditor')">保存</el-button>
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
  name: "TaskEdit",
  props: ["taskEditor", "getTask", "toDefaultShow", "dbLogo"],
  data() {
    return {
      dbTypeArr: [
        "vertica",
        "oracle",
        "mysql",
        "postgres",
      ],
      rules: {
        TaskName: [{required: true, message: '不能为空', trigger: "blur"}],
        SourceDBLinkId: [{required: true, message: '不能为空', trigger: "blur"}],
        TargetDBLinkId: [{required: true, message: '不能为空', trigger: "blur"}],
        ConfigDBLinkId: [{required: true, message: '不能为空', trigger: "blur"}],
        ConfigTableOwner: [{required: true, message: '不能为空', trigger: "blur"}],
        ConfigTableName: [{required: true, message: '不能为空', trigger: "blur"}],
        ResultDBLinkId: [{required: true, message: '不能为空', trigger: "blur"}],
        ResultTableOwner: [{required: true, message: '不能为空', trigger: "blur"}],
        ResultTableName: [{required: true, message: '不能为空', trigger: "blur"}],
      },
      dbLinkInfo: [],
    };
  },
  methods: {
    submitForm(formName) {
      let _this = this
      this.$refs[formName].validate((valid) => {
        if (valid) {
          request.put("/task/", _this.taskEditor)
              .then(function (response) {
                if (response.data.code !== 200) {
                  _this.$message.error({message: '任务更新失败' + response.data.msg, center: true});
                  return null
                }
                _this.$message.success({message: '任务更新成功', center: true});
                _this.getTask()
                _this.toDefaultShow()
              })
              .catch(function (err) {
                console.log(err);
                _this.$message.error({message: '任务更新失败', center: true});
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