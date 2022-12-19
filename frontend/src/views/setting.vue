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
      <el-main class="card" style="min-height: 100px">Main</el-main>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "setting",
  data() {
    return {}
  },
  methods:{
    clicked(index){

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

