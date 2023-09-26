package main

import "fmt"

func stringsAndNums() { // REF - https://go.dev/ref/spec#Numeric_types
	// strings
	var nameOne string = "Zane"
	var nameTwo = "2"
	var nameThree string
	nameFour := "4" // cannot be used outside of a function!!

	nameThree = "3"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// integers
	ageOne := 20
	ageTwo := 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	// var numOne int8 = 127 // max for int8
	// var numTwo int8 = -128 // min for int8
	// var numThree uint8 = 255 // cannot be negative, max for uint8

	// floats
	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 89438923328492387948239.9347294723923849324839175915
	// scoreThree := 1.5
}