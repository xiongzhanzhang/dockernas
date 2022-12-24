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
              v-model="instanceParam.name"
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
              v-model="instanceParam.summary"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
            >
            </el-input>
          </div>
        </div>
        <div class="input_div">
          <div class="first_input">{{ $t("store.appName") }}</div>
          <div>
            {{ instanceParam.appName }}
            <el-tooltip effect="dark" placement="bottom">
              <el-icon><InfoFilled /></el-icon>
              <template #content
                ><div style="width: 300px">{{ app.summary }}</div></template
              >
            </el-tooltip>
          </div>
        </div>
        <div class="input_div">
          <div class="first_input">{{ $t("store.appVersion") }}</div>
          <div>
            <el-select
              v-model="instanceParam.version"
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

        <div
          class="input_div"
          v-for="param in instanceParam.portParams"
          v-show="param.hide == false"
          :key="param.prompt"
        >
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
        <div
          class="input_div"
          v-for="param in instanceParam.dfsVolume"
          v-show="param.hide == false"
          :key="param.prompt"
        >
          <div class="first_input">{{ param.prompt }}</div>
          <!-- <div>
            <el-input
              v-model="param.value"
              class="w-50 m-2"
              style="width: 400px"
              size="large"
            ></el-input>
          </div> -->
          <el-tree-select
            class="w-50 m-2"
            style="width: 400px"
            size="large"
            v-model="param.value"
            check-strictly
            lazy
            :load="loadNode"
          >
            <template #default="{ data: { name } }">{{ name }}</template>
          </el-tree-select>
        </div>
        <div
          class="input_div"
          v-for="param in instanceParam.envParams"
          v-show="param.hide == false"
          :key="param.prompt"
        >
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
            :loading="btnLoading"
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
import { getDfsDirs } from "../api/filesystem";

export default {
  name: "createInstance",
  data() {
    return {
      btnLoading: false,

      app: {},
      editMode: false,

      instanceParam: {
        name: "",
        summary: "",
        appName: "",
        imageUrl: "",
        version: "",
        portParams: "",
        envParams: "",
        localVolume: "",
        dfsVolume: "",
        iconUrl: "",
        privileged: false,
      },

      dialogTableVisible: false,
      selectVersion: "",
      appVersions: [],
      title: this.$t("store.instanceApp"),
    };
  },
  methods: {
    loadNode(node, resolve) {
      console.log(node.data);
      if (node.isLeaf) return resolve([]);
      var curPath = node.data.value;
      if (curPath == null) {
        resolve([{
          name:"/",
          value:"/",
          label:"/",
        }])
      } else {
        getDfsDirs(curPath).then((response) => {
          console.log(response.data);
          resolve(response.data.list);
        });
      }
    },
    showDialog() {
      this.appVersions = [];
      for (let d of this.app.dockerVersions) {
        this.appVersions.push(d.version);
      }
      this.dialogTableVisible = true;
    },
    setParams(instanceParam) {
      this.instanceParam.name = instanceParam.name;
      this.instanceParam.summary = instanceParam.summary;
      this.instanceParam.version = instanceParam.version;
      this.instanceParam.portParams = instanceParam.portParams;
      this.instanceParam.dfsVolume = instanceParam.dfsVolume;
      this.instanceParam.envParams = instanceParam.envParams;
      this.instanceParam.localVolume = instanceParam.localVolume;
      this.instanceParam.imageUrl = instanceParam.imageUrl;
      this.instanceParam.privileged = instanceParam.privileged;
    },
    setApp(app) {
      this.app = app;
      this.instanceParam.iconUrl = app.iconUrl;
      this.instanceParam.appName = app.name;
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
        if (d.version == this.instanceParam.version) {
          this.instanceParam.imageUrl = d.imageUrl;
          this.instanceParam.privileged = d.privileged;
          this.instanceParam.portParams = JSON.parse(
            JSON.stringify(d.portParams)
          );
          this.instanceParam.dfsVolume = JSON.parse(
            JSON.stringify(d.dfsVolume)
          );
          this.instanceParam.envParams = JSON.parse(
            JSON.stringify(d.envParams)
          );
          this.instanceParam.localVolume = JSON.parse(
            JSON.stringify(d.localVolume)
          );
        }
      }
    },
    createApp() {
      this.btnLoading = true;
      if (this.editMode != true) {
        newInstance(this.instanceParam)
          .then((response) => {
            console.log(response);
            this.dialogTableVisible = false;
            this.btnLoading = false;
          })
          .catch((error) => {
            this.btnLoading = false;
          });
      } else {
        editInstance(
          this.instanceParam.name,
          JSON.stringify(this.instanceParam)
        )
          .then((response) => {
            console.log(response);
            this.dialogTableVisible = false;
            this.btnLoading = false;
            location.reload();
          })
          .catch((error) => {
            this.btnLoading = false;
          });
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
  width: 200px;
  text-align: right;
  margin-right: 30px;
}
</style>