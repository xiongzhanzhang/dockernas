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
          :src="instance.iconUrl"
        />
      </div>
      <div class="vertical_div" style="">
        <div class="main_text">{{ instance.name }}</div>
        <div class="secondary_text">{{ instance.appName }}</div>
      </div>
    </div>
    <!-- <createCard ref="createCard" :app="instance"></createCard> -->
  </div>
</template>
  
  <script>
import createCard from "../components/createApp.vue";
export default {
  name: "instanceCard",
  components: {
    createCard,
  },
  props: ["instance"],
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