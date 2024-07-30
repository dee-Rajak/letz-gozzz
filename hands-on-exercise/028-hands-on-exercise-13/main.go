package main

import "fmt"

func main() {
	var a int8 = 127
	var b uint8 = 255
	fmt.Printf("%T \t\t %d \t\t %b\n", a, a, a)
	fmt.Printf("%T \t\t %d \t\t %b\n", b, b, b)
}
