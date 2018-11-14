# noughts-and-crosses

Nougths & crosses in go, playable via API.

## Requirements
  - go (tested with go1.11.2)
  - dep (for dependencies, tested with v0.5.0)
  - node (tested with v11.1.0)
  - npm or yarn (tested with yarn 1.12.3)
  - docker and docker-compose (optional, tested with docker 18.03.1-ce and docker-compose 1.21.0)
  
## How to start (without docker)

### Server (go)
  - `cd server`
  - `dep ensure`
  - `go build -o game main.go`
  - `./game`
  
The API can be accessed on http://localhost:8080, the port can be changed via `PORT` env var. 
The board size defaults to 3x3 and can be changed with env var `SIZE`.

### Client (Vue.js)
  - `cd client`
  - `yarn` or `npm -i`
  - edit .env file to configure accordingly with your API (defaults to http://localhost:8080/api/v1/)
  - `yarn serve` or `npm run serve` (dev mode)

You can play now by accessing the given URL.

## API
  - GET `/api/v1/status`: get game state, the winner, the next player, if the game has ended
    - res: ```{"gameover":false,"winner":null,"nextplayer":"P2","board":["P1","-","-","P2","-","-","P1","-","-"],"players":null,"movescounter":3}```
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

## Run in containers (docker & docker-compose needed)
  - `docker-compose build && docker-compose up -d`

Et voilà! You can play Noughts & Crosses via http://localhost:8080

## Demo
- A live demo can be found here: https://noughts-and-crosses.charuruflex.eu/
- API : https://noughts-and-crosses.charuruflex.eu/api/v1/status
