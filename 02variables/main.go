package main

import "fmt"

const LoginToken string = "typewriter" // Public

func main() {
	var username string = "eclipse"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.23534566766
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some alias
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is type of: %T \n", anotherVariable)

	// implicit type
	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style
	numberofuser := 300000.0
	fmt.Println(numberofuser)

	fmt.Println(numberofuser)
	fmt.Printf("Variable is of type: %T \n", numberofuser)
}
