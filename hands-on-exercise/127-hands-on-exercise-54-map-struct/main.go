package main

import "fmt"

type person struct {
	first string
	last  string
	favIC []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favIC: []string{"chocolate", "banana", "passion fruit with mango and guava"},
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		favIC: []string{"Strawberry", "Chocolate"},
	}

	m := make(map[string]person)
	m[p1.last] = p1
	m[p2.last] = p2

	for k, v := range m {
		fmt.Printf("%T\t%#v ---> %T\t%#v\n", k, k, v, v)
		fmt.Printf("%T\t%#v\n", v.first, v.first)
		fmt.Printf("%T\t%#v\n", v.last, v.last)
		fmt.Printf("%T\t%#v\n", v.favIC, v.favIC)
		fmt.Println("-----------------------")
	}

	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Println(p1.favIC)
	// fmt.Println(p2.favIC)

	// for _, v := range p1.favIC {
	// 	fmt.Println(p1.first, "favorite is", v)
	// }

	// for _, v := range p2.favIC {
	// 	fmt.Println(p2.first, "favorite is", v)
	// }
}
