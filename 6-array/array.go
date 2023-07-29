package main

import "fmt"

func main() {
	//premier methode: declare puis on assigne
	var List [2]int
	List[0] = 10
	List[1] = 20
	fmt.Println(List)
	fmt.Println(List[0])
	fmt.Println(List[1])

	//deuxieme methode: declare et assigne en meme temps
	Newlist := [...]int{40, 50, 60}
	fmt.Println(Newlist)
	fmt.Println(Newlist[0])
	fmt.Println(Newlist[1])
	fmt.Println(Newlist[2])

	//si on a une enorme array faut pas faire des println. il faut boucler
	Biglist := [...]int{70, 80, 90, 100, 7777, 6547, 545787, 32145}
	/*for i := 0; i <= 7; i++ {
		var j int = Biglist[i]
		fmt.Printf("Position %d vaut %d.\n", i, j)
	}
	soit cette methode soit tu use range*/
	for pos, value := range Biglist {
		fmt.Printf("Position %d vaut %d.\n", pos, value)
	}
}
