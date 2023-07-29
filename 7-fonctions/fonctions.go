package main

import (
	"errors"
	"fmt"
)

// simple fonction
func sayHello(name string) {
	fmt.Printf("Bonjour je suis %v \n", name)
}

// fonction calcule perimetre rectangle
func CalculatePerimRect(lo, la float64) float64 {
	return 2 * (lo + la)
}

// fonction avec erreur sis le user a oublier les paramettres
func Saybye(name string) (string, error) {
	if name == "" {
		return "", errors.New("\bLa fonction peux pas etre vide  veuillez specifier un nom")
	}
	byMessage := fmt.Sprintf("%v s'en va ! Bonne Soirée.", name)
	return byMessage, nil
}
func main() {
	//on appelle nos fonction
	sayHello("Birahim")
	perim := CalculatePerimRect(10.5, 5.4)
	fmt.Printf("Le perimetre vaut: %v \n", perim)
	fmt.Println(Saybye(""))
}
