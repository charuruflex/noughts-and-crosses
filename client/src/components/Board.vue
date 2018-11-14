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
  display: grid;
  grid-template-columns: auto auto auto;
  background-color: #2196f3;
  padding: 10px;
  width: 20em;
  margin: auto;
}
.cell {
  background-color: rgba(255, 255, 255, 0.8);
  border: 1px solid rgba(0, 0, 0, 0.8);
  padding: 20px;
  font-size: 30px;
  text-align: center;
}
</style>
