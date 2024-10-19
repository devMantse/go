package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	revenue, err1 := getUserInput("Enter Revenue: ")
	expenses, err2 := getUserInput("Enter Expenses: ")
	taxRate, err3 := getUserInput("Enter Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculateFinnacials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.1f\n", ratio)
	storeRes(ebt, profit, ratio)

}

func storeRes(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit:%.1f\nRatio: %.3f\n", ebt, profit, ratio)

	os.WriteFile("results.text", []byte(results), 0644)
}

func calculateFinnacials(revenue, expenses, taxRate float64) (float64, float64, float64) {

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("positive value required")
	}

	return userInput, nil
}
