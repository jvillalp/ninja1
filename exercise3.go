package main

import (
	"fmt"
)

var a int = 42
var b string = "James Bond"
var c bool = true

func printpackagelevel() {
	s := fmt.Sprintf("%v\t%v\t%v", a, b, c)
	fmt.Println(s)
}

func divideBy(a int64, b int64) (int64, error) {
	if b == 0 {
		return -1, fmt.Errorf("cannot divide by zero %d/%d", a, b)
	}
	return a / b, nil
}
