package main

import "fmt"

func main() {
	fmt.Println("This is array in Go")

	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[2] = "cherry"
	fruitlist[3] = "kiwi"

	fmt.Println("Fruit list is", fruitlist)
	fmt.Println("Fruit list is", len(fruitlist))

	var vegList = [5]string{"potato", "beans", "carrot"}
	fmt.Println("list is", vegList)
	fmt.Println("list is", len(vegList))
}