package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println("---", i, "---")
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("x = %v & y = %v\n", x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both less than 4.")
		case x > 6 && y > 6:
			fmt.Println("Both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("X : Greater than equal to 4 and Less than equal to 6")
		case y != 5:
			fmt.Println("Y : Is not 5")
		default:
			fmt.Println("None of the previous cases were met.")
		}
	}
}
