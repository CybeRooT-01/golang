package main

type contact struct {
	nom   string
	age   int
	phone map[string]string //on fais un map pour si la personna plusieurs num de tel il nous dis que si c'est un fix ou un portable ensuite il nous donne son numero
}

func newContact(nom string, age int) contact {
	c := contact{
		nom:   nom,
		age:   age,
		phone: map[string]string{},
	}

	return c
}