package main

import "fmt"

func main() {
	//SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v \n y=%v\n", x, y)

	// CONDITIONAL
	// if statements
	// switch statements

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}

	if x < 42 {
		fmt.Println("Less than the meaning of life")
	} else if x == 42 {
		fmt.Println("equal to the meaning of life")
	} else {
		fmt.Println("greater than the meaning of life")
	}

	/*
		"If" statements specify the conditional execution of two branches
		according to the value of a boolean expression. If the expression evaluates
		to true, the "if" branch is executed, otherwise, if present, the "else" branch is executed.
	*/
	// https://go.dev/ref/spec#If_statements

	/*
		Comparison operators
		Comparison operators compare two operands and yield an untyped boolean value.

		==    equal
		!=    not equal
		<     less
		<=    less or equal
		>     greater
		>=    greater or equal
	*/

	// https://go.dev/ref/spec#Comparison_operators

	if x < 42 && y < 42 {
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 {
		fmt.Println("x is not 42")
	}

	/*
		Logical operators
		Logical operators apply to boolean values
		and yield a result of the same type as the operands.
		The right operand is evaluated conditionally.

		&&    conditional AND    p && q  is  "if p then q else false"
		||    conditional OR     p || q  is  "if p then true else q"
		!     NOT                !p      is  "not p"
	*/
	// https://go.dev/ref/spec#Logical_operators

	/*
		The expression [evaluated in an if statement ]may be preceded by a simple statement,
		which executes before the expression is evaluated.
	*/
	// https://go.dev/ref/spec#If_statements
	/*
		The comma ok idiom is also like this
	*/
	// https://go.dev/play/p/OXGzjxVkag0
}