<template>
  <div ref="terminal" style="height: 100%"></div>
</template>

<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";

import { getInstanceLog } from "../../api/instance";

export default {
  name: "instanceLog",
  props: ["name"],
  data() {
    return {
      log: "",
      term:null,
    };
  },
  methods: {
    initXterm(){
      var term = new Terminal({
        fontSize: 14,
        cursorBlink: false,
        disableStdin: true,
        theme: {
          foreground: "#ECECEC",
          background: "#000000", 
          cursor: "#000000", 
        }
      });
      var fitAddon = new FitAddon();
      term.loadAddon(fitAddon);
      term.open(this.$refs.terminal);
      fitAddon.fit();
      term.focus();
      this.term = term;
    },
    initData() {
      getInstanceLog(this.name).then((response) => {
        this.log = "";
        for(var s of response.data.data.split('\n')){
            // console.log(s);
            this.log+=s.substring(8)+"\r\n";
        };
        this.term.write(this.log);
      });
    },
  },
  mounted(){
    this.initXterm();
    this.initData();
  }
};
</script>

<style>
</style>