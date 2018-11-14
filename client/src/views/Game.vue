<template>
  <div>
    <div id="nav">
      <span v-on:click="reset">New game</span>|
      <span v-on:click="load">Refresh</span>
    </div>
    <div>
      <board/>
    </div>
    <div class="state">
      <state/>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import Board from "@/components/Board.vue";
import State from "@/components/State.vue";

export default {
  name: "game",

  components: {
    Board, State
  },
  mounted() {
    this.load();
  },
  methods: {
    load: function () {
      this.$store.dispatch("status");
    },
    reset: function () {
      this.$store.dispatch("newGame");
    }
  }
};
</script>
