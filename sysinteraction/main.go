package main

import (
    "fmt"
    "os"
    "io"
)

func main() {
    fileName := os.Args[1]

    file, err := os.Open(fileName)
    if err != nil{
        fmt.Println("we hava a erro", err)
        os.Exit(1)
    }

    io.Copy(os.Stdout, file)

}
