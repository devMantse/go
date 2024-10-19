package main

import (
	"fmt"

	fileops "example.com/bank/fileops"
	options "example.com/bank/options"
)

const afile = "balance.txt" // Fixed filename

func main() {
	accountBalance, err := fileops.GetFromFile(afile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic(err)
	}

	for {
		options.PresentOptions()
		var choice int
		fmt.Print("Select: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Balance is: ", accountBalance)
		case 2:
			fmt.Println("Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated: ", accountBalance)
			fileops.ValueToFile(accountBalance, afile)
		case 3:
			fmt.Println("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Must be greater than 0")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Cannot withdraw more than balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("New Balance", accountBalance)
			fileops.ValueToFile(accountBalance, afile)

		default:
			fmt.Println("Goodbye")
			return
		}
	}
}
