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
  props: ["name"],
  data() {
    return {
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

      var ws = getWebTerminalWebsocket(this.name, term.cols);
      var attachAddon = new AttachAddon(ws);
      term.loadAddon(attachAddon);
      term.focus();
      this.term = term;

      window.onresize = function () {
        fitAddon.fit();
      };
    }
  },
  mounted() {
    this.connect();
  },
};
</script>

<style>
</style>