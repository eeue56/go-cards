package gocards

type Room struct {
    Players     []Player
    BlackHand   []BlackCard
    dealer      int
}

func NewRoom() *Room {
    room := Room{}
    room.dealer = 0
    room.Players = []Player{}
    room.BlackHand = []BlackCard{}

    return &room
}

func (room *Room) SetPlayers(players []Player) {
    room.Players = players
    room.dealer = 0
}

func (room *Room) AddPlayer(player Player) {
    dealer := room.Dealer()
    room.SetPlayers(append(room.Players, player))


    for player := range room.Players {
        if room.Players[player].Points == dealer.Points {
            room.dealer = player
            break
        }
    }
}

func (room *Room) Dealer() Player {
    return room.Players[room.dealer]
}

func (room *Room) NonDealers() chan Player {
     c := make(chan Player)

     go func(players []Player, dealerNumber int) {
        for i := 0; i < len(players); i++ {
            if i == dealerNumber {
                continue
            }

            c <- players[i]
        }

        close(c)
     }(room.Players, room.dealer)

     return c
}

func (room *Room) RotateDealer() {
    if room.dealer + 1 == len(room.Players) {
        room.dealer = 0
    } else {
        room.dealer++
    }
}