package main

import "fmt"

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