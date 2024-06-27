package main

func main() {
	// var card string = "Ace od spades"

	// cards := newDeck()

	// hand, remainingcards := deal(cards, 3)
	// hand.print()
	// remainingcards.print()

	// cards := newDeck()
	// cards.savetofile("mycards")

	// cards := newDeckfromfile("my")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
