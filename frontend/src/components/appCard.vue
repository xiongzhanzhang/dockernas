<template>
  <div class="card">
    <div class="vertical_div">
      <el-image
        style="width: 120px; height: 120px;Border-radius:3px;margin-top: 36px;"
        :src="url"
      />
    </div>
    <div class="vertical_div" style="">
      <div class="main_text">{{name}}</div>
      <div class="secondary_text">{{category}}</div>
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