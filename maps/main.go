package main

import "fmt"

func main() {
    //var colors map[string]string it creates a empty maps


    //    colors := make(map[string]string)

     //   colors["white"] = "#ffffff"
// here we cant acces toelements using structName.white sintaxys

    //delete(colors, "white")//that fuctions delete de element " white" from colors map

    color := map[string]string{// maps on go are static typed, it means that all the keys must have the same type and all the values must have the same type but it dont have to be the same
    //in this particular case keys ([string]) are string type and values too
    "red": "#ff0000",
    "green": "3bf745",
    "white": "ffffff",
    }

    printMap(color)
}

func printMap(c map[string]string) {
    for color, hex := range c{
        fmt.Println("for key",color, "the value is ",hex)

    }
}
