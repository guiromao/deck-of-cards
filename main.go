package main

func main() {

	//cards := newDeck()

	cards := newDeckFromFile("deck.txt")
	cards.print()

}
