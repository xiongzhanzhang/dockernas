<template>
  <div>
    <div class="card">
      <div class="input_div">
        <div class="first_input">主机名</div>
        <div>{{ hostData.hostname }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">系统</div>
        <div>{{ hostData.platform }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">DockerNAS版本</div>
        <div>{{ hostData.dockerNASVersion }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">cpu型号</div>
        <div>{{ hostData.modelName }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">内存大小</div>
        <div>{{ (hostData.memSize / 1024 / 1024 / 1024).toFixed(2) }} GB</div>
      </div>
      <div class="input_div">
        <div class="first_input">启动时间</div>
        <div>{{ new Date(hostData.bootTime).toLocaleString() }}</div>
      </div>
    </div>
    <div>
      <instanceMonitor name="dockernas"></instanceMonitor>
    </div>
    <div>
      <instanceMonitor></instanceMonitor>
    </div>
  </div>
</template>

<script>
import instanceMonitor from "../instance/instanceMonitor.vue";
import { getHostInfo } from "../../api/host";

export default {
  name: "host",
  components: {
    instanceMonitor,
  },
  data() {
    return {
      hostData: {},
    };
  },
  methods: {},
  mounted() {
    getHostInfo().then((response) => {
      this.hostData = response.data;
    });
  },
};
</script>

<style scoped>
.card {
  /* padding: 6px; */
  /* margin: 6px; */
  border-radius: 3px;
  background-color: white;
}

.input_div {
  display: flex;
  color: black !important;
  font-size: 18px;
  align-items: center;
  height: 80px;
  border-bottom: 1px solid rgb(222, 222, 222);
}
.first_input {
  width: 33%;
  padding-left: 30px;
  /* text-align: right; */
  margin-right: 80px;
}
</style>