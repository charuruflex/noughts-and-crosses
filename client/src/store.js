import Vue from "vue";
import Vuex from "vuex";
import api from "@/services/api.js";
Vue.use(Vuex);

export default new Vuex.Store({
  // strict: process.env.NODE_ENV !== "production",
  state: {
    status: {},
    loading: false,
    error: null
  },
  getters: {
    loading: state => state.loading,
    error: state => state.error,
    status: state => state.status
  },
  mutations: {
    SET_STATUS(state, status) {
      state.status = status;
    },
    SET_LOADING(state, loading) {
      state.loading = loading;
    },
    SET_ERROR(state, error) {
      state.error = error;
    }
  },
  actions: {
    status({ commit }) {
      api
        .status()
        .then(res => {
          commit("SET_STATUS", res);
        })
        .catch(res => {
          commit("SET_ERROR", res.response.data.error);
        })
        .finally(() => {
          commit("SET_LOADING", false);
        });
    },
    makeMove({ dispatch, state, commit }, { index }) {
      state.loading = true;
      state.error = null;
      api
        .makeMove(index, state.status.nextplayer)
        .then(() => {
          dispatch("status");
        })
        .catch(res => {
          state.error = res.response.data.error;
          commit("SET_LOADING", false);
        });
    },
    newGame({ dispatch, state, commit }) {
      state.loading = true;
      state.error = null;
      api
        .newGame()
        .then(() => {
          dispatch("status");
        })
        .catch(res => {
          state.error = res.response.data.error;
          commit("SET_LOADING", false);
        });
    }
  }
});
