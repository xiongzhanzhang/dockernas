<template>
  <div>
    <div class="card" @click="clicked">
      <div class="vertical_div">
        <el-image
          style="width: 120px; height: 120px; border-radius: 3px"
          :src="instance.iconUrl"
        />
      </div>
      <div class="text_div" style="padding-left: 6px">
        <div class="main_text">{{ instance.name }}</div>
        <div class="secondary_text">{{ instance.summary }}</div>
      </div>
      <div style="flex-grow: 1"></div>
      <div style="height: 100%">
        <div class="color_dot" v-if="instance.state == 0 || instance.state == 4" style="background-color:gray"></div>
        <div class="color_dot" v-if="instance.state == 1" style="background-color:yellow"></div>
        <div class="color_dot" v-if="instance.state == 2" style="background-color:red"></div>
        <div class="color_dot" v-if="instance.state == 3" style="background-color:green"></div>
      </div>
    </div>
    <!-- <createInstance ref="createCard" :app="instance"></createInstance> -->
  </div>
</template>
  
  <script>
import createInstance from "./createInstance.vue";
export default {
  name: "instanceCard",
  components: {
    createInstance,
  },
  props: ["instance"],
  data() {
    return {};
  },
  methods: {
    clicked() {
      this.$router.push("/index/instanceinfo/"+this.instance.name);
    },
  },
};
</script>
  
  <style scoped>
@import "../css/common.css";

.card {
  width: 95%;
  height: 150px;

  margin: 6px;
  padding-left: 6px;

  background-color: white;
  border-radius: 3px;
  box-shadow: 0 0 2px 2px #ccc;

  display: flex;
  align-items: center;
}

.card:hover {
  box-shadow: 0 0 8px 8px #ccc;
}

.color_dot {
  width: 10px;
  height: 10px;
  margin: 8px;
  border-radius: 10px;
  /* background-color: green; */
}

.text_div {
    display: flex;
    justify-content: center;
    flex-direction: column;
}

</style>