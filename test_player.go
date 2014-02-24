package gocards

import "testing"

func TestEqual(t *testing.T) {
    players := []Player{ Player{Points:1},
        Player{Points:5},
        Player{Points:6},
        Player{Points:8},
        Player{Points:0},
    }

    if players[0].Equal(players[1]) { 
        t.Errorf("Player 0 should not equal player 1!")
    }

    if !players[0].Equal(players[0]) { 
        t.Errorf("Player 0 should equal player 0!")
    }
}