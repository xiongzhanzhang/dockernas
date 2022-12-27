<template>
  <div class="common-layout" style="height: 100%">
    <el-container style="height: 100%">
      <el-aside class="card" width="200px">
        <div>
          <el-menu
            active-text-color="rgb(64,158,255)"
            text-color="#000"
            class="el-menu-demo"
            default-active="1"
            style="width: 100%"
          >
            <el-menu-item index="1" @click="clicked(0)"
              ><div class="menu-item">基本信息</div></el-menu-item
            >
            <el-menu-item index="2" @click="clicked(1)"
              ><div class="menu-item">存储</div></el-menu-item
            >
            <el-menu-item index="3" @click="clicked(2)"
              ><div class="menu-item">网络</div></el-menu-item
            >
          </el-menu>
        </div>
      </el-aside>
      <el-main style="padding:0px; margin:6px">
        <div v-show="divShow[0]">
          <host ref="host"></host>
        </div>
        <div v-show="divShow[1]">
          <storage ref="storage"></storage>
        </div>
        <div v-show="divShow[2]">
          <network ref="network"></network>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import host from "../components/setting/host.vue";
import storage from "../components/setting/storage.vue";
import network from "../components/setting/network.vue";

export default {
  name: "setting",
  components: {
    host,storage,network
  },
  data() {
    return {
      divShow: [true, false, false],
    }
  },
  methods:{
    clicked(index){
      this.divShow[0] = false;
      this.divShow[1] = false;
      this.divShow[2] = false;
      this.divShow[index] = true;

      this.$refs.storage.flush();
    }
  }
}
</script>

<style scoped>
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

