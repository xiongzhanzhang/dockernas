<template>
  <div style="white-space: pre-wrap;">{{ log }}</div>
</template>

<script>
import { getInstanceLog } from "../../api/instance";

export default {
  name: "instanceLog",
  props: ["name"],
  data() {
    return {
      log: "",
    };
  },
  methods: {
    initData(instance) {
      getInstanceLog(this.name).then((response) => {
        this.log = "";
        for(var s of response.data.data.split('\n')){
            // console.log(s);
            this.log+=s.substring(8)+"\n";
        };
      });
    },
  },
};
</script>

<style>
</style>