<template>
  <div>
    <div>
      <div class="input_div">
        <div class="first_input">实例名</div>
        <div>{{ instance.name }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">注备</div>
        <div>{{ instance.summury }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">状态</div>
        <div v-if="instance.state == 0" style="color:gray">新建作业</div>
        <div v-if="instance.state == 1" style="color:yellow">创建失败</div>
        <div v-if="instance.state == 2" style="color:red">运行失败</div>
        <div v-if="instance.state == 3" style="color:green">运行中</div>
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
      <el-button type="primary" style="height: 40px; width: 200px">配置</el-button>
      <el-button type="success" style="height: 40px; width: 200px">启动</el-button>
      <el-button type="warning" style="height: 40px; width: 200px">停止</el-button>
      <el-button type="danger" style="height: 40px; width: 200px">删除</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "instanceBasicInfo",
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