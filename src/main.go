package main

import (
	"fmt"
	ds "lib/dataStructure"
	dsl "lib/dataStructure/linkedList"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func main() {

	fmt.Println("fibonacci : ",
		ds.Fibo_recursion(4))

	list := dsl.New()

	var i Shape
	i = &Circle{5}
	i.Area()

	fmt.Println("Print i : ", i)

	list.InsertFirst(dsl.NewNode(&Circle{4}))

	fmt.Println("Print list : ", list)
}
