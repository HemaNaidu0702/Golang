package main

import "fmt"

func main() {
	// var card string = "Ace od spades"

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Ace of Hearts")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
