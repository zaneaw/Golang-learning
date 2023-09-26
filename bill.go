package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println(opt)
}

// create bill with user input
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

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
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", billObj.tip)

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