package main

import (
	"fmt"
	"sort"
)

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