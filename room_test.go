package gocards

import "testing"

func TestNonDealers(t *testing.T) {
    players := []Player{ Player{Points:1},
        Player{Points:5},
        Player{Points:6},
        Player{Points:8},
        Player{Points:0},
    }

    room := NewRoom() 
    room.SetPlayers(players)

    for i := 0; i < len(players); i++ {
        room.RotateDealer()

        for p := range room.NonDealers() {
            if p.Points == room.Dealer().Points {
                t.Errorf("Player should not be in dealers! %v", p)
            }
        }
    }

    room.SetPlayers([]Player{})

    for _ = range room.NonDealers() {
        t.Errorf("Should not be in here!")
    }

    room.SetPlayers([]Player{Player{Points:1}})

    for _ = range room.NonDealers() {
        t.Errorf("Should not be in here!")
    }
}

func TestRotateDealers(t *testing.T) {
    players := []Player{ Player{Points:1},
        Player{Points:2},
        Player{Points:3},
        Player{Points:4},
        Player{Points:5},
    }

    room := NewRoom() 
    room.SetPlayers(players)

    if room.Dealer().Points != players[0].Points {
        t.Errorf("Incorrect dealer set as default")
    }


    for i := 1; i < len(players); i++ {
        room.RotateDealer()

        if room.Dealer().Points != players[i].Points {
            t.Errorf("Expected dealer with points %v, got dealer with points %v", 
                players[1],
                room.Dealer().Points)
        }
    }

    room.RotateDealer()

    if room.Dealer().Points != players[0].Points {
        t.Errorf("Incorrect dealer set as default")
    }

}