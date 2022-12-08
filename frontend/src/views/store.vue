<template>
  <div class="main_page">
    <div class="two_item_div">
      <div class="two_item_div">
        <div
          style="
            margin-left: 10px;
            margin-right: 10px;
            font-weight: 500;
            font-size: large;
          "
          >{{ $t("store.appType") }}</div
        >
        <el-select
          v-model="selectTypes"
          multiple
          collapse-tags
          placeholder="Select"
          size="large"
          @change="filterChange"
        >
          <el-option
            v-for="item in appTypes"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          />
        </el-select>
      </div>
      <div>
        <el-input
          v-model="searchStr"
          class="w-50 m-2"
          style="width:300px"
          placeholder="search"
          size="large"
        >
          <template #prefix>
            <el-icon class="el-input__icon"><search /></el-icon>
          </template>
        </el-input>
      </div>
    </div>

    <div style="padding:6px">
        <el-row>
            <el-col :xs="8" :sm="6" :md="4" :lg="4"><appCard></appCard></el-col>
            <el-col :xs="8" :sm="6" :md="4" :lg="4"><appCard></appCard></el-col>
            <el-col :xs="8" :sm="6" :md="4" :lg="4"><appCard></appCard></el-col>
            <el-col :xs="8" :sm="6" :md="4" :lg="4"><appCard></appCard></el-col>
            <el-col :xs="8" :sm="6" :md="4" :lg="4"><appCard></appCard></el-col>
            <el-col :xs="8" :sm="6" :md="4" :lg="4"><appCard></appCard></el-col>
        </el-row>
    </div>
  </div>
</template>


<script>
import appCard from "../components/appCard.vue";

export default {
    components: {
        appCard,
  },
  name: "store",
  data() {
    return {
      selectTypes: [],
      searchStr: "",
      appTypes: [
        {
          value: "Option1",
          label: "Option1",
        },
        {
          value: "Option2",
          label: "Option2",
        }
      ],
    };
  },
  methods: {
    filterChange() {},
  },
};
</script>


<style scoped>
@import "../css/common.css";
</style>