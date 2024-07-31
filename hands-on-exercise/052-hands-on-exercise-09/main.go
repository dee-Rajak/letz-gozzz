package main

import "fmt"

var (
	a int
	b string
)

const (
	c float64 = 42.42
	d string  = "fourty two"
)

func main() {
	e := "Hey Boi"
	a = 42
	b = "FOURTY TWO"
	fmt.Printf("%v\n%v\n%v\n%v\n%v\n", e, a, b, c, d)
}
