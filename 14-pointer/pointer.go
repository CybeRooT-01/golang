package main

import "fmt"

func UpdateA(pointerOfNumber *int, value int) {
	//*pointerOfNumber = 5
	*pointerOfNumber = value
	/* on va changer le type du parametre en pointeur et maintenant la valeur de ce pointeur (qui commence par un *) sera egal a 5
	et pour synamiser on ajoute un second paramettre qui va remplacer le 5. qui sera un entier pure */
}

func main() {
	number := 10
	fmt.Println(number) //10
	fmt.Println("Addrsse memoire de number :", &number)
	mypointer := &number                                                              //le pointeur doit pointer vers une cage memoire ou ya une variable donc on declare notre pointeur en tant que "&" etla variable auquel il va pointer
	fmt.Println(mypointer)                                                            //on aurra un chiffre bizzare: c'est le pointeur
	fmt.Printf("La valeur de l'adresse memoire %v vaut %d \n", mypointer, *mypointer) //pour retrouver la avaleur de la variable auquel notre pointeur pointe, on utilise un * devant notre pointeur
	UpdateA(mypointer, 20)
	/*sa va changer le 10 en 20vu que le mypointer contient l'addresse memoire de la variable number (qui est egal a 10) dnc en modifiant la valeur
	du *pointeur, on modifie la variable 10 et on le remplace par value */
	fmt.Printf("la valeur de l'address memoire %v a changé en %v", mypointer, number)

}
