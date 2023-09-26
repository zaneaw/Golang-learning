package main

import "fmt"

func maps() {
	// keys are strings and values are float64's
	menu := map[string]float64{
		"soup": 4.99,
		"pie": 7.99,
		"salad": 6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		267584967: "Zane",
		984759373: "Lilly",
		845775485: "Theodore",
		458458315: "Penelope",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584967])

	phonebook[984759373] = "Lilly Tonne-Daims"
	fmt.Println(phonebook[984759373])

	fmt.Println(phonebook)
}