<template>
  <div>
    <div class="card">
      <div class="input_div">
        <div class="first_input">IPv4地址</div>
        <div>{{ networkData.ip }}</div>
      </div>
      <div class="input_div">
        <div class="first_input">IPv6地址</div>
        <div>{{ networkData.ipv6 }}</div>
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
        <div>
          <el-switch
            v-model="networkData.httpsEnable"
            @change="changeHttpsState"
            :loading="changeHttpStateLoding"
          />
        </div>
      </div>
      <div class="input_div">
        <div class="first_input">ssl证书路径
          <el-tooltip effect="dark" placement="bottom">
              <el-icon><InfoFilled /></el-icon>
              <template #content
                ><div style="width: 300px">证书需要命名为example.com.key，example.com.key.crt或example.com.key.cer或example.com_bundle.crt。设置后需要重启网关</div></template
              >
            </el-tooltip>
        </div>
        <div>
          {{ networkData.sslCertificatePath }}
          <el-icon class="click_able" @click="showSetCaFilePath"
            ><EditPen
          /></el-icon>
        </div>
      </div>

    </div>

    <div class="card" style="margin-top: 16px; min-height: 300px">
      <div class="center_div" style="padding-top: 10px; padding-bottom: 10px; margin-right: 20px;">
        <div style="flex-grow: 1"></div>
        <el-button
        type="primary"
          style="height: 35px; width: 100px"
          @click="showCreateProxy"
          >添加代理
        </el-button>
      </div>

      <el-table
        :data="httpProxyConfigs"
        stripe
        style="width: 100%; font-size: 18px"
        :row-style="{ height: '50px' }"
        :cell-style="{ padding: '0px' }"
      >
        <el-table-column prop="hostName" label="主机名" min-width="20%" />
        <el-table-column label="访问链接" min-width="40%" #default="scope">
          <a  v-if="!networkData.httpsEnable" target="_blank" :href="'http://' + scope.row.url">{{scope.row.url}}</a>
          <a  v-if="networkData.httpsEnable" target="_blank" :href="'https://' + scope.row.url">{{scope.row.url}}</a>
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

    <el-dialog
      v-model="setCaFilePathVisiable"
      title="设置证书文件路径"
      style="min-height: 300px; width: 600px"
    >
      <div class="center_div">
        <div>
          <div class="input_div2">
            <div>
              <el-tree-select
                class="w-50 m-2"
                style="width: 400px"
                size="large"
                v-model="caFilePath"
                check-strictly
                lazy
                :load="loadNode"
              >
                <template #default="{ data: { name } }">{{ name }}</template>
              </el-tree-select>
            </div>
          </div>
          <div class="center_div" style="margin-top: 50px">
            <el-button
              type="primary"
              style="height: 40px; width: 200px"
              @click="trySetCaPath"
              :loading="setCaFilePathLoading"
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
  setCaPath,
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

      setCaFilePathLoading: false,
      setCaFilePathVisiable: false,
      caFilePath: "",
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
          });
      } else {
        disableHttps()
          .then((response) => {
            this.changeHttpStateLoding = false;
            this.flush();
          })
          .catch((error) => {
            this.changeHttpStateLoding = false;
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
    showSetCaFilePath() {
      this.setCaFilePathVisiable = true;
    },
    trySetCaPath() {
      this.setCaFilePathLoading = true;
      setCaPath(this.caFilePath)
        .then((response) => {
          this.setCaFilePathLoading = false;
          this.setCaFilePathVisiable = false;
          this.flush();
        })
        .catch((error) => {
          this.setCaFilePathLoading = false;
        });
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