package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// variables
	// 1. explicitly typing
	// var nameOne string = "Mario"
	// fmt.Println(nameOne)

	// 2. implicitly typing
	// var nameTwo = "Luigi"
	// fmt.Println(nameTwo)

	// 3. defining without initializing
	// var nameThree string
	// nameThree = "Princess"
	// fmt.Println(nameThree)

	// 4. short hand version
	//!! NOTE: CANT BE USED OUTSIDE OF A FUNCTION
	// nameFour := "yoshi"
	// fmt.Println(nameFour)

	// redefining variables
	// nameOne = "bowser"
	// nameThree = "peach"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// 5. Intergers
	var ageOne int = 30 // explicitly typed
	var ageTwo = 25     // implicitly typed
	ageThree := 20      // short hand

	var ageFour int // defining without initializing
	ageFour = 35

	// 5a. Bits and Memory (we wouldnt need this most of the time, but its good to know)
	// int8, int16, int32, int64
	var smallNumber int8 = 27 // see go docs for refrerence (https://go.dev/ref/spec#Numeric_types)
	var numOne uint = 100     // unsigned int, no negative values
	fmt.Println(ageOne, ageTwo, ageThree, ageFour, smallNumber, numOne)

	// 5b. Floats
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 75384954.8 // default float type
	scoreThree := 100.5               // notice the inference here, it will be a float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// TRICK QUESTION: can we assign a float to an int?
	// var numFour int = scoreOne // NO, this will not work, we cannot assign a float to an int
	// it can be converted though, but there has to be a sacrifice :). just kidding, it will just round down
	var numFour int = int(scoreOne) // this will convert the float to an int, but it will round down
	fmt.Println(numFour)            // this will print -1, because it rounded down
}
