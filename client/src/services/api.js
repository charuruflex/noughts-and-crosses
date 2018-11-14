import axios from "axios";

const client = axios.create({
  baseURL: process.env.VUE_APP_API,
  timeout: 2000,
  json: true,
  mode: "no-cors"
});

export default {
  status: () => client.get(`status`).then(res => res.data),
  newGame: () => client.get(`newgame`).then(res => res.data),
  makeMove: (index, player) => client.post(`makemove`, { index, player })
};
