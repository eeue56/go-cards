package gocards

import "testing"

func TestNonDealers(t *testing.T) {
    players := []Player{ Player{Points:1},
        Player{Points:5},
    }

    room := NewRoom() 
    room.Players = players

    for p := range room.NonDealers() {
        if p.Points == room.Dealer().Points {
            t.Errorf("Player should not be in dealers! %v", p)
        }
    }

    room.Players = []Player{}

    for _ = range room.NonDealers() {
        t.Errorf("Should not be in here!")
    }
}