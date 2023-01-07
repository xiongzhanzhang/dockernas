<template>
  <div>
    <div class="card" style="min-height: 600px">
      <div class="center_div" style="padding-top: 10px; padding-bottom: 20px">
        <div style="flex-grow: 1"></div>
        <el-button
          type="primary"
          style="height: 40px; width: 150px; margin-right: 30px"
          @click="createSubcribe"
          >添加订阅</el-button
        >
        <el-button
          type="success"
          style="height: 40px; width: 150px; margin-right: 30px"
          @click="tryFlushSubcribe"
          >刷新订阅</el-button
        >
      </div>

      <el-table
        :data="subscribeData"
        stripe
        style="width: 100%; font-size: 18px"
        :row-style="{ height: '50px' }"
        :cell-style="{ padding: '0px' }"
      >
        <el-table-column prop="name" label="名字" min-width="10%" />
        <el-table-column prop="url" label="git链接" min-width="40%" />
        <el-table-column prop="state" label="状态" min-width="10%" />
        <el-table-column prop="updateTime" label="更新时间" min-width="15%" />
        <el-table-column label="操作" min-width="10%" #default="scope">
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
      style="min-height: 300px"
    >
      <div class="center_div">
        <div>
          <div class="input_div2">
            <div class="first_input">名字</div>
            <div>
              <el-input
                class="w-50 m-2"
                style="width: 400px"
                size="large"
                v-model="curSubscribeName"
              >
              </el-input>
            </div>
          </div>
          <div class="input_div2">
            <div class="first_input">git链接</div>
            <div>
              <el-input
                class="w-50 m-2"
                style="width: 400px"
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

<style scoped>
.card {
  /* padding: 6px; */
  /* margin: 6px; */
  border-radius: 3px;
  background-color: white;
}

.input_div {
  display: flex;
  color: black !important;
  font-size: 18px;
  align-items: center;
  height: 80px;
  border-bottom: 1px solid rgb(222, 222, 222);
}
.first_input {
  width: 33%;
  padding-left: 30px;
  /* text-align: right; */
  margin-right: 80px;
}

.input_div2 {
  display: flex;
  color: black !important;
  font-size: 18px;
  align-items: center;
  height: 80px;
}
</style>