package main

import "fmt"

func main() {

	var messageOne string = "My first go program."
	var messageTwo string = " I am happy right now."
	var messageThree string 

	fmt.Println(messageOne, messageTwo, messageThree)

	messageOne = "Message One"
	messageTwo = "Message Two"
	messageThree = "Message Three"

	// fmt.Println(messageOne, messageTwo, messageThree)

	// asign a variable without the var keyword - applicable if the type will be the same as previous variable and also works only inside the function
	messageFour := "Message Four"
	
	fmt.Println(messageFour)
	
	// integers
	var age int = 25
	var ageTwo = 32
	ageThree := 33

	fmt.Println(age, ageTwo, "\nAnd soon I will have: ", ageThree)

	// bits and memory
	var numberOne int8 = 125
	var numberTwo int16 = 255
	var numberThree uint8 = 24

	fmt.Println(numberOne, numberTwo, numberThree)

	// floats
	var score float32 = 124.32
	var score2 float64 = 17366187246164143.784264663284
	
	score3 := 1.5

	fmt.Println(score, score2, score3)



}