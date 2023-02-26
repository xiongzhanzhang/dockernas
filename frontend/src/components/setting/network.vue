<template>
  <div>
    <div class="card_style">
      <div class="input_div bottom_border">
        <div class="table_first_input">IPv4地址</div>
        <div class="table_second_input">{{ networkData.ip }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">IPv6地址</div>
        <div class="table_second_input">{{ networkData.ipv6 }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">域名</div>
        <div class="table_second_input">
          {{ networkData.domain }}
          <el-icon class="click_able" @click="showSetDomain"
            ><EditPen
          /></el-icon>
        </div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">HTTP网关</div>
        <div class="table_second_input">
          <el-switch
            v-model="networkData.httpGatewayEnable"
            :loading="networkData.HttpGatewayLoading"
            @change="changeGatewayState"
          />
        </div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">开启HTTPS</div>
        <div class="table_second_input">
          <el-switch
            v-model="networkData.httpsEnable"
            @change="changeHttpsState"
            :loading="changeHttpStateLoding"
          />
        </div>
      </div>

    </div>

    <div class="card_style" style="min-height: 600px">
      <div class="center_div" style="padding-top: 10px; padding-bottom: 10px;">
        <div style="flex-grow: 1"></div>
        <el-button
        type="primary"
          class="big_button"
          @click="showCreateProxy"
          >添加代理
        </el-button>
      </div>

      <el-table
        :data="httpProxyConfigs"
        stripe
        class="table_css"
        :row-style="{ height: '50px' }"
        :cell-style="{ padding: '0px' }"
      >
        <el-table-column prop="hostName" label="主机名" min-width="20%" />
        <el-table-column label="访问链接" min-width="40%" #default="scope">
          <a  v-if="!networkData.httpsEnable" target="_blank" :href="'http://' + scope.row.url">{{scope.row.url}}</a>
          <a  v-if="networkData.httpsEnable" target="_blank" :href="'https://' + scope.row.url">{{scope.row.url}}</a>
        </el-table-column>
        <el-table-column prop="instanceName" label="代理实例" min-width="20%" />
        <el-table-column prop="port" label="代理端口" min-width="20%" />
        <el-table-column label="操作" width="90px" #default="scope">
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
      class="small_dialog"
    >
      <div class="center_div">
        <div>
          <div class="input_div">
            <div class="first_input">实例</div>
            <div>
              <el-select
                collapse-tags
                size="large"
                class="big_input"
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
          <div class="input_div">
            <div class="first_input">主机名</div>
            <div>
              <el-input
                class="big_input"
                size="large"
                v-model="curHostName"
              >
              </el-input>
            </div>
          </div>
          <div class="input_div">
            <div class="first_input">端口</div>
            <div>
              <el-select
                collapse-tags
                filterable
                allow-create
                size="large"
                class="big_input"
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
      class="small_dialog"
    >
      <div class="center_div">
        <div>
          <div class="input_div">
            <div>
              <el-input
                class="big_input"
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
  stopHttpGateway,
  restartHttpGateway,
  enableHttps,
  disableHttps,
} from "../../api/host";
import { getInstanceHttpPort } from "../../api/instance";
import { getDfsDirs } from "../../api/filesystem";

export default {
  name: "network",
  data() {
    return {
      networkData: {},
      httpProxyConfigs: [],
      createProxyVisible: false,
      instancesPorts: [],
      curPorts: [],
      curInstanceName: "",
      curSelectPort: "",
      curHostName: "",

      setDomainVisiable: false,
      curDomain: "",

      restartGatewayLoading: false,
      changeHttpStateLoding: false,
    };
  },
  methods: {
    loadNode(node, resolve) {
      console.log(node.data);
      if (node.isLeaf) return resolve([]);
      var curPath = node.data.value;
      if (curPath == null) {
        resolve([
          {
            name: "/",
            value: "/",
            label: "/",
          },
        ]);
      } else {
        getDfsDirs(curPath).then((response) => {
          console.log(response.data);
          resolve(response.data.list);
        });
      }
    },
    tryRestartHttpGateway() {
      this.restartGatewayLoading = true;
      restartHttpGateway().then((response) => {
        this.restartGatewayLoading = false;
      });
    },
    changeHttpsState() {
      this.changeHttpStateLoding = true;
      if (this.networkData.httpsEnable == true) {
        enableHttps()
          .then((response) => {
            this.changeHttpStateLoding = false;
            this.flush();
          })
          .catch((error) => {
            this.changeHttpStateLoding = false;
            this.flush();
          });
      } else {
        disableHttps()
          .then((response) => {
            this.changeHttpStateLoding = false;
            this.flush();
          })
          .catch((error) => {
            this.changeHttpStateLoding = false;
            this.flush();
          });
      }
    },
    changeGatewayState() {
      this.networkData.HttpGatewayLoading = true;
      if (this.networkData.httpGatewayEnable == true) {
        startHttpGateway().then((response) => {
          this.flush();
        }).catch((error) => {
          this.flush();
        });
      }else{
        stopHttpGateway().then((response) => {
          this.flush();
        }).catch((error) => {
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
      this.curHostName=this.curInstanceName;
      for (var instancePort of this.instancesPorts) {
        if (instancePort.instanceName == this.curInstanceName) {
          this.curPorts = instancePort.ports;
          if(this.curPorts.length>0){
            this.curSelectPort=this.curPorts[0];
          }
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
        });
      });
      getInstanceHttpPort().then((response) => {
        this.instancesPorts = response.data.list;
      });
    },
  },
  mounted() {
    this.flush();
  },
};
</script>

<style>
@import "../../css/common.css";
@import "../../css/picture.css";
@import "../../css/menu.css";
@import "../../css/text.css";
@import "../../css/dialog.css";
</style>