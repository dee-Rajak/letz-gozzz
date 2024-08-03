package main

import "fmt"

type person struct {
	f_name string
	l_name string
	fav_ic string
}

func main() {
	x := person{
		f_name: "dee",
		l_name: "nut",
		fav_ic: "barfi",
	}
	y := person{
		f_name: "sumi",
		l_name: "jaheen",
		fav_ic: "bomb",
	}

	fmt.Printf("%T ---> %#v\n", x, x)
	fmt.Printf("%T ---> %#v\n", y, y)

	z := []person{x, y}
	for _, v := range z {
		fmt.Println(v.fav_ic)
	}
}
