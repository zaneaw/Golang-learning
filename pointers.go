package main

import "fmt"

func pointers() {
	name := "Zane"

	// updateName(name)

	// fmt.Println("Memory Address of name is:", &name)

	m := &name
	// fmt.Println("Memory Address of m is:", m)
	// fmt.Println("Value at memory address m is:", *m)

	fmt.Println("Name is:", name)
	updateName(m)
	fmt.Println("Name is now:", name)
}

func updateName(name *string) {
	*name = "Name Changed"
}