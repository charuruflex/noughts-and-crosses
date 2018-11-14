# noughts-and-crosses

Nougths & crosses in go, playable via API

## requirements
  - go (tested with go1.11.2)
  - dep (for dependencies, tested with v0.5.0)
  - docker and docker-compose (optional, tested with docker 18.03.1-ce and docker-compose 1.21.0)
  
## How to start (without docker)
  - clone this repository
  - `dep ensure`
  - `go build -o nac main.go`
  - `PORT=8080 ./nac`
  
The API can be accessed on http://localhost:8080
The board size defaults to 3x3 and can be changed with env var `SIZE`

## API
  - GET `/api/v1/status`: get game state, the winner, the next player, if the game has ended
    - res: ```{
    "gameover": false,
    "winner": null,
    "nextplayer": "P1",
    "board": [
        "-",
        "-",
        "-",
        "-",
        "-",
        "-",
        "-",
        "-",
        "-"
    ],
    "players": null
}```
  - POST `/api/v1/makemove`: make a move for a player
    - req: ```{
	"player": "P1",
	"index": 1
}``` with 0 < index < 8 (if `SIZE=3`)

The board looks like this (if `SIZE=3`):

  0  1  2

  3  4  5

  6  7  8

  - GET `/api/v1/newgame`: reset the game

## Front (tbd)
Very basic front end in Vue.js

## docker (tbd)

## How to play (tbd)
