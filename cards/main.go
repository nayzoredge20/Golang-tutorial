package main

func main() {
	// creates a new cards and assigns that to old cards
	//cards.print()
	// cards := newDeck()
	// cards.shuffle()
	// cards.print()
	new_cards := newDeckFromFile("my_cards")
	new_cards.print()
}
