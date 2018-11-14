<template>
    <div v-on:click="play(index)">{{ value }}</div>
</template>

<script>
// @ is an alias to /src
// import api from "@/services/api.js";

export default {
    name: "cell",
    props: ["value", "index"],

    methods: {
        play(index) {
            if (this.value === '-' && !this.$store.getters.status.gameover)
                this.$store.dispatch('makeMove', { index })
        }

    }
};
</script>
<style scoped>
</style>