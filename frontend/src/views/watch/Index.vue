<template>
  <el-row :gutter="24" class="watch-row">
    <el-col :span="16" :offset="4" class="watch-col">
      <el-tabs type="border-card" style="border:0.5px solid #DCDFE6;box-shadow:none" @tab-click="handleClick">
        <el-tab-pane label="所有调度">
          <el-col :span="22" :offset="1">
            <el-card v-if=" schedulerAllItems.length === 0 " class="box-card">
              <div class="text item">
                <i class="el-icon-coffee-cup" style="margin-right: 10px"></i> 空空如也
              </div>
            </el-card>
            <el-card v-for="o in schedulerAllItems" :key="o" class="box-card"
                     :style=" (o['SchedulerEnable'] && o['SchedulerStatus'] ? 'box-shadow: 0 2px 20px 0 rgb(1 180 55 / 30%);border: 1px solid #f2fff3;':'')+
                      (o['SchedulerEnable'] && !o['SchedulerStatus'] ? 'box-shadow: 0 2px 20px 0 rgb(250 0 0 / 20%);border: 1px solid #ffdfdd;':'')"
            >
              <div class="text item">
                <SchedulerInfoPopover :popover-msg="o"></SchedulerInfoPopover>
                {{ o["SchedulerName"] }}
                <MsgPopover style="float: right" v-if="o['SchedulerEnable'] && !o['SchedulerStatus']"
                            :popoverMsg="{'title':'错误信息','content':o['ErrorMsg']}"></MsgPopover>
              </div>
            </el-card>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="启动成功调度">
          <el-col :span="22" :offset="1">
            <el-card v-if=" schedulerSuccessItems.length === 0 " class="box-card">
              <div class="text item">
                <i class="el-icon-coffee-cup" style="margin-right: 10px"></i> 空空如也
              </div>
            </el-card>
            <el-card v-for="o in schedulerSuccessItems" :key="o" class="box-card box-card-success">
              <div class="text item">
                <SchedulerInfoPopover :popover-msg="o"></SchedulerInfoPopover>
                {{ o["SchedulerName"] }}
              </div>
            </el-card>
          </el-col>

        </el-tab-pane>
        <el-tab-pane label="启动失败调度">
          <el-col :span="22" :offset="1">
            <el-card v-if=" schedulerFailItems.length === 0 " class="box-card">
              <div class="text item">
                <i class="el-icon-coffee-cup" style="margin-right: 10px"></i> 空空如也
              </div>
            </el-card>
            <el-card v-for="o in schedulerFailItems" :key="o" class="box-card box-card-fail">
              <div class="text item">
                <SchedulerInfoPopover :popover-msg="o"></SchedulerInfoPopover>
                {{ o["SchedulerName"] }}
                <MsgPopover style="float: right" :popoverMsg="{'title':'错误信息','content':o['ErrorMsg']}"></MsgPopover>
              </div>
            </el-card>
          </el-col>
        </el-tab-pane>
        <el-tab-pane label="未启用调度">
          <el-col :span="22" :offset="1">
            <el-card v-if=" schedulerBanItems.length === 0 " class="box-card">
              <div class="text item">
                <i class="el-icon-coffee-cup" style="margin-right: 10px"></i> 空空如也
              </div>
            </el-card>
            <el-card v-for="o in schedulerBanItems" :key="o" class="box-card">
              <div class="text item">
                <SchedulerInfoPopover :popover-msg="o"></SchedulerInfoPopover>
                {{ o["SchedulerName"] }}
              </div>
            </el-card>
          </el-col>

        </el-tab-pane>
      </el-tabs>
    </el-col>
  </el-row>


</template>

<script>
import request from "@/utils/request";
import storageService from "@/service/storageService";
import MsgPopover from "@/components/MsgPopover";
import SchedulerInfoPopover from "@/views/watch/SchedulerInfoPopover";

export default {
  name: "Index",
  components: {
    MsgPopover,SchedulerInfoPopover
  },
  data() {
    return {
      schedulerAllItems: [],
      schedulerSuccessItems: [],
      schedulerFailItems: [],
      schedulerBanItems: [],
    }
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    },
    getSchedulerStatus() {
      let _this = this
      request.get('/scheduler/watch_scheduler/')
          .then(function (response) {
            if (response.data.code !== 200) {
              _this.$message.error(response.data.msg)
              return null
            }
            _this.chedulerStatus = response.data.data
            _this.schedulerAllItems = _this.chedulerStatus["schedulerAllItems"]
            _this.schedulerSuccessItems = _this.chedulerStatus["schedulerSuccessItems"]
            _this.schedulerFailItems = _this.chedulerStatus["schedulerFailItems"]
            _this.schedulerBanItems = _this.chedulerStatus["schedulerBanItems"]
            // 存入浏览器
            storageService.set(storageService.SCHEDULER_STATUS, JSON.stringify(response.data.data))
          })
          .catch(function (err) {
            _this.$message.error(err)
          })
    }
  },
  mounted() {
    this.getSchedulerStatus()
  }
}
</script>

<style scoped>
.watch-col, .watch-row {
  /*background-color: #e5e5e5;*/
  /*width: 100%;*/
  height: 100%;
}

.box-card {
  margin: 10px 0 20px 0;
}

.box-card-success {
  box-shadow: 0 2px 20px 0 rgb(1 180 55 / 30%);
  border: 1px solid #f2fff3;
}

.box-card-fail {
  box-shadow: 0 2px 20px 0 rgb(250 0 0 / 20%);
  border: 1px solid #ffdfdd;
}

</style>