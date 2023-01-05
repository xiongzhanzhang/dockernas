<template>
  <div ref="terminal" style="height: 100%"></div>
</template>

<script>
import "xterm/css/xterm.css";
import { Terminal } from "xterm";
import { FitAddon } from "xterm-addon-fit";
import { AttachAddon } from "xterm-addon-attach";

import { getWebTerminalWebsocket } from "../../api/instance";

export default {
  name: "instanceTerminal",
  props:["instanceData"],
  data() {
    return {
      instance: {},
    };
  },
  methods: {
    connect() {
      if (this.term != null) {
        return;
      }

      var term = new Terminal({
        fontSize: 14,
        cursorBlink: true,
      });
      var fitAddon = new FitAddon();
      term.loadAddon(fitAddon);
      term.open(this.$refs.terminal);
      fitAddon.fit();

      var ws = getWebTerminalWebsocket(this.instance.containerID, term.cols);
      var attachAddon = new AttachAddon(ws);
      term.loadAddon(attachAddon);
      term.focus();
      this.term = term;

      window.onresize = function () {
        fitAddon.fit();
      };
    },
    initData(instance) {
        if (instance==null){
            return
        }
      this.instance = instance;
      console.log(instance);
      this.connect();
    },
  },
  mounted() {
    this.instance=this.instanceData;
    this.initData(this.instance);
  },
};
</script>

<style>
</style>