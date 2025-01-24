package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := (revenue - expenses) * ((100 - taxRate) / 100)
	ratio := earningBeforeTax / earningAfterTax

	fmt.Println(earningBeforeTax)
	fmt.Println(earningAfterTax)
	fmt.Println(ratio)
}
