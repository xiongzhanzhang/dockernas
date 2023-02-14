<template>
  <div>
    <div class="card_style">
      <div class="input_div bottom_border">
        <div class="table_first_input">主机名</div>
        <div class="table_second_input">{{ hostData.hostname }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">系统</div>
        <div class="table_second_input">{{ hostData.platform }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">DockerNAS版本</div>
        <div class="table_second_input">{{ hostData.dockerNASVersion }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">cpu型号</div>
        <div class="table_second_input">{{ hostData.modelName }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">内存大小</div>
        <div class="table_second_input">{{ (hostData.memSize / 1024 / 1024 / 1024).toFixed(2) }} GB</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">启动时间</div>
        <div class="table_second_input">{{ new Date(hostData.bootTime).toLocaleString() }}</div>
      </div>
    </div>
    <div class="card_style">
      <instanceMonitor name="dockernas"></instanceMonitor>
    </div>
    <div class="card_style">
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

<style>
@import "../../css/common.css";
@import "../../css/picture.css";
@import "../../css/menu.css";
@import "../../css/text.css";
</style>