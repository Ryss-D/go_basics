package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of ' deck'
// which is a slice of strings

type deck []string //that []strings means taht deck extendes []string

func newDeck() deck {
	cardSuits := []string{"Spades", "diamonds", "Hearts", "Clubs"} //to acces slice we can do it like on python
	cardValues := []string{"As", "Two", "Theree", "Four"}

	cards := deck{}

	for _, suit := range cardSuits { //using underscore "_" makes go now that we now that here is a variable but we will not use it (in this particular time if we didnt use _ we will have ad erro ron termianl trying to run iti)
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() { // (d deck) act as a reciver, that means that any variable of type deck now gets access to the "print" method
	// on (1, 2) reciver 2 is the type and 1 is a reference of the instance were we are working, it means the curred 2 type object(or variable) that we are using something like use this reserved word
	// by convencion on reciver we use one or two letters of current type
	for i, card := range d {

		fmt.Println(i, card)

	}
}

func deal(d deck, handSize int) (deck, deck) { //(deck, deck) => able us to return multiple values and spicify this type
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //function structure is func WriteFile(filename string, data []byte, perm os.FileMode) error
}

func readFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	//null in go its represented as nil
	if err != nil {
		// Option #1 - log the error and renurt a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1) //default value for argument of os.Exit() function its - that means that us program run successfully

	} // else {

	//}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // here we are creating a new seed for random generator
	r := rand.New(source)                           //here we are creating a random generator with a random seed, making sure that the every randomizen value is different

	//thats because go use the same seed  for randomize by default

	for i := range d { //we can iterate just with index not always with elemsts and index
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
