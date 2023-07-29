package main

import "fmt"

func main() {
	//printf
	x, y := 10, 10
	fmt.Printf("Salut a tous \n")                        //fmt.printnf ne fais pas de saut de ligne donc faut mettre \n
	fmt.Printf("Ma variable vaut : %v", x)               //valeur de v
	fmt.Printf("Ma variable en base 2 vaut : %b \n", x)  //base 2 on utilise %b
	fmt.Printf("Ma variable en base 8 vaut : %o \n", x)  //bas eo octal onpeut ecrire %O (en maj) pour avoir le prefix octal 0o
	fmt.Printf("Ma variable en base 10 vaut : %v \n", x) //base 10 (c'est notre nombre par defaut)
	fmt.Printf("Ma variable en base 16 vaut : %X \n", x) //base hexadecimal on peut aussi ecrire %x en minuscul pour avoir en minuscule

	fmt.Printf(" x et y sont egaux : %t \n", x == y) //on teste si x == y

	fmt.Printf("l'ecriture scientifique de 12345.12345879 vaut %E \n", 12345.12345879) //ecriture scientifique

}
