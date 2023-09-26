package main

import "fmt"

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