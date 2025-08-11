package main

import "fmt"

func main(){
	fmt.Println("This is pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer", ptr)
	fmt.Println("Value of acutal pointer", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber) 
}