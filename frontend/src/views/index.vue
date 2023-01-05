<template>
  <div class="full_size">
    <div class="common-layout full_size">
      <el-container class="full_size">
        <el-header class="no-padding" style="height: 50px">
          <el-menu
            active-text-color="#ffd04b"
            background-color="#545c64"
            text-color="#fff"
            class="el-menu-demo"
            mode="horizontal"
            :default-active="
              splitRouterPathByIndex($router.currentRoute.value.path, 3)
            "
            router
            style="height: 100%"
          >
            <el-menu-item index="/index/instances">{{
              $t("index.apps")
            }}</el-menu-item>
            <el-menu-item index="/index/store">{{
              $t("index.store")
            }}</el-menu-item>
            <el-menu-item index="/index/setting">{{
              $t("index.setting")
            }}</el-menu-item>
          </el-menu>
        </el-header>

        <el-main class="main_router">
          <RouterView v-slot="{ Component }">
            <keep-alive>
              <component :is="Component" />
            </keep-alive>
          </RouterView>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script>
import { splitRouterPathByIndex } from "../utils/url";

export default {
  name: "index",
  data() {
    return {};
  },
  methods: {
    splitRouterPathByIndex,
  },
  mounted() {
    console.log(this.$router);
  },
};
</script>

<style>
@import "../css/common.css";
</style>