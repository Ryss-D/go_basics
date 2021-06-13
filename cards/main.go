package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	hand, remainingDeck := deal(cards, 3) // hand, reaminingDeck means that we are defining two variables

	cards.print()
	fmt.Println("----- end of deck -----")

	hand.print()
	fmt.Println("----- end of hand -----")
	remainingDeck.print()

	fmt.Println(cards.toString())

	cards2 := readFromFile("my_cards")
	cards2.print()

	cards2.shuffle()
	cards2.print()
}

// to convert types on go we use de notation 1(2) where 1 is the type we want an 2 the value we have ieg. []byte("hi")
