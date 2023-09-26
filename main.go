package main

import (
	"fmt"
)

var score = 99.5

func main() {
	myBill := createBill()
	promptOptions(&myBill)
	fmt.Println(myBill.formatBill())

	// myBill := newBill("Zane's Bill")

	// myBill.addItem("Onion Soup", 4.50)
	// myBill.addItem("Veggie Burger", 8.99)
	// myBill.addItem("Coffee", 2.99)
	// myBill.addItem("Soda", 2.50)

	// myBill.updateTip(8)

	// fmt.Println(myBill.formatBill())

	// pointers()

	// maps()

	// sayHello("Zane")

	// for _, value := range points {
	// 	fmt.Println(value)
	// }

	// showScore()

	// fn1, sn1 := getInitials("zane Wilson")
	// fmt.Println(fn1, sn1)
	// fn2, sn2 := getInitials("Lilly")
	// fmt.Println(fn2, sn2)

	// sayGreeting("Zane")
	// sayGreeting("Lilly")
	// sayBye("Zane")
	// cycleNames([]string{"Zane", "Lilly", "Theodore", "Penelope"}, sayGreeting)
	// cycleNames([]string{"Zane", "Lilly", "Theodore", "Penelope"}, sayBye)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("Circle 1 is %0.3f and Circle 2 is %0.3f\n", a1, a2)

	// booleansAndConditionals()
	// loops()
	// standardLib()
	// arraysAndSlices()
	// stringsAndNums()
	// fmtStuff()
}