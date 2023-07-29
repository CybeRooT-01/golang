package main

import "fmt"

type exemple struct {
	a string
	b int
	c bool
}

func main() {
	myexemple := exemple{
		a: "Birahim",
		b: 21,
		c: true,
		//cette facon pas besoin de tout specifier il va mettre les int a 0 et les bools a false par defaut
	}
	myexemple2 := exemple{"Fallou", 24, false} //cette facon d'ecrire t'oblige a respecter tout les 3 donnés du struct
	fmt.Println(myexemple)
	fmt.Println(myexemple2)

	//la fonction importé dans contact.go
	myContact := newContact("Charo", 29)
	fmt.Println(myContact)
}
