package gocards

import "testing"

func TestEqual(t *testing.T) {
    players := []Player{ Player{Points:1},
        Player{Points:2},
        Player{Points:3},
        Player{Points:4},
        Player{Points:5},
    }

    if players[0].Equal(players[1]) { 
        t.Errorf("Player 0 should not equal player 1!")
    }

    if !players[0].Equal(players[0]) { 
        t.Errorf("Player 0 should equal player 0!")
    }
}

func TestRemoveCard(t *testing.T) {

    hand := []WhiteCard{WhiteCard{Card{"hello"}},
        WhiteCard{Card{"dog"}}, 
    }
    
    player := Player{Points:1, Hand:hand}

    player.RemoveCard(hand[0])

    if len(player.Hand) == 2 {
        t.Errorf("Failed to remove card 0 from the hand")
    }

    player.RemoveCard(hand[1])

    if len(player.Hand) != 0 {
        t.Errorf("Failed to remove final card")
    }

}