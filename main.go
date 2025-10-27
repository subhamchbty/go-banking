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

	ebt := getEbt(revenue, expenses)
	profit, ratio := getProfitRatio(ebt, taxRate)

	formattedRatio := fmt.Sprintf("Earning ratio: %.3f\n", ratio)

	fmt.Println("Earnings before tax:", ebt)
	fmt.Println("Profit:", profit)
	fmt.Print(formattedRatio)
}

func getEbt(revenue, expenses float64) float64 {
	return revenue - expenses
}

func getProfitRatio(ebt, taxRate float64) (float64, float64) {
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return profit, ratio
}
