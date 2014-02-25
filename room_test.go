package gocards

import "testing"

func TestNonDealers(t *testing.T) {
    players := []*Player{ &Player{Points:1},
        &Player{Points:5},
        &Player{Points:6},
        &Player{Points:8},
        &Player{Points:0},
    }

    room := NewRoom() 
    room.SetPlayers(players)

    for i := 0; i < len(players); i++ {
        room.RotateDealer()

        for p := range room.NonDealers() {
            if p.Equal(*room.Dealer()) {
                t.Errorf("Player should not be in dealers! %v", p)
            }
        }
    }

    room.SetPlayers([]*Player{})

    for _ = range room.NonDealers() {
        t.Errorf("Should not be in here!")
    }

    room.SetPlayers([]*Player{&Player{Points:1}})

    for _ = range room.NonDealers() {
        t.Errorf("Should not be in here!")
    }
}

func TestRotateDealers(t *testing.T) {
    players := []*Player{ &Player{Points:1},
        &Player{Points:2},
        &Player{Points:3},
        &Player{Points:4},
        &Player{Points:5},
    }

    room := NewRoom() 
    room.SetPlayers(players)

    if !room.Dealer().Equal(*players[0]) {
        t.Errorf("Incorrect dealer set as default")
    }


    for i := 1; i < len(players); i++ {
        room.RotateDealer()

        if !room.Dealer().Equal(*players[i]) {
            t.Errorf("Expected dealer with points %v, got dealer with points %v", 
                players[1],
                room.Dealer().Points)
        }
    }

    room.RotateDealer()

    if !room.Dealer().Equal(*players[0]) {
        t.Errorf("Incorrect dealer set as default")
    }
}

func TestAddPlayer(t *testing.T) {
    players := []*Player{ &Player{Points:1},
        &Player{Points:2},
        &Player{Points:3},
        &Player{Points:4},
        &Player{Points:5},
    }

    room := NewRoom() 
    room.SetPlayers(players)
    room.RotateDealer()

    room.AddPlayer(Player{Points:6});

    if (room.Players[len(room.Players) - 1].Points != 6){
        t.Errorf("Failed to add new player correctly")
    }

    if (room.Dealer().Points != 2){
        t.Errorf("Failed to correctly keep track of dealer when adding")
    }
}

func TestRemovePlayer(t *testing.T) {
    players := []*Player{ &Player{Points:1},
        &Player{Points:2},
        &Player{Points:3},
    }

    room := NewRoom() 
    room.SetPlayers(players)
    room.RotateDealer()

    room.RemovePlayer(*players[1])

    if (len(room.Players) != 2 || room.Players[1].Points == 2){
        t.Errorf("Failed to remove player correctly")
    }

    if (room.dealer != 1 || room.Dealer().Points != 3){
        t.Errorf("Incorrect reassign of dealer, expected dealer id of 1. Dealer was %v",
            room.dealer)
    }


    room.RemovePlayer(*players[2])

    if (len(room.Players) != 1 || room.Players[0].Points != 1){
        t.Errorf("Failed to remove player correctly")
    }

    if room.dealer != 0 {
        t.Errorf("Incorrect reassign of dealer, expected dealer id of 0, got %v",
            room.dealer)
    }

    room.RemovePlayer(*players[0])

    if len(room.Players) != 0 {
        t.Errorf("Failed to remove player correctly")
    }

    if room.dealer != -1 {
        t.Errorf("Incorrect reassign of dealer, expected dealer id of -1, got %v",
            room.dealer)
    }
}

func TestDealer(t *testing.T) {
    players := []*Player{}

    room := NewRoom() 


    if room.dealer != -1 {
        t.Errorf("Failed to correctly set dealer to -1")
    }

    room.SetPlayers(players)

    if room.dealer != -1{
        t.Errorf("Failed to correctly set dealer to -1 when using SetPlayers")
    }

    players = append(players, &Player{Points:5})

    room.SetPlayers(players)

    if room.dealer != 0{
        t.Errorf("Failed to correctly set dealer to 0 when using SetPlayers")
    }
}

func TestDealBlackCard(t *testing.T) {
    room := NewRoom()

    if room.DealBlackCard() != nil {
        t.Errorf("Failed to return nil when BlackHand empty")
    }

    blacks := []*BlackCard{ &BlackCard{ Card{"hello"}}}

    room.BlackHand = blacks

    if room.DealBlackCard() != blacks[0] {
        t.Errorf("Incorrect card dealt!")
    }

    if len(room.BlackHand) != 0 {
        t.Errorf("Failed to remove card from deck")
    }
}