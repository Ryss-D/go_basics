package main

import "fmt"

type shape interface{
    getArea() float64
}
type square struct{
    sideLength float64
}
type triangle struct{
    base float64
    height float64
}

func main () {
    triangulo := triangle{1, 2}
    cuadrado := square{2}
    printArea(triangulo)
    printArea(cuadrado)

}

func (t triangle) getArea() float64 {
    area:= t.base * t.height * 0.5
    return area
}

func (s square) getArea() float64 {
    area:= s.sideLength * s.sideLength
    return area
}

func printArea(s shape){
    fmt.Println("The area is", s.getArea())
}
