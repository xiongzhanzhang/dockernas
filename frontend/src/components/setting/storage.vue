<template>
  <div>
    <div class="card_style">
      <div class="input_div bottom_border">
        <div class="table_first_input">存储路径</div>
        <div>{{ storageData.baseDir }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">设备</div>
        <div>{{ storageData.device }}</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">容量</div>
        <div>{{ getSizeInGB(storageData.capacity, 0) }} GB</div>
      </div>
      <div class="input_div bottom_border">
        <div class="table_first_input">格式</div>
        <div>{{ storageData.fstype }}</div>
      </div>
    </div>
    <div class="card_style" style="margin-top: 16px">
      <el-row>
        <el-col :xs="24" :sm="12" :md="12" :lg="12">
          <div ref="all" class="chart"></div>
        </el-col>
        <el-col :xs="24" :sm="12" :md="12" :lg="12">
          <div ref="local" class="chart"></div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import * as echarts from "echarts";
import { getStorageInfo } from "../../api/host";

export default {
  name: "storage",
  data() {
    return {
      storageData: {},
    };
  },
  methods: {
    getSizeInGB(size, factor) {
      return (size / 1024 / 1024 / 1024).toFixed(factor);
    },
    init() {
      if (this.allChart != null) {
        return;
      }

      var allChart = echarts.init(this.$refs.all);
      var localChart = echarts.init(this.$refs.local);

      window.onresize = function () {
        allChart.resize();
        localChart.resize();
      };

      this.allChart = allChart;
      this.localChart = localChart;
    },
    setData() {
      var instanceDataSize = [];
      for (let k of Object.keys(this.storageData.instanceSizeMap)) {
        instanceDataSize.push({
          value: this.getSizeInGB(this.storageData.instanceSizeMap[k], 2),
          name: k,
        });
      }

      var localOption = {
        title: {
          text: "实例存储使用(GB)",
          left: "center",
        },
        tooltip: {
          trigger: "item",
        },
        series: [
          {
            type: "pie",
            radius: "70%",
            data: instanceDataSize,
          },
        ],
      };

      var allOption = {
        title: {
          text: "设备存储使用(GB)",
          left: "center",
        },
        tooltip: {
          trigger: "item",
        },
        series: [
          {
            type: "pie",
            radius: "70%",
            data: [
              {
                value: this.getSizeInGB(this.storageData.freeSize, 0),
                name: "空闲",
              },
              {
                value: this.getSizeInGB(this.storageData.otherSize, 0),
                name: "其他",
              },
              {
                value: this.getSizeInGB(this.storageData.dfsSize, 0),
                name: "共享存储",
              },
              {
                value: this.getSizeInGB(this.storageData.localSize, 0),
                name: "实例存储",
              },
            ],
          },
        ],
      };

      this.allChart.clear();
      this.localChart.clear();
      this.allChart.setOption(allOption, true);
      this.localChart.setOption(localOption, true);
      this.allChart.resize();
      this.localChart.resize();
    },
    flush() {
      this.init();
      getStorageInfo().then((response) => {
        this.storageData = response.data;
        this.setData();
        console.log(response.data);
      });
    },
  },
  mounted() {
    this.flush();
  },
};
</script>

<style scoped>
@import "../../css/common.css";
@import "../../css/picture.css";
@import "../../css/menu.css";
@import "../../css/text.css";
.chart {
  margin-top: 18px;
  height: 400px;
}
</style>