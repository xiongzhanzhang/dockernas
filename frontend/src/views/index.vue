<template>
  <div style="width: 100%; height: 100%">
    <div class="common-layout">

      <el-container>
        <el-header class="no-padding">
          <el-menu
            active-text-color="#ffd04b"
            background-color="#545c64"
            text-color="#fff"
            class="el-menu-demo"
            mode="horizontal"
            default-active="apps"
            router
          >
            <el-menu-item index="apps">{{$t("index.apps")}}</el-menu-item>
            <el-menu-item index="store">{{$t("index.store")}}</el-menu-item>
            <el-menu-item index="setting">{{$t("index.setting")}}</el-menu-item>
          </el-menu>
        </el-header>

        <el-main class="no-padding">
          <RouterView />
        </el-main>

      </el-container>
    </div>
  </div>
</template>

<style>
.no-padding{
  padding: 0px;
}

</style>