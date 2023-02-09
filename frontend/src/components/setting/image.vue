<template>
  <div>
    <div class="card_style" style="min-height: 600px">
      <el-table
        :data="image"
        stripe
        class="table_css"
        :row-style="{ height: '50px' }"
        :cell-style="{ padding: '0px' }"
        :row-key="row => { return row.id+row.name }"
      >
        <el-table-column prop="name" label="名字" sortable  min-width="30%" />
        <el-table-column prop="size"  label="大小" min-width="10%" sortable   #default="scope">
            <div>{{ (parseInt(scope.row.size, 10) / 1024 / 1024).toFixed(0) }} MB</div>
        </el-table-column>
        <el-table-column prop="state" label="拉取进度" sortable  min-width="15%" />
        <el-table-column label="操作" width="90px" #default="scope">
          <el-button size="small" type="danger" @click="delImage(scope.row)"
            >Delete</el-button
          >
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
import { getImages, delImages } from "../../api/image";

export default {
  name: "image",
  methods:{
    flush(){
        getImages().then((response) => {
            console.log(response);
            this.image=response.data.list;
        });
    },
    delImage(row){
        delImages(row).then((response) => {
            this.flush();
        })
    }
  },
  data(){
    return {
        image:[]
    }
  },
  mounted() {
    this.flush();
  },
};
</script>

<style>
@import "../../css/common.css";
@import "../../css/picture.css";
@import "../../css/menu.css";
@import "../../css/text.css";
@import "../../css/dialog.css";
</style>