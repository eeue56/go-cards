package gocards

type Player struct {
    Hand    []WhiteCard
    Points  int
}

func (player *Player) Equal(other Player) bool {

    if player.Points != other.Points {
        return false
    }

    for i := range player.Hand {
        if player.Hand[i] != other.Hand[i] {
            return false
        }
    }

    return true
}

func (player *Player) RemoveCard(card WhiteCard) {
    hand := []WhiteCard{}

    var currentCard WhiteCard

    for i := range player.Hand {
        currentCard = player.Hand[i]
        if !card.Equal(currentCard) {
            hand = append(hand, currentCard)
        }
    }

    player.Hand = hand
}