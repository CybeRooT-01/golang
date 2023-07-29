package main

import "fmt"

func main() {
	//fmt.Printf("Hello World") on va afficher hello world
	var (
		x int
		y string
		z bool
	)
	/*
		var x int
		x = 19
		y := 16
		alors on va pouvoir utilisé sa mais la declaration + assignation ne marche pas en dehor d'une fonction donc
		on va tout regrouper dans un var
	*/
	x = 21
	y = "Birahim"
	z = true
	//fmt.Printf("Mon age est: %v ! ", x)
	//fmt.Printf("\nMon nom est: %v ! ", y)
	//fmt.Printf("\nJ'ai plus de 18 ans: %v ! ", z)
	//Thier tu peux comprendre sa. ta fais du C regarde bien wsh
	//on declare nos variable
	//on les assigne
	//et on lesaffiche avec fmt.printf
	fmt.Printf("Mon nom est %v, Mon age est %vans et j'ai plus de 18ans est %v", y, x, z)
}
