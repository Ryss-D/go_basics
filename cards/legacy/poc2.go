package poc2

// import "fmt"

// func main() {
// 	card := newCard()
// 	cards := []string{newCard(), newCard()} //that how we declase a slice
// 	cards = append(cards, newCard())        // that how we append elements, but actually it not modify the variable, create a new value and reasign it to variable

// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	} // every time we iterate we are throwing away the last index

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}

// 	fmt.Println(card)
// 	fmt.Println(cards)
// }

// func newCard() string { //the string changes de default value of void fuction to a gunction that returns strings

// 	return "Five of Diamonds"
// }

// we can create multiple files linked to main package and they will compile

// go has two "LIST LIKE" data structures "Array" == Fixed length list of things and "Slice" == An array that can grow or shrink
// Every item on slice must be of same type
//next a way to emulate a while loop
// for{
//do something
// if !contition {
//break
// }
// }
