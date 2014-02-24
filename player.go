package gocards

type Player struct {
    Hand    []WhiteCard
    Points  int
}

func (player Player) Equal(other Player) bool {

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