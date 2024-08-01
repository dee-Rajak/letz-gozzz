package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization for my program occurs")
}

func main() {
	x := rand.Intn(300)
	fmt.Println("Random Number: ", x)

	// if x >= 0 && x <= 100 {
	// 	fmt.Println("Between 0 and 100")
	// } else if x >= 101 && x <= 200 {
	// 	fmt.Println("Between 101 and 200")
	// } else if x >= 201 && x <= 250 {
	// 	fmt.Println("Between 201 and 250")
	// } else {
	// 	fmt.Println("Not inside the defined range")
	// }

	switch {
	case x <= 100 && x >= 0:
		fmt.Println("Between 0 and 100")
	case x <= 200 && x >= 101:
		fmt.Println("Between 101 and 200")
	case x <= 250 && x >= 201:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("Not inside the defined range")
	}
}
