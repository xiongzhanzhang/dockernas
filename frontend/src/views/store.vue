<template>
  <div class="main_page">
    <div class="two_item_div">
      <div class="two_item_div">
        <div style="margin-right: 10px; font-weight: 500; font-size: large">
          {{ $t("store.appType") }}
        </div>
        <el-select
          v-model="selectTypes"
          multiple
          collapse-tags
          placeholder="Select"
          size="large"
          @change="filterApps"
        >
          <el-option
            v-for="item in appTypes"
            :key="item"
            :label="item"
            :value="item"
          />
        </el-select>
      </div>
      <div>
        <el-input
          v-model="searchStr"
          class="w-50 m-2"
          style="width: 300px"
          placeholder="search"
          size="large"
          @input="filterApps"
        >
          <template #prefix>
            <el-icon class="el-input__icon"><search /></el-icon>
          </template>
        </el-input>
      </div>
    </div>

    <div style="padding: 6px">
      <el-row>
        <el-col
          :xs="8"
          :sm="6"
          :md="4"
          :lg="4"
          v-for="app in filtedApps"
          :key="app.name"
        >
          <appCard :app="app"></appCard>
        </el-col>
      </el-row>
    </div>
  </div>
</template>


<script>
import appCard from "../components/appCard.vue";
import { getApps } from "../api/store";

export default {
  components: {
    appCard,
  },
  name: "store",
  data() {
    return {
      apps: [],
      filtedApps:[],
      selectTypes: [],
      searchStr: "",
      appTypes: [],
    };
  },
  methods: {
    flushApps() {
      getApps()
        .then((response) => {
          console.log(response);
          this.apps = response["data"]["list"];

          this.appTypes = [];
          let categorys = [];
          for (let app of this.apps) {
            categorys = categorys.concat(app.category);
          }
          for (let category of new Set(categorys)) {
            this.appTypes.push(category);
          }

          this.filterApps();
        })
        .catch((error) => {
          console.log(error);
        });
    },
    filterApps(){
      this.filtedApps=[];
      console.log(this.searchStr);
      console.log(this.selectTypes);

      for(let app of this.apps){
        console.log(app);
        if(app.name.toLowerCase().indexOf(this.searchStr.toLowerCase())>=0 || this.searchStr.length==0){
          if(this.selectTypes.length>0){
            for(let category of app.category){
              if(this.selectTypes.indexOf(category)>=0){
                this.filtedApps.push(app);
                break;
              }
            }
          }else{
            this.filtedApps.push(app);
          }
        }
      }
    }
  },
  mounted() {
    this.flushApps();
  },
};
</script>


<style scoped>
@import "../css/common.css";
</style>