<template>
  <div class="card_style">
    <div style="padding: 16px; display: flex">
      <el-date-picker
        v-model="time"
        type="datetimerange"
        size="large"
        class="big_input"
        range-separator="To"
        start-placeholder="Start date"
        end-placeholder="End date"
        @change="setTime"
      />
      <div style="flex-grow: 1000"></div>
      <div class="click_able" @click="initData">
        <el-icon size="20"><Refresh /></el-icon>
      </div>
    </div>
    <el-row>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div ref="cpu" class="chart"></div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div ref="mem" class="chart"></div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div ref="network_in" class="chart"></div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div ref="network_out" class="chart"></div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div ref="disk_in" class="chart"></div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div ref="disk_out" class="chart"></div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import * as echarts from "echarts";
import { getInstanceStats, getInstanceStatsByName } from "../../api/instance";

export default {
  name: "instanceMonitor",
  props: ["name"],
  data() {
    return {
      isTimeSeted: false,
      time: [new Date(new Date().getTime() - 3 * 60 * 60 * 1000), new Date()],
    };
  },
  methods: {
    init() {
      var cpuChart = echarts.init(this.$refs.cpu);
      var memChart = echarts.init(this.$refs.mem);
      var networkInChart = echarts.init(this.$refs.network_in);
      var networkOutChart = echarts.init(this.$refs.network_out);
      var diskInChart = echarts.init(this.$refs.disk_in);
      var diskOutChart = echarts.init(this.$refs.disk_out);

      window.onresize = function () {
        cpuChart.resize();
        memChart.resize();
        networkInChart.resize();
        networkOutChart.resize();
        diskInChart.resize();
        diskOutChart.resize();
      };

      this.cpuChart = cpuChart;
      this.memChart = memChart;
      this.networkInChart = networkInChart;
      this.networkOutChart = networkOutChart;
      this.diskInChart = diskInChart;
      this.diskOutChart = diskOutChart;
    },
    setTime() {
      this.isTimeSeted = true;
      this.initData();
    },
    computeOption(dataList, title, field, factor, precision) {
      var seriesMap = {};
      for (var d of dataList) {
        if (seriesMap[d.name] == null) {
          seriesMap[d.name] = {
            name: d.name,
            type: "line",
            // stack: "Total",
            // areaStyle: {},
            emphasis: {
              focus: "series",
            },
            data: [],
          };
        }
        seriesMap[d.name]["data"].push([
          d.createTime,
          (d[field] / factor).toFixed(precision),
        ]);
      }

      var series = [];
      var legend = [];
      for (let k of Object.keys(seriesMap)) {
        legend.push(k);
        series.push(seriesMap[k]);
      }

      var option = {
        title: {
          left: "center",
          text: title,
        },
        tooltip: {
          trigger: "axis",
          axisPointer: {
            type: "cross",
            label: {
              backgroundColor: "#6a7985",
            },
          },
        },
        legend: {
          top: "7%",
          data: legend,
        },
        grid: {
          left: "3%",
          right: "4%",
          bottom: "3%",
          containLabel: true,
        },
        xAxis: [
          {
            type: "time",
            boundaryGap: false,
          },
        ],
        yAxis: [
          {
            type: "value",
          },
        ],
        series: series,
      };

      return option;
    },
    setData(data) {
      this.cpuChart.clear();
      this.memChart.clear();
      this.networkInChart.clear();
      this.networkOutChart.clear();
      this.diskInChart.clear();
      this.diskOutChart.clear();

      this.cpuChart.setOption(
        this.computeOption(data, "cpu使用百分百", "CPUPercentage", 1, 2),
        true
      );
      this.memChart.setOption(
        this.computeOption(data, "内存使用(MB)", "memory", 1024 * 1024, 0),
        true
      );
      this.networkInChart.setOption(
        this.computeOption(data, "网络下行速度(KB/s)", "networkRx", 1024, 0),
        true
      );
      this.networkOutChart.setOption(
        this.computeOption(data, "网络上行速度(KB/s)", "networkTx", 1024, 0),
        true
      );
      this.diskInChart.setOption(
        this.computeOption(data, "磁盘写速度(KB/s)", "blockWrite", 1024, 0),
        true
      );
      this.diskOutChart.setOption(
        this.computeOption(data, "磁盘读速度(KB/s)", "blockRead", 1024, 0),
        true
      );

      this.cpuChart.resize();
      this.memChart.resize();
      this.networkInChart.resize();
      this.networkOutChart.resize();
      this.diskInChart.resize();
      this.diskOutChart.resize();
    },
    initData() {
      if (this.isTimeSeted == false) {
        this.time[1] = new Date();
      }
      if (this.name == null || this.name == "") {
        getInstanceStats(this.time[0].getTime(), this.time[1].getTime()).then(
          (response) => {
            this.setData(response.data.list);
          }
        );
      } else {
        getInstanceStatsByName(
          this.name,
          this.time[0].getTime(),
          this.time[1].getTime()
        ).then((response) => {
          this.setData(response.data.list);
        });
      }
    },
  },
  mounted() {
    this.init();
    this.initData();
  },
};
</script>

<style>
@import "../../css/common.css";
@import "../../css/picture.css";
@import "../../css/text.css";

.chart {
  margin-top: 18px;
  height: 400px;
}
</style>