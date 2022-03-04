<template>
  <el-row :gutter="20">
    <el-col :span="14" :offset= "3" >
      <el-form ref="form" :model="form" label-width="80px" label-suffix=":" :disabled="true" size="mini">
        <el-form-item label="活动名称">
          <el-input class="input-item" v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item  label="活动区域">
          <el-input class="input-item" v-model="form.region" placeholder="请选择活动区域"></el-input>
        </el-form-item>
      </el-form>
    </el-col>
  </el-row>

</template>

<script>
export default {
  name: "DBLinkMsg",
  data() {
    return {
      form: {
        name: '',
        region: '',
        date1: '',
        date2: '',
        delivery: false,
        type: [],
        resource: '',
        desc: ''
      }
    }
  },
}
</script>

<style scoped>
.input-item >>> .el-input__inner {
  -webkit-appearance: none;
  border-radius: 4px;
  border: none;
  cursor: auto;
}
</style>