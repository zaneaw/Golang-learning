package main

import "fmt"

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