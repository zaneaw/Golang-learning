package main

import (
	"fmt"
	"sort"
	// "strings"
)

func main() {
	booleansAndConditionals()
	// loops()
	// standardLib()
	// arraysAndSlices()
	// stringsAndNums()
	// fmtStuff()
}

func booleansAndConditionals() { // REF -
	age := 35

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is greater than or equal to 40")
	}

	names := []string{"Zane", "Lilly", "Theodore", "Penelope", "Koopa"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at position", index)
			continue // skips the rest of the code in the loop and goes to the next iteration
		}
		if index > 2 {
			fmt.Println("Breaking at position", index)
			break // breaks out of the loop
		}

		fmt.Printf("The name at index %v is %v\n", index, value)
	}
}

func loops() { // REF -
	// x := 0

	// for x < 5 {
	// 	fmt.Println("Value of x is: ", x);
	// 	x++
	// }

	// for i := 0; i <  5; i++ {
	// 	fmt.Println("Value of i is: ", i);
	// }

	names := []string{"Zane", "Lilly", "Theodore", "Penelope"}

	// for index, value := range names {
	// 	fmt.Printf("The name at index %v is %v\n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The name is %v\n", value)
	}

	fmt.Println(names)

}

func standardLib() { // REF - https://pkg.go.dev/std
	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello!")) // returns true or false
	// fmt.Println((strings.ReplaceAll(greeting, "hello", "hi"))) // replace all instances of "hello" with "hi"
	// fmt.Println(strings.ToUpper(greeting)) // convert to uppercase
	// fmt.Println(strings.Index(greeting, "ll")) // returns index of first instance of "ll"
	// fmt.Println(strings.Split(greeting, " ")) // splits string into array of strings based on the separator

	// fmt.Println("Original string value: ", greeting)

	ages := []int{45, 20, 35, 30, 75, 70, 50, 25}

	sort.Ints(ages) // sorts the array of ints and alters the original array
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) // returns the index of the value 30
	fmt.Println(index)

	indexUnreachable := sort.SearchInts(ages, 90) // returns the index of the value 90, but since it is not in the array, it returns 1 more than the slice length
	fmt.Println(indexUnreachable)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names) // sorts the array of strings and alters the original array
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) // returns the index of the value "bowser"

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