package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is Slices")

	var fruitlist = []string{"Apple", "Cherry", "Peach"}
	fmt.Println("Type of fruitlist is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3])
	fmt.Println(fruitlist)	

	highscores := make([]int, 4)

	highscores[0] = 234
	highscores[1] = 945
	highscores[2] = 465
	highscores[3] = 867
	// highscores[4] = 777

	highscores = append(highscores, 666, 555, 321)
	
	fmt.Println(highscores)

	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println(highscores)

}