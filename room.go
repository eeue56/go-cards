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

func (room *Room) Dealer() Player {
    return room.Players[room.dealer]
}

func (room *Room) NonDealers() chan Player {
     c := make(chan Player)

     go func(count int, players []Player, dealerNumber int) {
        for i := 0; i < count; i++ {
            if i == dealerNumber {
                continue
            }

            c <- players[i]
        }

        close(c)
     }(len(room.Players), room.Players, room.dealer)

     return c
}

func (room *Room) RotateDealer() {
    if room.dealer + 1 == len(room.Players) {
        room.dealer = 0
    } else {
        room.dealer++
    }
}