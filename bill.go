package main

import "fmt"

type bill struct {
	name string
	items map[string]float64
	tip float64
}

// make new bills
func createNewBill(name string) bill {
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