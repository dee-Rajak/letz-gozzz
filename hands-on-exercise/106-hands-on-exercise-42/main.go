package main

import "fmt"

func main() {
	// x := [5]int{9, 8, 7}
	// x[3] = 6
	// x[4] = 5
	// for _, v := range x {
	// 	fmt.Printf("%v\t%#v\t%T\n", v, v, v)
	// }
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// for _, v := range x {
	// 	fmt.Printf("%v\t%#v\t%T\n", v, v, v)
	// }
	// fmt.Println(x[:5])
	// fmt.Println(x[5:])
	// fmt.Println(x[2:7])
	// fmt.Println(x[1:6])
	fmt.Println(x)

	// x = append(x, 52)
	// fmt.Println(x)
	// x = append(x, 53, 54, 55)
	// fmt.Println(x)
	// y := []int{56, 57, 58, 59, 60}
	// x = append(x, y...)
	// fmt.Println(x)
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
