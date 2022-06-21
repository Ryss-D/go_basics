package main

import "fmt"

type bot interface {
	getGreeting() string
	//any type that succes the condition of have a function getGreeting() and return a string now can access to fuctions
	//of type bot something like a abstract class
}

// the default interface sintax is
// type bot interface {
//  function (type of  arguments) (type of retuns)
// on go we can talk about concrete types and interfaces types(we cant create interface types from scratch or directrly)
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//func (englishBot) will by totatlly equivalent
	//TO DO VERY CUSTOM ENGLISH LOGIC
	return "Hello there"
}

func (sb spanishBot) getGreeting() string {
	//func (sb spanishBot) will by totatlly equivalent
	//TO DO VERY CUSTOM SPANISH LOGIC
	return "Hola"
}
