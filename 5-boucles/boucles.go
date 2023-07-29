package main

import "fmt"

func main() {
	//l'ecriture basique tu connais
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	//autre ecriture (equivalent du while)
	x := 10 //x vaut 10

	for x >= 0 { //tant que x est sup ou egal a 0
		fmt.Println(x) //tu affiche x
		x--            //et tu diminue de 1s
	}

	//autre facon de faire avec le break
	y := 10
	for {
		if y >= 15 { //on verra pas 15 puisque quand y sera egal a 15 il va casser la boucle
			break
		}
		fmt.Println(y)
		y++
	}
	//autre facon avec le continue
	z := 0
	for ; z <= 10; z++ {
		if z%2 == 1 {
			continue //on arrete la boucle mais on reprend aud ebut de cette derniere
		}
		fmt.Println(z)
	}
	/*On a dabord z qui vaut 0. au premier tour on a 0 on verifie si 0 est impaire (avec z%2==1) et bah c'est pas
	le cas, donc on println le z et on incremente de 1. Là z vaut 1 et on va reverifier si 1 est impaire.
	C'est le cas donc on atterit dans le continue qui va recommencer le boucle au debut. c-a-d va pas
	vvenir jusquau println il va juste incrementer donc z vaut 2 et in recommence au debut et rebelote*/

}
