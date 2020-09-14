package main

import (
	"fmt"
)

type Human struct {
	age    int8
	name   string
	height Height
}
type Height struct {
	feet   int8
	inches int8
}

func printHuman(n string, a int8, f int8, i int8) {
	h := Height{
		feet:   f,
		inches: i,
	}
	human := Human{
		age:    a,
		name:   n,
		height: h,
	}
	fmt.Printf("%+v\n", human)
}

func printHello() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Printf("Hello World %d %s %t", x, y, z)
}

// type Driveable interface {

// }
