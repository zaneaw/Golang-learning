package main

import "fmt"

func main() {
	arraysAndSlices()
	// stringsAndNums()
	// fmtStuff()
}

func arraysAndSlices() { // REF - https://go.dev/ref/spec#Array_types

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} // type inference

	names := [4]string{"Zane", "Lilly", "Theodore", "Penelope"}
	names[0] = "Poppa"
	names[1] = "Momma"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	scores := []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3] // does not include index 3
	rangeTwo := names[2:] // includes index 2 to the end
	rangeThree := names[:3] // includes index 0 to 2

	rangeOne = append(rangeOne, "Koopa")

	fmt.Printf("Names One and Two: %v\n", rangeOne)
	fmt.Printf("Names Two to the end: %v\n", rangeTwo)
	fmt.Printf("Names Zero to Two: %v\n", rangeThree)

}
	
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

func fmtStuff() { // REF - https://pkg.go.dev/fmt
	// vars
	age := 28
	name := "Zane"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World!\n")
	fmt.Print("New Line...\n")

	// Print lines
	fmt.Println("Hello, Zane...")
	fmt.Println("Goodbye, Zane...")
	// fmt.Println("my age is", age, "and my name is", name)

	// Printf - Formatted strings
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q - quotes (does not work on ints)
	fmt.Printf("age is of type %T \n", age) // %T - type
	fmt.Printf("you scored %f points! \n", 225.55) // %f - float
	fmt.Printf("you scored %0.1f points! \n", 225.55) // %f - float

	// Sprintf - save formatted strings
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)

}