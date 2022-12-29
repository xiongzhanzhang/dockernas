<template>
  <div>
    <div class="card">
      <div class="input_div">
        <div class="first_input">IP地址</div>
        <div>{{ networkData.ip }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">域名</div>
        <div>
          {{ networkData.domain }}
          <el-icon class="click_able" @click="showSetDomain"
            ><EditPen
          /></el-icon>
        </div>
      </div>
      <div class="input_div">
        <div class="first_input">HTTP网关</div>
        <div>
          <el-switch
            v-model="networkData.httpGatewayEnable"
            :loading="networkData.HttpGatewayLoading"
            @change="changeGatewayState"
          />
        </div>
      </div>
      <div class="input_div">
        <div class="first_input">开启HTTPS</div>
        <div><el-switch v-model="networkData.httpsEnable" /></div>
      </div>
      <div class="input_div">
        <div class="first_input">ssl证书路径</div>
        <div>
          {{ networkData.sslCertificatePath }}
          <el-icon><EditPen /></el-icon>
        </div>
      </div>

      <div class="center_div" style="padding-top: 30px; padding-bottom: 20px">
        <el-button
          type="primary"
          style="height: 40px; width: 250px; margin-right: 100px"
          >重启HTTP网关</el-button
        >
        <el-button
          type="success"
          style="height: 40px; width: 250px"
          @click="showCreateProxy"
          >添加HTTP转发</el-button
        >
      </div>
    </div>

    <div class="card" style="margin-top: 16px; min-height: 200px">
      <el-table
        :data="httpProxyConfigs"
        stripe
        style="width: 100%; font-size: 18px"
        :row-style="{ height: '50px' }"
        :cell-style="{ padding: '0px' }"
      >
        <el-table-column prop="hostName" label="主机名" min-width="20%" />
        <el-table-column label="访问链接" min-width="40%" #default="scope">
          <a target="_blank" :href="'http://' + scope.row.url">{{
            scope.row.url
          }}</a>
        </el-table-column>
        <el-table-column prop="instanceName" label="代理实例" min-width="20%" />
        <el-table-column prop="port" label="代理端口" min-width="10%" />
        <el-table-column label="操作" min-width="10%" #default="scope">
          <el-button
            size="small"
            type="danger"
            @click="delHttpProxyConfig(scope.row)"
            >Delete</el-button
          >
        </el-table-column>
      </el-table>
    </div>

    <el-dialog
      v-model="createProxyVisible"
      title="创建HTTP代理"
      style="min-height: 500px"
    >
      <div class="center_div">
        <div>
          <div class="input_div2">
            <div class="first_input">主机名</div>
            <div>
              <el-input
                class="w-50 m-2"
                style="width: 400px"
                size="large"
                v-model="curHostName"
              >
              </el-input>
            </div>
          </div>
          <div class="input_div2">
            <div class="first_input">实例</div>
            <div>
              <el-select
                collapse-tags
                size="large"
                style="width: 400px"
                @change="versionChange"
                v-model="curInstanceName"
              >
                <el-option
                  v-for="item in instancesPorts"
                  :key="item.instanceName"
                  :label="item.instanceName"
                  :value="item.instanceName"
                />
              </el-select>
            </div>
          </div>
          <div class="input_div2">
            <div class="first_input">端口</div>
            <div>
              <el-select
                collapse-tags
                size="large"
                style="width: 400px"
                v-model="curSelectPort"
              >
                <el-option
                  v-for="item in curPorts"
                  :key="item"
                  :label="item"
                  :value="item"
                />
              </el-select>
            </div>
          </div>
          <div class="center_div" style="margin-top: 50px">
            <el-button
              type="primary"
              style="height: 40px; width: 200px"
              @click="createHttpProxyConfig"
              >{{ $t("common.yes") }}</el-button
            >
          </div>
        </div>
      </div>
    </el-dialog>

    <el-dialog
      v-model="setDomainVisiable"
      title="设置域名"
      style="min-height: 300px; width: 600px"
    >
      <div class="center_div">
        <div>
          <div class="input_div2">
            <div>
              <el-input
                class="w-50 m-2"
                style="width: 350px"
                size="large"
                v-model="curDomain"
              >
              </el-input>
            </div>
          </div>
          <div class="center_div" style="margin-top: 50px">
            <el-button
              type="primary"
              style="height: 40px; width: 200px"
              @click="editDomain"
              >{{ $t("common.yes") }}</el-button
            >
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getNetworkInfo,
  getHttpProxyConfig,
  postHttpProxyConfig,
  delHttpProxyConfig,
  postDomain,
  startHttpGateway,
} from "../../api/host";
import { getInstanceHttpPort } from "../../api/instance";

export default {
  name: "network",
  data() {
    return {
      networkData: {},
      value1: true,
      httpProxyConfigs: [],
      createProxyVisible: false,
      instancesPorts: [],
      curPorts: [],
      curInstanceName: "",
      curSelectPort: "",
      curHostName: "",

      setDomainVisiable: false,
      curDomain: "",
    };
  },
  methods: {
    changeGatewayState() {
      console.log("changeGatewayState");
      if (this.networkData.httpGatewayEnable == true) {
        this.networkData.HttpGatewayLoading=true;
        startHttpGateway().then((response) => {
          this.flush();
        });
      }
    },
    showSetDomain() {
      this.curDomain = this.networkData.domain;
      this.setDomainVisiable = true;
    },
    showCreateProxy() {
      this.createProxyVisible = true;
    },
    editDomain() {
      postDomain(this.curDomain).then((response) => {
        this.setDomainVisiable = false;
        this.flush();
      });
    },
    versionChange() {
      for (var instancePort of this.instancesPorts) {
        if (instancePort.instanceName == this.curInstanceName) {
          this.curPorts = instancePort.ports;
        }
      }
    },
    createHttpProxyConfig() {
      postHttpProxyConfig(
        this.curHostName,
        this.curInstanceName,
        this.curSelectPort
      ).then((response) => {
        this.createProxyVisible = false;
        this.flush();
      });
    },
    delHttpProxyConfig(row) {
      delHttpProxyConfig(row.hostName).then((response) => {
        this.flush();
      });
    },
    flush() {
      getNetworkInfo().then((response) => {
        this.networkData = response.data;
        console.log(this.networkData);
        getHttpProxyConfig().then((response) => {
          this.httpProxyConfigs = response.data.list;
          for (var config of this.httpProxyConfigs) {
            config.url = config.hostName + "." + this.networkData.domain;
          }
          console.log(this.instancesPort);
        });
      });
      getInstanceHttpPort().then((response) => {
        this.instancesPorts = response.data.list;
        console.log(this.instancesPort);
      });
    },
  },
  mounted() {
    this.flush();
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

.input_div2 {
  display: flex;
  color: black !important;
  font-size: 18px;
  align-items: center;
  height: 80px;
}
</style>