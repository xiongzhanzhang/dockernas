<template>
  <div>
    <div>
      <div class="input_div">
        <div class="first_input">实例名</div>
        <div>{{ instance.name }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">创建时间</div>
        <div>{{ new Date(instance.createTime).toLocaleString() }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">注备</div>
        <div>{{ instance.summary }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">状态</div>
        <div v-if="instance.state == 0" style="color: green">
          拉取镜像 {{ instance.imagePullState }}
        </div>
        <div v-if="instance.state == 1" style="color: yellow">创建失败</div>
        <div v-if="instance.state == 2" style="color: red">运行失败</div>
        <div v-if="instance.state == 3" style="color: green">运行中</div>
        <div v-if="instance.state == 4" style="color: gray">已停止</div>
        <div v-if="instance.state == 5" style="color: red">拉取失败</div>
      </div>
      <div class="input_div">
        <div class="first_input">应用名</div>
        <div v-if="instanceParam.appUrl != ''">
          <a target="_blank" :href="instanceParam.appUrl">{{ instance.appName }}</a>
        </div>
        <div v-if="instanceParam.appUrl == ''">{{instance.appName }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">仅本地访问</div>
        <div>
          <el-switch v-model="instanceParam.hostOnly" disabled></el-switch>
        </div>
      </div>
      <div class="input_div">
        <div class="first_input">版本</div>
        <div>{{ instance.version }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.portParams"
        v-show="param.hide == false"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div v-if="param.protocol == 'http'">
          <a target="_blank" :href="getInstanceWebUrl(instance.name, param.value)">{{ param.value }}</a>
        </div>
        <div v-if="param.protocol != 'http'">{{ param.value }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.envParams"
        v-show="param.hide == false && param.passwd != true"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.otherParams"
        v-show="param.hide == false && param.passwd != true"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>

      <div
        class="input_div"
        v-for="param in instanceParam.dfsVolume"
        v-show="param.hide == false"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div>

      <!-- <div
        class="input_div"
        v-for="param in instanceParam.localVolume"
        :key="param.prompt"
      >
        <div class="first_input">{{ param.prompt }}</div>
        <div>{{ param.value }}</div>
      </div> -->
    </div>

    <div class="center_div" style="margin-top: 100px">
      <el-button type="primary" style="height: 40px; width: 200px" @click="edit"
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
        @click="showDelete"
        >删除</el-button
      >
    </div>

    <createInstance ref="createCard"></createInstance>

    <el-dialog v-model="showDeleteDialog" title="确定删除?">
      <div class="center_div">
        <el-button
          type="primary"
          style="height: 40px; width: 200px"
          :disabled="cancelBtnDisable"
          @click="showDeleteDialog = false"
          >取消</el-button
        >
        <el-button
          type="danger"
          style="height: 40px; width: 200px"
          :loading="delBtnLoading"
          @click="requestDeleteInstance"
          >确认</el-button
        >
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getInstanceWebUrl } from "../../utils/url";
import createInstance from "../createInstance.vue";
import {
  stopInstance,
  startInstance,
  deleteInstance,
  getInstance,
} from "../../api/instance";

export default {
  name: "instanceBasicInfo",
  props: ["name"],
  components: {
    createInstance,
  },
  data() {
    return {
      instanceParam: {},
      instance: {},
      showDeleteDialog: false,

      delBtnLoading: false,
      cancelBtnDisable: false,
    };
  },
  methods: {
    getInstanceWebUrl,
    initData(instance) {
      this.instance = instance;
      this.instanceParam = JSON.parse(this.instance.instanceParamStr);

      this.$refs.createCard.setAppName(this.instance.appName);
      this.$refs.createCard.setEditMode();
      this.$refs.createCard.setParams(this.instanceParam);
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
    showDelete() {
      this.showDeleteDialog = true;
    },
    requestDeleteInstance() {
      (this.delBtnLoading = true),
        (this.cancelBtnDisable = true),
        deleteInstance(this.instance.name)
          .then((response) => {
            this.$router.push("/index/instances");
          })
          .catch((error) => {
            this.delBtnLoading = false;
            this.cancelBtnDisable = false;
          });
    },
    edit() {
      this.$refs.createCard.showDialog();
    },
  },
  mounted() {
    getInstance(this.name).then((response) => {
      this.initData(response.data);
    });
  },
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