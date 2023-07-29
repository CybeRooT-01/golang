package main

import "fmt"

func main() {
	//disonc c'est l'quivalent des tableau associatif en php cle/valeur
	SuperMarketPrice := map[string]int{
		"prince":   8,
		"baguette": 16,
		"Bonbon":   2,
	}
	fmt.Println(SuperMarketPrice["prince"])
	for key, value := range SuperMarketPrice {
		fmt.Println(key, "vaut", value)
	}
	annee := 0
	Mois := map[string]int{
		"Janvier":   31,
		"Fevrier":   28,
		"Mars":      31,
		"Avril":     30,
		"Mai":       31,
		"Juin":      30,
		"Julliet":   31,
		"Aout":      31,
		"Septembre": 30,
		"Octobre":   31,
		"Novembre":  30,
		"Decembre":  31,
	}
	for _, value := range Mois { //le underscore pour dire on va pas utiliser le key
		annee += value //a chaque fois on ajoute la valeur courante de value dns annee
	}
	fmt.Printf("Le nombre de jours dans l'annee est %v", annee)
}
