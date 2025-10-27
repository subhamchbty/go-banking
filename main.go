package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	taxRate := 12.0

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	formattedRatio := fmt.Sprintf("Earning ratio: %.3f\n", ratio)

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Print(formattedRatio)
}
