<template>
  <div class="card">
    <div class="vertical_div">
      <el-image
        style="width: 120px; height: 120px;Border-radius:3px;margin-top: 36px;"
        :src="app.IconUrl"
      />
    </div>
    <div class="vertical_div" style="">
      <div class="main_text">{{app.Name}}</div>
      <div class="secondary_text">{{app.Category[0]}}</div>
    </div>
  </div>
</template>

<script>
export default {
  name: "appCard",
  props: ["app"],
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
  Border-radius:3px;
  box-shadow: 0 0 2px 2px #ccc;
}

.card:hover {
  box-shadow: 0 0 8px 8px #ccc;
}
</style>