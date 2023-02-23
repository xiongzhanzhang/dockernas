<template>
  <div>
    <div class="card" @click="clicked">
      <div class="vertical_div">
        <el-image
          style="
            margin-top: 10%;
          "
          class="image_icon"
          :src="getIconUrl(app.iconUrl)"
        />
      </div>
      <div class="vertical_div" style="">
        <div class="main_text">{{ app.name }}</div>
        <div class="secondary_text">{{ app.category[0] }}</div>
      </div>
    </div>
    <createInstance ref="createCard"></createInstance>
  </div>
</template>

<script>
import createInstance from "./createInstance.vue";
import {getIconUrl} from "@/utils/url"
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
    getIconUrl,
    clicked() {
      this.$refs.createCard.showDialog();
    },
  },
  mounted(){
    this.$refs.createCard.setApp(this.app);
  }
};
</script>

<style scoped>
@import "../css/common.css";
@import "../css/picture.css";
@import "../css/text.css";
</style>