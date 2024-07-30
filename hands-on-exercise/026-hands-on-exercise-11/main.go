package main

import "fmt"

func main() {
	var a int = 42
	var b float64 = 42.24
	var c string = "Hello Love"

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	fmt.Printf("%v is of type %T\n", c, c)
}
