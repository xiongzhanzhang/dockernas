<template>
  <div class="card">
    <div class="vertical_div">
      <el-image
        style="width: 120px; height: 120px;Border-radius:3px;margin-top: 36px;"
        src="https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg"
      />
    </div>
    <div class="vertical_div" style="">
      <div class="main_text">mysql</div>
      <div class="secondary_text">数据库</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "appCard",
  props: ["name", "summary", "url", "category"],
  data() {
    return {};
  },
};
</script>

<style scoped>
.card {
  width: 95%;
  height: 250px;
  background-color: white;
  margin: 6px;
}
</style>