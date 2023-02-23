<template>
  <div>
    <div class="card" @click="clicked" style="position:relative">
      <div class="vertical_div">
        <el-image
          :class='[url==null ?"":"click_able","image_icon"]'
          @click="tryOpen"
          :src="getIconUrl(instance.iconUrl)"
        />
      </div>
      <div class="text_div" style="padding-left: 6px; flex-grow: 1;">
        <div class="main_text" >{{ instance.name }}</div>
        <div class="secondary_text">{{ instance.summary }}</div>
      </div>
      <div style="position:absolute;right:0;top:0">
        <div class="color_dot" v-if="instance.state == 4" style="background-color:gray"></div>
        <div class="color_dot" v-if="instance.state == 0 || instance.state == 1" style="background-color:yellow"></div>
        <div class="color_dot" v-if="instance.state == 2" style="background-color:red"></div>
        <div class="color_dot" v-if="instance.state == 3" style="background-color:green"></div>
      </div>
    </div>
    <!-- <createInstance ref="createCard" :app="instance"></createInstance> -->
  </div>
</template>
  
  <script>
import createInstance from "./createInstance.vue";
import {getFirstHttpPortUrl, getIconUrl} from "../utils/url"

export default {
  name: "instanceCard",
  components: {
    createInstance,
  },
  props: ["instance"],
  data() {
    return {
      url:null
    };
  },
  methods: {
    getIconUrl,
    clicked() {
      this.$router.push("/index/instances/"+this.instance.name+"/basicInfo");
    },
    tryOpen(event){
      if(this.url!=null){
        event.stopPropagation();
        window.open(this.url);
      }
    }
  },
  mounted(){
    if(this.instance.url!=""){
      this.url=this.instance.url;
    }else{
      this.url=getFirstHttpPortUrl(JSON.parse(this.instance.instanceParamStr));
    }
  }
};
</script>
  
  <style scoped>
@import "../css/common.css";
@import "../css/picture.css";
@import "../css/text.css";

.card {
  margin: 6px;
  padding-left: 6px;

  background-color: white;
  border-radius: 3px;
  box-shadow: 0 0 2px 2px #ccc;

  display: flex;
  align-items: center;
}

@media (max-width: 600px) {
  .card{
    height: 100px;
  }
}
@media (min-width: 600px) {
  .card{
    height: 150px;
  }
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