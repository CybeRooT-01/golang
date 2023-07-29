package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//generer un nombre aleatire faut utiliser le rand.seed sinon on aura tjr le mm nombre
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Int()) randint permet d'avoir un nbr aleatoire

	//on peut directement declarer une var dans le if (specificité de go) et on definit la condition
	if x := rand.Int(); x%2 == 0 {
		fmt.Println(x, "Est un nombre paire")
	} else {
		fmt.Println(x, "Est un nombre impaire")
	}

	//on peut aussi separer les choses y va prendre le resultat du modulo(reste de la division) et on va definir notre condition a part
	y := rand.Int() % 2

	if y == 0 {
		fmt.Println(y, "Est un nombre paire")
	} else {
		fmt.Println(y, "Est un nombre impaire")
	}

	// age

	age := 18
	if age < 18 {
		fmt.Println("Vous etes mineur")
	} else if age > 18 {
		fmt.Printf("Vous etes un adulte \n")
	} else {
		fmt.Printf("Vous avez pile 18 ans\n")
	}
}
