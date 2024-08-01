package main

import (
	"fmt"
)

func main() {
	// for i := 1; i <= 100; i++ {
	// 	if x := rand.Intn(5); x == 3 {
	// 		fmt.Println(i, "--- ", "x is 3")
	// 	}
	// 	fmt.Println(i, "--- ")
	// }
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
