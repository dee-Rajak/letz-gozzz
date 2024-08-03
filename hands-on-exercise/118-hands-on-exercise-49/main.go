package main

import "fmt"

func main() {
	x := make(map[string][]string)
	x[`bond_james`] = []string{"shaken, not stirred", "martinis", "fast cars"}
	x["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	x["no_dr"] = []string{`cats`, `ice cream`, `sunsets`}

	fmt.Printf("%#v\n", x)

	for k, v := range x {
		fmt.Println("---: ", k, " :---")
		for i, a := range v {
			fmt.Println(i, a)
		}
	}

	delete(x, "bond_james")

	for k, v := range x {
		fmt.Println("---: ", k, " :---")
		for i, a := range v {
			fmt.Println(i, a)
		}
	}

	// m := map[string][]string{
	// 	"A": []string{"", ""},
	// }
}
