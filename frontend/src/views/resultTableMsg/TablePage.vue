<template>
  <el-table
      :data="tableData"
      border
      style="width: 100%">
    <el-table-column
        align="center"
        prop="date"
        label="OWNER"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="date"
        label="TABLENAME"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="date"
        label="SOURCE_NUM"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="date"
        label="TARGET_NUM"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="name"
        label="SOURCE_MAX"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="address"
        label="TARGET_MAX">
    </el-table-column>
    <el-table-column
        align="center"
        prop="address"
        label="CHECK_TIME">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "TablePage",
  props:["tableData"],
}
</script>

<style scoped>

</style>