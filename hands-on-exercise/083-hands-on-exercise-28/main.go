package main

import "fmt"

func main() {
	// x := 0
	// for x <= 50 {
	// 	if x%2 == 0 {
	// 		fmt.Println(x, " is even.")
	// 		x++
	// 		continue
	// 	}
	// 	fmt.Println("I am printed if number's odd.")
	// 	x++
	// }

	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println(i, " is an odd number.")
			continue
		}
		fmt.Println("For programmers, Learning curves are infinite loops.")
	}
}
