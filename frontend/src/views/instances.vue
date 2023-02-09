<template>
  <div class="main_page">
    <div style="padding: 6px">
      <el-row>
        <el-col
          :xs="12"
          :sm="12"
          :md="8"
          :lg="6"
          v-for="instance in instances"
          :key="instance.name"
        >
          <instanceCard :instance="instance"></instanceCard>
        </el-col>
      </el-row>
    </div>

    <div v-show="instances.length==0" class="tip_div"><div>没有实例</div></div>

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

<style scoped>
.tip_div{
  color: rgb(94, 94, 94);
  font-size: 18px;
  display: flex;
  justify-content: center;
  margin-top: 30px;
}
</style>

  