<template>
  <div>
    <div class="card" @click="clicked">
      <div class="vertical_div">
        <el-image
          style="
            width: 120px;
            height: 120px;
            border-radius: 3px;
            margin-top: 36px;
          "
          :src="app.iconUrl"
        />
      </div>
      <div class="vertical_div" style="">
        <div class="main_text">{{ app.name }}</div>
        <div class="secondary_text">{{ app.category[0] }}</div>
      </div>
    </div>
    <createInstance ref="createCard" :app="app"></createInstance>
  </div>
</template>

<script>
import createInstance from "./createInstance.vue";
export default {
  name: "appCard",
  components: {
    createInstance,
  },
  props: ["app"],
  data() {
    return {};
  },
  methods: {
    clicked() {
      this.$refs.createCard.showDialog();
    },
  },
};
</script>

<style scoped>
@import "../css/common.css";
</style>