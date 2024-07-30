package main

import "fmt"

var a int

func main() {
	println(a)
	b := 42
	println(b)
	c, d := 10, "Ten"
	println(c, d)
	var e float32 = 45.5
	fmt.Printf("%v is of type %T\n", e, e)
	f, g, _ := 5, 6, 7
	fmt.Println(f, g,)
}
