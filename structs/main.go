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
    //contact contactInfo  //we can add custom types to definition
    contactInfo// will act as as last line
}

func main() {
    //alex := person{"Alex","Anderson"}
    //alex := person{firstName:"Alex", lastName:"Anderson"}
    //fmt.Println(alex)
    //var alex person
    //go charges empty variables with zero default values string -> "", int -> 0, float -> 0, bool -> false


    jim := person{
        firstName: "Jim",
        lastName: "Party",
        //contact: contactInfo{
    contactInfo: contactInfo{
            email: "jim@gmail.com:",
            zipCode: 94000,
        },
    }

    //basic way
   // jimPointer := &jim//now jimPointer is a reference to memory addres, where are alocated jim, that because we use & operator
    //jimPointer.updateName("jimmy")
    //jim.print()
    //shortcut way
    jim.updateName("jimmy")
    jim.print()
}

//func (p person) updateName(newFirtsName string) {
 //   p.firstName = newFirtsName//here we are not updating the original structure we are updating a copy
    //created by go on other ram "slot" indentifyied with "p" that because go its a "pass by value" lenguaje
    //"pointers" solve this problem
//}

func (pointerToPerson *person) updateName(newFirstName string){
    (*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
    fmt.Printf("%+v", p) //v will print all name - value pairs on struct
  }

  // & is a opperator that give us th ememory address of the value this variable is pointing at, es decir give uss access to certain memory address who alocated a value
  // *pointer give us the value this memory address is pointing at, es decir give us access to certain value allocated on memory 
  //

  //that dont apply to slice vecause when we create one slice go by himself create two data structures a array taht reperesent the actual slice and a slice to use 
  // go data structures are dived into two types value types (that means that we need to use pointer to change these thing in a function)
  //{int, float, string, bool, structs} and referenced values (Dont worry about pointers with these){slices, maps, channels, pointers, functions}
  // will be nice check how the word ieg
  //slices contain [lenght, capacity, and pointer to array who represent his elements]
