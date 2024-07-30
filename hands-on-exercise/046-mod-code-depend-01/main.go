package main

import (
	"fmt"

	"github.com/dee-Rajak/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())
}
