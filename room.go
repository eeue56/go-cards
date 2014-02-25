package gocards

type Room struct {
    Players     []*Player
    BlackHand   []*BlackCard
    dealer      int
}

func NewRoom() *Room {
    room := Room{}
    room.dealer = -1
    room.Players = []*Player{}
    room.BlackHand = []*BlackCard{}

    return &room
}

func (room *Room) SetPlayers(players []*Player) {
    room.Players = players

    if len(players) == 0 {
        room.dealer = -1
    } else {
        room.dealer = 0
    }
}

func (room *Room) AddPlayer(player Player) {
    dealer := room.Dealer()
    room.SetPlayers(append(room.Players, &player))


    for player := range room.Players {
        // TODO: add uniq for players
        if room.Players[player].Equal(*dealer) {
            room.dealer = player
            break
        }
    }
}

func (room *Room) RemovePlayer(player Player) {
    dealer := room.dealer

    players := []*Player{}

    for i := range room.Players {
        if !room.Players[i].Equal(player) {
            players = append(players, room.Players[i])
        } 
    }

    room.SetPlayers(players)


    if len(players) == 0 {
        dealer = -1
    } else if dealer + 1 > len(players) {
        dealer = 0
    }


    room.dealer = dealer
}

func (room *Room) Dealer() *Player {
    if room.dealer == -1 {
        return &Player{}
    }
    return room.Players[room.dealer]
}

func (room *Room) NonDealers() chan *Player {
     c := make(chan *Player)

     go func(players []*Player, dealerNumber int) {
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

func (room *Room) DealBlackCard() *BlackCard {
    if len(room.BlackHand) == 0 {
        return nil
    }

    black := room.BlackHand[0]
    room.BlackHand = room.BlackHand[1:]

    return black
}