<template>
  <div class="center_div">
    <div class="login_card">
      <div class="login_txt center_div">设置存储路径</div>
      <div class="vertical_div">
        <el-tree-select
          class="login_input_div"
          size="large"
          v-model="path"
          check-strictly
          lazy
          :load="loadNode"
        >
          <template #default="{ data: { name } }">{{ name }}</template>
        </el-tree-select>

        <el-button
          type="primary"
          class="login_input_div"
          :loading="btnLoading"
          @click="trySetBasePath"
          >确定</el-button
        >
      </div>
    </div>
  </div>
</template>

<script>
import { getSystemDirs, setBasePath } from "../api/filesystem";

export default {
  name: "basePath",
  data() {
    return {
      path: "",
      btnLoading: false,
    };
  },
  methods: {
    loadNode(node, resolve) {
      console.log(node.data);
      if (node.isLeaf) return resolve([]);
      var curPath = node.data.value;
      if (curPath == null) {
        curPath = "";
      }

      getSystemDirs(curPath).then((response) => {
        console.log(response.data);
        resolve(response.data.list);
      });
    },
    trySetBasePath() {
        this.btnLoading=true;
        setBasePath(this.path).then((response) => {
            this.btnLoading=false;
            this.$router.push("/index/store");
        }).catch((error) => {
            this.btnLoading=false;
        })
    },
  },
};
</script>


<style scoped>
@import "../css/common.css";
@import "../css/login.css";

.login_input_div{
  margin-top: 50px;
}
</style>