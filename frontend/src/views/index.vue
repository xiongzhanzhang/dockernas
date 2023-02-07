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
            <el-menu-item style="font-weight: bold; font-size: 20px"
              >DockerNAS</el-menu-item
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

            <!-- <div style="flex-grow: 1"></div> -->

            <el-sub-menu>
              <template #title>帮助</template>
              <el-menu-item @click="gotoLink('https://doc.dockernas.com')" index="">文档</el-menu-item>
              <el-menu-item @click="gotoLink('https://doc.dockernas.com')" index="">社区</el-menu-item>
              <el-menu-item @click="gotoLink('https://github.com/xiongzhanzhang/dockernas')" index="">github</el-menu-item>
              <el-menu-item @click="gotoLink('https://gitee.com/xiongzhanzhang/dockernas')" index="">gitee</el-menu-item>
              <el-menu-item @click="logout" index="">退出</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-header>

        <el-main class="main_router">
          <RouterView v-slot="{ Component }">
            <component :is="Component" />
          </RouterView>
        </el-main>
      </el-container>
    </div>
  </div>
</template>

<script>
import { splitRouterPathByIndex } from "../utils/url";
import storage from "@/utils/storage";

export default {
  name: "index",
  data() {
    return {};
  },
  methods: {
    splitRouterPathByIndex,
    gotoLink(url) {
      window.open(url);
    },
    logout() {
      storage.rm("token");
      storage.rm("user");
      this.$router.push("/login");
    },
  },
  mounted() {
    console.log(this.$router);
  },
};
</script>

<style>
@import "../css/common.css";
</style>