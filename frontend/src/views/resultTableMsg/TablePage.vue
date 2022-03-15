<template>
  <el-table
      :data="tableData"
      border
      :row-style="rowStyle"
      style="width: 100%"
  >
    <el-table-column
        align="center"
        prop="owner"
        label="OWNER"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="tablename"
        label="TABLENAME"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="source_num"
        label="SOURCE_NUM"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="target_num"
        label="TARGET_NUM"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="source_max"
        label="SOURCE_MAX"
    >
    </el-table-column>
    <el-table-column
        align="center"
        prop="target_max"
        label="TARGET_MAX">
    </el-table-column>
    <el-table-column
        align="center"
        prop="check_time"
        label="CHECK_TIME">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  name: "TablePage",
  props: ["tableData"],
  methods: {
    rowStyle({row, rowIndex}) {
      if (row["source_num"] !== row["target_num"]) {
        return {
          color: "red"
        }
      }
    }
  }
}
</script>

<style scoped>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>