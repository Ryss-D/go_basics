package main
// structs are like a python dictionary
import "fmt"

type contactInfo struct {
    email string
    zipCode int
}


type person struct {
    firstName string
    lastName string
}

func main() {
    //alex := person{"Alex","Anderson"}
    //alex := person{firstName:"Alex", lastName:"Anderson"}
    //fmt.Println(alex)
    var alex person
    //go charges empty variables with zero default values string -> "", int -> 0, float -> 0, bool -> false
    alex.firstName = "Alex"
    alex.lastName = "Anderson"

    fmt.Println(alex)
    fmt.Printf("%+v", alex)//%+v will print all name - value pairs on struct
}
