package main

import "fmt"

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