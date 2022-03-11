<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset="4">
      <el-form
          :model="schedulerCreate"
          :rules="rules"
          ref="schedulerCreate"
          label-width="160px"
          label-suffix=" : "
          size="mini"
      >
        <el-form-item prop="SchedulerName" label="调度名称">
          <el-input class="input-item" v-model="schedulerCreate['SchedulerName']"></el-input>
        </el-form-item>
        <el-form-item prop="TaskId" label="任务">
          <el-select v-model="schedulerCreate['TaskId']" placeholder="请选择任务">
            <el-option v-for="item in taskInfo" :key="item['TaskName']+'task'" :label="item['TaskName']"
                       :value="item['ID']">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item prop="TaskSchedule" label="调度信息">
          <el-input class="input-item" v-model="schedulerCreate['TaskSchedule']"></el-input>
        </el-form-item>
        <el-form-item prop="SchedulerStatus" label="是否启用">
          <el-switch v-model="schedulerCreate['SchedulerStatus']"></el-switch>
        </el-form-item>
        <el-form-item prop="TaskConcurrent" label="调度并发数">
          <el-input class="input-item" v-model.number="schedulerCreate['TaskConcurrent']"></el-input>
        </el-form-item>


        <el-form-item label="描述信息" v-show="false">
          <el-input class="input-item" type="textarea" :placeholder="schedulerCreate['Desc']"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('schedulerCreate')">保存</el-button>
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
  name: "SchedulerCreate",
  props: ["getScheduler", "toDefaultShow", "dbLogo"],
  data() {
    return {
      dbTypeArr: [
        "vertica",
        "oracle",
        "mysql",
        "postgres",
      ],
      schedulerCreate: {},
      rules: {
        SchedulerName: [{required: true, message: '不能为空', trigger: "blur"}],
        TaskId: [{required: true, message: '不能为空', trigger: "blur"}],
        TaskSchedule: [{required: true, message: '不能为空', trigger: "blur"}],
        // SchedulerStatus: [{required: true, message: '不能为空', trigger: "blur"}],
        TaskConcurrent: [
          {required: true, message: '不能为空', trigger: "blur"},
          {type: "number", message: '请输入数字', trigger: "blur"},
          {type: "number", min: 1, message: '最小为1', trigger: "blur"},
          {type: "number", max: 1000, message: '最大为1000', trigger: "blur"}
        ],
      },
      taskInfo: [],
    };
  },
  methods: {
    submitForm(formName) {
      let _this = this
      this.$refs[formName].validate((valid) => {
        if (valid) {
          request.post("/scheduler/", _this.schedulerCreate)
              .then(function (response) {
                if (response.data.code !== 200) {
                  _this.$message.error({center: true, message: '调度创建失败' + response.data.msg,});
                  return null
                }
                _this.$message.success({center: true, message: '调度创建成功',});
                _this.getScheduler()
                _this.toDefaultShow()
                _this.schedulerCreate = {}
              })
              .catch(function (err) {
                console.log(err);
                _this.$message.error({center: true, message: '调度创建失败'});
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
    getTaskInfo() {
      this.taskInfo = JSON.parse(storageService.get(storageService.TASK_INFO_LIST))
    },
  },
  mounted() {
    this.getTaskInfo()
  }
}
</script>

<style scoped>

</style>