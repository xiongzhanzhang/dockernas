<template>
  <div>
    <div class="card_style" style="min-height: 600px">
      <div class="center_div" style="padding-top: 5px; padding-bottom: 10px">
        <div style="flex-grow: 1"></div>
        <el-button
          type="primary"
          style="height: 35px; width: 100px; margin-right: 10px"
          @click="createSubcribe"
          >添加订阅</el-button
        >
        <el-button
          type="success"
          style="height: 35px; width: 100px; margin-right: 10px"
          @click="tryFlushSubcribe"
          >刷新订阅</el-button
        >
      </div>

      <el-table
        :data="subscribeData"
        stripe
        class="table_css"
        :row-style="{ height: '50px' }"
        :cell-style="{ padding: '0px' }"
      >
        <el-table-column prop="name" label="名字" min-width="10%" />
        <el-table-column prop="url" label="git链接" min-width="40%" />
        <el-table-column prop="state" label="状态" min-width="10%" />
        <el-table-column prop="updateTime" label="更新时间" min-width="15%" />
        <el-table-column label="操作" width="90px" #default="scope">
          <el-button
            size="small"
            type="danger"
            @click="delSubscribe(scope.row)"
            >Delete</el-button
          >
        </el-table-column>
      </el-table>
    </div>

    <el-dialog
      v-model="createSubcribeVisible"
      title="创建订阅"
      class="big_dialog"
    >
      <div class="center_div">
        <div>
          <div class="input_div">
            <div class="first_input">名字</div>
            <div>
              <el-input
                class="big_input"
                size="large"
                v-model="curSubscribeName"
              >
              </el-input>
            </div>
          </div>
          <div class="input_div">
            <div class="first_input">git链接</div>
            <div>
              <el-input
                class="big_input"
                size="large"
                v-model="curUrl"
              >
              </el-input>
            </div>
          </div>
          <div class="center_div" style="margin-top: 50px">
            <el-button
              type="primary"
              style="height: 40px; width: 200px"
              @click="tryCreateSubcribe"
              >{{ $t("common.yes") }}</el-button
            >
          </div>
        </div>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getSubscribes, createSubscribe, deleteSubscribe, flushSubscribe} from "../../api/subscribe";

export default {
  name: "subscribe",
  data() {
    return {
      subscribeData: [],
      createSubcribeVisible: false,
      curSubscribeName: "",
      curUrl: "",
    };
  },
  methods: {
    getStateStr(state) {
      if (state == 0) return "初始化";
      if (state == 1) return "初始化失败";
      if (state == 2) return "更新成功";
      if (state == 3) return "更新中";
      if (state == 4) return "更新失败";

      return "未知";
    },
    createSubcribe() {
      this.curSubscribeName="";
      this.curUrl="";
      this.createSubcribeVisible = true;
    },
    tryCreateSubcribe() {
      createSubscribe(this.curSubscribeName, this.curUrl).then((response) => {
        this.createSubcribeVisible = false;
        this.init();
      });
    },
    tryFlushSubcribe() {
      flushSubscribe().then((response) => {
        this.init();
      });
    },
    delSubscribe(row){
      deleteSubscribe(row.name).then((response) => {
        this.init();
      })
    },
    init() {
      getSubscribes().then((response) => {
        this.subscribeData=[];
        for (var d of response.data.list) {
          this.subscribeData.push({
            updateTime: new Date(d.updateTime).toLocaleString(),
            state: this.getStateStr(d.state),
            url: d.url,
            name: d.name,
          });
        }
      });
    },
  },
  mounted() {
    this.init();
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