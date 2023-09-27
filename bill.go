package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bill struct {
	name string
	items map[string]float64
	tip float64
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(billObj *bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	
	switch opt {
		case "a":
			fmt.Println("You chose Add Item!")
			name, _ := getInput("Item name: ", reader)
			price, _ := getInput("Item price: ", reader)

			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number.")
				promptOptions(billObj)
			}

			billObj.addItem(name, p)
			fmt.Printf("Item added -\n%-25v ...$%0.2f\n", name+":", p)

			promptOptions(billObj)
		case "s":
			fmt.Println("You chose Save!")
			fmt.Println(billObj.formatBill())
			billObj.saveBill()
		case "t":
			fmt.Println("You chose Tip!")

			tip, _ := getInput("Enter tip amount ($): ", reader)

			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The tip must be a number.")
				promptOptions(billObj)
			}

			billObj.updateTip(t)
			fmt.Printf("Tip added - $%0.2f\n", t)

			promptOptions(billObj)
		default:
			fmt.Println("That is not a valid option...")
			promptOptions(billObj)
	}
}

// create bill with user input
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created bill for -", b.name)

	return b
}

// make new bills
func newBill(name string) bill {
	newBill := bill{
		name: name,
		items: map[string]float64{"pie": 5.99, "cake": 3.99},
		tip: 0,
	}

	return newBill
}

// format the bill
func (billObj *bill) formatBill() string {
	fs := "Bill Breakdown:\n"
	var total float64 = 0

	// list items
	for key, val := range billObj.items {
		fs += fmt.Sprintf("%-25v ...$%0.2f\n", key+":", val)
		total += val
	}

	// if tip is less than 0, set it to 0
	if (billObj.tip < 0) {
		billObj.tip = 0
	}

	// add tip line to output
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "tip:", billObj.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+billObj.tip)

	return fs
}

// update tip
func (billObj *bill) updateTip(tipAmount float64) {
	billObj.tip = tipAmount
}

// add item to bill
func (billObj *bill) addItem(name string, price float64) {
	billObj.items[name] = price
}

// save bill to file
func (billObj *bill) saveBill() {
	data := []byte(billObj.formatBill())

	err := os.WriteFile("bills/"+billObj.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file!")
}