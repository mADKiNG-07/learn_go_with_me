package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// variables
	// 1. explicitly typing
	var nameOne string = "Mario"
	fmt.Println(nameOne)

	// 2. implicitly typing
	var nameTwo = "Luigi"
	fmt.Println(nameTwo)

	// 3. defining without initializing
	var nameThree string
	nameThree = "Princess"
	fmt.Println(nameThree)

	// 4. short hand version
	//!! NOTE: CANT BE USED OUTSIDE OF A FUNCTION
	nameFour := "yoshi"
	fmt.Println(nameFour)

	// redefining variables
	nameOne = "bowser"
	nameThree = "peach"
	fmt.Println(nameOne, nameTwo, nameThree)

}
