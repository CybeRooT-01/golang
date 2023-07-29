package main

import "fmt"

func main() {
	f := func() { //pas de nom pas de return et on stock directemetn dans une variable
		fmt.Printf("Je suis une fonction anonyme")
	}
	f() //on appelle juste f notre variable pourappeer la foncction

	//----------------fonction anonyme avec return-----------------------------

	j := func() string {
		fmt.Printf("Je suis une fonction anonyme avec return")
		return "Birahim"
	}() //on appelle directement la fonction (obligatoire)
	println(j) //le j return contient notre return

	//autre methode
	x, y := func() (int, int) {
		println("Aucun param deux return")
		return 5, 7
	}()
	println(x)
	println(y)

	//autre methode
	func(a, b int) {
		println("x+x*y*y = ", x+x*y*y)
	}(x, y) //on pouvais aussi mettre 2, 7 ou whatever là c juste qu'on avais deux variable qui traine autant les utiliser
}
