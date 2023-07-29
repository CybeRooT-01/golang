package main

import "fmt"

var x int //portée globale

func main() {
	x = 5 //portée locale
	fmt.Println(x)
	f()
	//on aura 5 et 10 puisque dans main x vaut 5 et dans f x vaut 10
	showY()
}
func f() {
	x := 10 //portée locale
	fmt.Println(x)
}

// vscode gueule masi tout va bien. quand on execute sa passe
