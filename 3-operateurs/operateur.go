package main

import "fmt"

func main() {
	//declare nos vars
	var (
		x int
		y int
	)
	//assigner nos vars
	x = 15
	y = 15
	// operation: arithmethique + - * / %
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(x % y)
	//operateurs relationnelle
	fmt.Println(x == y)
	fmt.Println(x != y)
	fmt.Println(x < y)
	fmt.Println(x > y)
	fmt.Println(x <= y)
	fmt.Println(x >= y)
	//operateurs logique
	fmt.Println(x == y && x != y)
	fmt.Println(x == y || x != y)
}
