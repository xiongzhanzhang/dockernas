<template>
  <div class="main_page">
    <div style="padding: 6px">
      <el-row>
        <el-col
          :xs="8"
          :sm="6"
          :md="4"
          :lg="4"
          v-for="instance in instances"
          :key="instance.name"
        >
          <instanceCard :instance="instance"></instanceCard>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import instanceCard from "../components/instanceCard.vue";
import { getAllInstance } from "../api/instance";

export default {
  name: "instances",
  components: {
    instanceCard,
  },
  data() {
    return {
        instances:[]
    };
  },
  methods: {
    flushInstances() {
      getAllInstance()
        .then((response) => {
          console.log(response);
          this.instances = response["data"]["list"];
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
  mounted() {
    this.flushInstances();
  },
};
</script>

<style>
</style>

  