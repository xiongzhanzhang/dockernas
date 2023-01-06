<template>
  <div class="common-layout" style="height: 100%">
    <el-container style="height: 100%">
      <el-aside class="card" width="200px">
        <div class="vertical_div">
          <el-image
            style="width: 160px; height: 160px; border-radius: 3px"
            :src="iconUrl"
          ></el-image>

          <el-menu
            active-text-color="rgb(64,158,255)"
            text-color="#000"
            class="el-menu-demo"
            :default-active="$router.currentRoute.value.path"
            style="width: 100%"
            router
          >
            <el-menu-item :index="`/index/instances/${name}/basicInfo`"
              ><div class="menu-item">基本信息</div></el-menu-item
            >
            <el-menu-item :index="`/index/instances/${name}/event`"
              ><div class="menu-item">事件记录</div></el-menu-item
            >
            <el-menu-item :index="`/index/instances/${name}/log`"
              ><div class="menu-item">日志</div></el-menu-item
            >
            <el-menu-item :index="`/index/instances/${name}/monitor`" 
              ><div class="menu-item">监控</div></el-menu-item
            >
            <el-menu-item :index="`/index/instances/${name}/terminal`"
              ><div class="menu-item">终端</div></el-menu-item
            >
          </el-menu>
        </div>
      </el-aside>
      <el-main class="card" style="min-height: 100px">
        <RouterView  v-slot="{ Component }">
          <keep-alive>
            <component  ref="view" :is="Component"/>
          </keep-alive>
        </RouterView>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { getInstance } from "../api/instance";

export default {
  name: "instanceInfo",
  data() {
    return {
      name: this.$route.params.name,
      appName: "",
      iconUrl: "",
      state: null,
      summary: "",
      instanceData: {},
    };
  },
  methods: {
    flushInstance() {
      getInstance(this.$route.params.name)
        .then((response) => {
          this.instanceData = response.data;
          this.appName = response.data.appName;
          this.iconUrl = response.data.iconUrl;
          this.state = response.data.state;
          this.summary = response.data.summary;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  created(){
      this.flushInstance();
  },
  mounted() {},
};
</script>

<style scoped>
@import "../css/common.css";

.card {
  /* padding: 6px; */
  margin: 6px;
  border-radius: 3px;
  background-color: white;
}

.menu-item {
  width: 100%;
  text-align: center;
  font-size: 20px;
}
</style>