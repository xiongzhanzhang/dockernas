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
    <createCard ref="createCard" :app="app"></createCard>
  </div>
</template>

<script>
import createCard from "../components/createApp.vue";
export default {
  name: "appCard",
  components: {
    createCard,
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
.card {
  width: 95%;
  height: 250px;
  background-color: white;
  margin: 6px;
  border-radius: 3px;
  box-shadow: 0 0 2px 2px #ccc;
}

.card:hover {
  box-shadow: 0 0 8px 8px #ccc;
}
</style>