<template>
    <div>
        <div class="state">{{player}}</div>
        <div class="loading" v-if="$store.getters.loading">...loading...</div>
        <div class="error" v-if="$store.getters.error">{{$store.getters.error}}</div>
    </div>
</template>

<script>
// @ is an alias to /src

export default {
    name: "state",
    computed: {
        player: function () {
            if (this.$store.getters.status.gameover) {
                if (this.$store.getters.status.winner) {
                    return this.$store.getters.status.winner + " WINS!"
                }
                else {
                    return "DRAW!"
                }
            } else {
                return "next player: " + this.$store.getters.status.nextplayer
            }
        }
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