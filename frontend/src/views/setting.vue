<template>
  <div class="common-layout" style="height: 100%">
    <el-container style="height: 100%">
      <el-aside class="card" width="200px">
        <div>
          <el-menu
            active-text-color="rgb(64,158,255)"
            text-color="#000"
            class="el-menu-demo"
            :default-active="$router.currentRoute.value.path"
            style="width: 100%"
            router
          >
            <el-menu-item index="/index/setting/host"
              ><div class="menu-item">基本信息</div></el-menu-item
            >
            <el-menu-item index="/index/setting/storage"
              ><div class="menu-item">存储</div></el-menu-item
            >
            <el-menu-item index="/index/setting/network"
              ><div class="menu-item">网络</div></el-menu-item
            >
            <el-menu-item index="/index/setting/subscribe"
              ><div class="menu-item">订阅</div></el-menu-item
            >
            <el-menu-item index="/index/setting/image"
              ><div class="menu-item">镜像</div></el-menu-item
            >
          </el-menu>
        </div>
      </el-aside>
      <el-main style="padding: 0px; margin: 6px">
        <RouterView v-slot="{ Component }">
          <component ref="view" :is="Component" />
        </RouterView>
      </el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "setting",
  data() {
    return {
    };
  },
  methods: {
  },
};
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

