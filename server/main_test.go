package main

import "testing"

func TestVictoryCases(t *testing.T) {
	initialInit()
	tables := [][]move{
		{
			{Index: 0, Player: "P1"},
			{Index: 3, Player: "P2"},
			{Index: 1, Player: "P1"},
			{Index: 4, Player: "P2"},
			{Index: 2, Player: "P1"},
		},
		{
			{Index: 3, Player: "P1"},
			{Index: 0, Player: "P2"},
			{Index: 4, Player: "P1"},
			{Index: 1, Player: "P2"},
			{Index: 5, Player: "P1"},
		},
		{
			{Index: 6, Player: "P1"},
			{Index: 0, Player: "P2"},
			{Index: 7, Player: "P1"},
			{Index: 1, Player: "P2"},
			{Index: 8, Player: "P1"},
		},
		{
			{Index: 0, Player: "P1"},
			{Index: 1, Player: "P2"},
			{Index: 3, Player: "P1"},
			{Index: 2, Player: "P2"},
			{Index: 6, Player: "P1"},
		},
		{
			{Index: 1, Player: "P1"},
			{Index: 0, Player: "P2"},
			{Index: 4, Player: "P1"},
			{Index: 2, Player: "P2"},
			{Index: 7, Player: "P1"},
		},
		{
			{Index: 2, Player: "P1"},
			{Index: 0, Player: "P2"},
			{Index: 5, Player: "P1"},
			{Index: 1, Player: "P2"},
			{Index: 8, Player: "P1"},
		},
		{
			{Index: 0, Player: "P1"},
			{Index: 1, Player: "P2"},
			{Index: 4, Player: "P1"},
			{Index: 2, Player: "P2"},
			{Index: 8, Player: "P1"},
		},
		{
			{Index: 2, Player: "P1"},
			{Index: 0, Player: "P2"},
			{Index: 4, Player: "P1"},
			{Index: 1, Player: "P2"},
			{Index: 6, Player: "P1"},
		},
	}

	for _, moves := range tables {
		for _, move := range moves {
			err := makeMove(move.Index, move.Player)
			if err != nil {
				t.Error(err)
			}
		}
		if !status.Gameover || status.Winner == nil || !(*status.Winner == "P1") {
			t.Errorf("error with winning moves %+v\n", moves)
		}
		initAll()
	}
}

func TestMakeMoveErrors(t *testing.T) {

}
