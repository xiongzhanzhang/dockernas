<template>
  <div>
    <div>
      <div class="input_div">
        <div class="first_input">实例名</div>
        <div>{{ instance.name }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">注备</div>
        <div>{{ instance.summary }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">状态</div>
        <div v-if="instance.state == 0" style="color: gray">新建作业</div>
        <div v-if="instance.state == 1" style="color: yellow">创建失败</div>
        <div v-if="instance.state == 2" style="color: red">运行失败</div>
        <div v-if="instance.state == 3" style="color: green">运行中</div>
        <div v-if="instance.state == 4" style="color: gray">已停止</div>
      </div>
      <div class="input_div">
        <div class="first_input">应用名</div>
        <div>{{ instance.appName }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">版本</div>
        <div>{{ instance.version }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.portParams"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.envParams"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.dfsVolume"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.localVolume"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>
    </div>

    <div class="center_div" style="margin-top: 100px">
      <el-button
        type="primary"
        style="height: 40px; width: 200px"
        :disabled="instance.state == 3"
        @click="edit"
        >配置</el-button
      >
      <el-button
        type="success"
        style="height: 40px; width: 200px"
        :disabled="instance.state == 3"
        @click="start"
        >启动</el-button
      >
      <el-button
        type="warning"
        style="height: 40px; width: 200px"
        :disabled="instance.state != 3"
        @click="stop"
        >停止</el-button
      >
      <el-button
        type="danger"
        style="height: 40px; width: 200px"
        :disabled="instance.state == 3"
        @click="deleteInstancec"
        >删除</el-button
      >
    </div>

    <createInstance ref="createCard"></createInstance>
  </div>
</template>

<script>
import createInstance from "../createInstance.vue";
import {
  stopInstance,
  startInstance,
  deleteInstance,
} from "../../api/instance";

export default {
  name: "instanceBasicInfo",
  components: {
    createInstance,
  },
  data() {
    return {
      instanceParam: {},
      instance: {},
    };
  },
  methods: {
    initData(instance) {
      this.instance = instance;
      this.instanceParam = JSON.parse(this.instance.instanceParamStr);

      this.$refs.createCard.setAppName(this.instance.appName);
      this.$refs.createCard.setEditMode();
      this.$refs.createCard.setParams(
        this.instanceParam.name,
        this.instanceParam.summary,
        this.instanceParam.version,
        this.instanceParam.portParams,
        this.instanceParam.dfsVolume,
        this.instanceParam.envParams,
        this.instanceParam.localVolume
      );
    },
    stop() {
      stopInstance(this.instance.name).then((response) => {
        location.reload();
      });
    },
    start() {
      startInstance(this.instance.name).then((response) => {
        location.reload();
      });
    },
    deleteInstancec() {
      deleteInstance(this.instance.name).then((response) => {
        this.$router.push("/index/instances");
      });
    },
    edit() {
      this.$refs.createCard.showDialog();
    },
  },
  mounted() {},
};
</script>

<style scoped>
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