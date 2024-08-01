package main

import "fmt"

func main() {
	// xi := []int{42, 43, 44, 45, 46, 47}
	// for i, v := range xi {
	// 	fmt.Println("Index: ", i, "\tValue: ", v)
	// }
	m := map[string]int{
		"Dee":        10,
		"James":      42,
		"Moneypenny": 32,
	}
	// for k, v := range m {
	// 	fmt.Println("Key: ", k, "\tValue: ", v)
	// }
	age := m["James"]
	fmt.Println(age)
	if value, ok := m["Dee"]; ok {
		fmt.Println("It is stored in the Map.", " Value: ", value)
	} else {
		fmt.Println("It is not stored in the Map.")
	}
}
