package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func getNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2 //31
	mul := num1 * num2 //210

	return sum, mul
}

func main() {
	a := 10
	b := 21
	// sum := add(a, b)
	p, q := getNumbers(a, b)

	fmt.Println(p)
	fmt.Println(q)
}
