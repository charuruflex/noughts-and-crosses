<template>
  <div class="board">
    <div
      class="cell"
      v-for="(cell, i) in $store.getters.status.board"
      v-bind:key="i"
    >
      <Cell :value="cell" :index="i"></Cell>
    </div>
  </div>
</template>

<script>
import Cell from "@/components/Cell.vue";

export default {
  name: "board",
  components: {
    Cell
  }
};
</script>

<style scoped>
.board {
  display: flex;
  flex-wrap: wrap;
}
.cell {
  flex: 1 0 30%; /* explanation below */
  border: 1px solid;
}
</style>
