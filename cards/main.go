package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
