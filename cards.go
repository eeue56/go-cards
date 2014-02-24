package gocards

type Card struct {
    Text    string
}

type BlackCard struct {
    Card
}

type WhiteCard struct {
    Card
}

func (card *Card) Equal (other Card) bool {
    return card.Text == other.Text
}

func (card *BlackCard) Equal (other BlackCard) bool {
    return card.Text == other.Text
}

func (card *WhiteCard) Equal (other WhiteCard) bool {
    return card.Text == other.Text
}