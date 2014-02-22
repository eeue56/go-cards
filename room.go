package gocards

type Room struct {
    Players     []Player
    dealer      int
    BlackHand   []BlackCard
}

func (room *Room) Dealer() Player {
    return room.Players[room.dealer]
}

func (room *Room) RotateDealer() {
    if room.dealer + 1 == len(room.Players) {
        room.dealer = 0
    } else {
        room.dealer++
    }
}