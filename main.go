package main

func main() {
	//cards := newDeckFromFile("deck.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
