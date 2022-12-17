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
            default-active="apps"
            style="width: 100%"
          >
            <el-menu-item index="1" @click="clicked(0)"><div class="menu-item">基本信息</div></el-menu-item>
            <el-menu-item index="2" @click="clicked(1)"><div class="menu-item">事件记录</div></el-menu-item>
            <el-menu-item index="3" @click="clicked(2)"><div class="menu-item">日志</div></el-menu-item>
            <el-menu-item index="4" @click="clicked(3)"><div class="menu-item">监控</div></el-menu-item>
          </el-menu>
          
        </div>
      </el-aside>
      <el-main class="card" style="min-height: 100px">
        <div v-show="divShow[0]">basicDiv</div>
        <div v-show="divShow[1]">eventDiv</div>
        <div v-show="divShow[2]">logDiv</div>
        <div v-show="divShow[3]">monitorDiv</div>
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

      divShow:[true,false,false,false]
    };
  },
  methods: {
    flushInstance() {
      getInstance(this.$route.params.name)
        .then((response) => {
          (this.appName = response.data.appName),
            (this.iconUrl = response.data.iconUrl),
            (this.state = response.data.state),
            (this.summary = response.data.summary),
            console.log(response);
        })
        .catch((error) => {
          console.log(error);
        });
    },
    clicked(index){
      this.divShow[0]=false;
      this.divShow[1]=false;
      this.divShow[2]=false;
      this.divShow[3]=false;
      this.divShow[index]=true;
    }
  },
  mounted() {
    this.flushInstance();
  },
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