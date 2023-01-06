<template>
  <el-table
    :data="events"
    stripe
    style="width: 100%; font-size: 18px"
    :row-style="{ height: '50px' }"
    :cell-style="{ padding: '0px' }"
  >
    <el-table-column prop="createTime" label="时间" width="250" />
    <el-table-column prop="eventType" label="事件类别" width="250" />
    <el-table-column prop="msg" label="具体信息" />
  </el-table>
</template>

<script>
import { getInstanceEvent } from "../../api/instance";

export default {
  name: "instanceEvent",
  props: ["name"],
  data() {
    return {
      events: [],
    };
  },
  methods: {
    getFineMsg(msg) {
      if (msg.length == 0) {
        return "成功";
      }
      return msg;
    },
    getEventTypeStr(eventType) {
      if (eventType == 0) return "创建实例";
      if (eventType == 1) return "停止实例";
      if (eventType == 2) return "启动实例";
      if (eventType == 3) return "配置实例";
      if (eventType == 4) return "删除实例";
      if (eventType == 5) return "重启实例";

      return "未知事件";
    },
    initData() {
      this.events=[]
      getInstanceEvent(this.name).then((response) => {
        for (var d of response.data.list) {
          this.events.push({
            createTime: new Date(d.createTime).toLocaleString(),
            eventType: this.getEventTypeStr(d.eventType),
            msg: this.getFineMsg(d.msg),
          });
        }
        console.log(this.events);
      });
    },
  },
  mounted(){
    this.initData();
  }
};
</script>

<style scoped>
</style>