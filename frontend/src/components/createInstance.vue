<template>
  <el-dialog
    v-model="dialogTableVisible"
    :title="title"
    style="min-height: 500px"
  >
    <div class="center_div">
      <div>
        <div class="input_div">
          <div class="first_input">{{ $t("store.instanceName") }}</div>
          <div>
            <el-input
              v-model="instanceName"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
              :disabled="editMode == true"
            >
            </el-input>
          </div>
        </div>
        <div class="input_div">
          <div class="first_input">{{ $t("store.instanceSummary") }}</div>
          <div>
            <el-input
              v-model="summary"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
            >
            </el-input>
          </div>
        </div>
        <div class="input_div">
          <div class="first_input">{{ $t("store.appName") }}</div>
          <div>{{ appName }}</div>
        </div>
        <div class="input_div">
          <div class="first_input">{{ $t("store.appVersion") }}</div>
          <div>
            <el-select
              v-model="selectVersion"
              collapse-tags
              size="large"
              style="width: 400px"
              @change="versionChange"
            >
              <el-option
                v-for="item in appVersions"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select>
          </div>
        </div>

        <div class="input_div" v-for="param in portParams" :key="param.prompt">
          <div class="first_input">{{ param.prompt }}</div>
          <div>
            <el-input
              v-model="param.value"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
            ></el-input>
          </div>
        </div>
        <div class="input_div" v-for="param in dfsVolume" :key="param.prompt">
          <div class="first_input">{{ param.prompt }}</div>
          <div>
            <el-input
              v-model="param.value"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
            ></el-input>
          </div>
        </div>
        <div class="input_div" v-for="param in envParams" :key="param.prompt">
          <div class="first_input">{{ param.prompt }}</div>
          <div>
            <el-input
              v-model="param.value"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
            ></el-input>
          </div>
        </div>

        <div class="center_div" style="margin-top: 50px">
          <el-button
            type="primary"
            style="height: 40px; width: 200px"
            @click="createApp"
            >{{ $t("common.yes") }}</el-button
          >
        </div>
      </div>
    </div>
  </el-dialog>
</template>

<script>
import { newInstance, editInstance } from "../api/instance";
import { getAppsByName } from "../api/store";

export default {
  name: "createInstance",
  data() {
    return {
      app: {},
      editMode: false,

      instanceName: "",
      summary: "",
      appName: "",
      // appVersion: "",
      dialogTableVisible: false,
      selectVersion: "",
      appVersions: [],
      title: this.$t("store.instanceApp"),

      portParams: [],
      dfsVolume: [],
      envParams: [],
      localVolume: [],
    };
  },
  methods: {
    showDialog() {
      this.appName = this.app.name;
      this.appVersions = [];
      for (let d of this.app.dockerVersions) {
        this.appVersions.push(d.version);
      }

      this.dialogTableVisible = true;
    },
    setParams(
      name,
      summary,
      appVersion,
      portParams,
      dfsVolume,
      envParams,
      localVolume
    ) {
      this.instanceName = name;
      this.summary = summary;
      this.selectVersion = appVersion;
      this.portParams = portParams;
      this.dfsVolume = dfsVolume;
      this.envParams = envParams;
      this.localVolume = localVolume;
    },
    setApp(app) {
      this.app = app;
    },
    setAppName(name) {
      getAppsByName(name).then((response) => {
        this.setApp(response.data);
      });
    },
    setEditMode() {
      this.title = "编辑应用";
      this.editMode = true;
    },
    versionChange() {
      for (let d of this.app.dockerVersions) {
        if (d.version == this.selectVersion) {
          this.imageUrl = d.imageUrl;
          this.portParams = JSON.parse(JSON.stringify(d.portParams));
          this.dfsVolume = JSON.parse(JSON.stringify(d.dfsVolume));
          this.envParams = JSON.parse(JSON.stringify(d.envParams));
          this.localVolume = JSON.parse(JSON.stringify(d.localVolume));
        }
      }
    },
    createApp() {
      if (this.editMode != true) {
        newInstance(
          this.instanceName,
          this.summary,
          this.appName,
          this.imageUrl,
          this.selectVersion,
          this.portParams,
          this.envParams,
          this.localVolume,
          this.dfsVolume,
          this.app.iconUrl
        )
          .then((response) => {
            console.log(response);
            this.dialogTableVisible = false;
          })
          .catch((error) => {
            console.log(error);
          });
      } else {
        editInstance(this.instanceName, "");
      }
    },
  },
};
</script>

<style scoped>
@import "../css/common.css";
.input_div {
  display: flex;
  color: black !important;
  font-size: 16px;
  align-items: center;
  height: 50px;
}
.first_input {
  width: 100px;
  text-align: right;
  margin-right: 30px;
}
</style>