package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var players []player
var status gameStatus
var config gameConfig

type gameConfig struct {
	Size int
	Port string
}

type player struct {
	Name     string
	Symbol   string
	Counters []int
}

type move struct {
	Player string `json:"player"`
	Index  int    `json:"index"`
}

type gameStatus struct {
	Gameover     bool     `json:"gameover"`
	Winner       *string  `json:"winner"`
	NextPlayer   string   `json:"nextplayer"`
	Board        []string `json:"board"`
	Players      []player `json:"players"`
	MovesCounter int      `json:"movescounter"`
}

type nAndCErr struct {
	Msg string `json:"error"`
}

func (e *nAndCErr) Error() string {
	return e.Msg
}

func initialInit() {
	config = gameConfig{Size: 3, Port: fmt.Sprintf(":%d", 8080)}
	if sizeEnv, ok := os.LookupEnv("SIZE"); ok {
		config.Size, _ = strconv.Atoi(sizeEnv)
	}
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		config.Port = ":" + portEnv
	}

	players = []player{
		player{Name: "P1", Symbol: "P1"},
		player{Name: "P2", Symbol: "P2"},
	}
	initAll()
}

func initAll() {
	status = gameStatus{Gameover: false, Winner: nil, Board: initBoard(config.Size)}
	players[0].Counters = make([]int, config.Size*2+2)
	players[1].Counters = make([]int, config.Size*2+2)
	status.NextPlayer = "P1"
}

func initBoard(boardSize int) (board []string) {
	board = make([]string, config.Size*config.Size)
	for i := 0; i < config.Size*config.Size; i++ {
		board[i] = "-"
	}
	return
}

func printBoard() {
	var boardStr string
	for i := 0; i < config.Size; i++ {
		boardStr += fmt.Sprintf("%s\n", strings.Join(status.Board[i*config.Size:(i+1)*config.Size], " "))
	}
	fmt.Println(boardStr)
}

func jsonBoard() []byte {
	return []byte(fmt.Sprintf(`["%s"]`, strings.Join(status.Board, `","`)))
}

func updateCounters(index int, player player) bool {
	player.Counters[index]++
	victory := player.Counters[index] >= config.Size

	if victory {
		fmt.Println("WINNER!")
		status.Winner = &player.Name
		status.Gameover = true
		status.NextPlayer = ""
	}
	return victory
}

func makeMove(index int, playerName string) (err error) {

	var currentPlayer, otherPlayer player

	for _, p := range players {
		if p.Name == playerName {
			currentPlayer = p
		} else {
			otherPlayer = p
		}
	}

	switch {
	case cmp.Equal(currentPlayer, player{}):
		err = &nAndCErr{"Given player not valid"}
	case index < 0 || index >= config.Size*config.Size:
		err = &nAndCErr{"Given index out of bounds"}
	case status.Gameover:
		err = &nAndCErr{"Game has ended"}
	case currentPlayer.Name != status.NextPlayer:
		err = &nAndCErr{"It's not your turn!"}
	case status.Board[index] != "-":
		err = &nAndCErr{"Move already played before"}
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	row, col := index/config.Size, index%config.Size
	fmt.Printf("new move for %s: row %d, col %d\n", currentPlayer.Name, row, col)
	status.Board[index] = currentPlayer.Symbol
	status.NextPlayer = otherPlayer.Name
	printBoard()

	if row == col {
		updateCounters(0, currentPlayer)
	}
	if row == config.Size-1-col {
		updateCounters(1, currentPlayer)
	}
	updateCounters(2+row, currentPlayer)
	updateCounters(2+config.Size+col, currentPlayer)

	status.MovesCounter++

	if status.Winner == nil && status.MovesCounter == config.Size*config.Size {
		status.Gameover = true
	}
	return
}

func about(w http.ResponseWriter, r *http.Request) {
	blabla := "----- Noughts & Crosses -----"
	w.Write([]byte(blabla))
}

func setMove(w http.ResponseWriter, r *http.Request) {
	// TODO: need more checking (type, value)
	// TODO: need to handle errors
	body, _ := ioutil.ReadAll(r.Body)
	var m move
	_ = json.Unmarshal(body, &m)

	err := makeMove(m.Index, m.Player)

	var res []byte
	if err != nil {
		res, _ = json.Marshal(err)
		w.WriteHeader(http.StatusNotAcceptable)
	}
	w.Write(res)
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	statusJSON, _ := json.Marshal(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(statusJSON)
}

func newGame(w http.ResponseWriter, r *http.Request) {
	initAll()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"newGame":"ok"}`))
}

func main() {
	fmt.Printf("----- Noughts & Crosses -----\n\n")
	initialInit()

	router := mux.NewRouter()
	router.HandleFunc("/", about).Methods("GET")
	router.HandleFunc("/api/v1/makemove", setMove).Methods("POST")
	router.HandleFunc("/api/v1/status", getStatus).Methods("GET")
	router.HandleFunc("/api/v1/newgame", newGame).Methods("GET")

	printBoard()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	log.Fatal(http.ListenAndServe(config.Port, handlers.CORS(corsObj, headersOk)(router)))
}
