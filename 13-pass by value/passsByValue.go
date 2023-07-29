package main

import "fmt"

func UpdateA(number int) int {
	number = 5
	return number
}
func UpdateB(element map[string]int) {
	element["Bonbon"] = 4     //tout les donnés qui sont ici seront redirige vers l'endroit en memoire ou ya le map
	element["Brouettes"] = 50 //ListeDeCourses. pourquoi ListeDeCourses? prsk c'est lui qu'on donne comme parametre ala fction
}
func main() {
	number := 10
	//UpdateA(number)
	number = UpdateA(number)
	fmt.Println(number) /*on a 10 prsk notre ordi stock une copie du variable number avec la valeur 5. C'est deux endroits
	different dans la memoire... et pour avoir 5, on stock le retour de la fonction dans number*/

	//----------------------conccernant les maps et fonction----------------
	ListeDeCourses := map[string]int{
		"Poire":   10,
		"Beignet": 5,
		"parfums": 50,
	}
	UpdateB(ListeDeCourses)
	fmt.Println(ListeDeCourses)
}
