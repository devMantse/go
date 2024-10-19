package main

import (
	"fmt"
	"os"
)

const afile = "balance.txt" // Fixed filename

func balanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(afile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getFromFile(afile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
		panic(err)
	}

	for {
		presentOptions()
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
			balanceToFile(accountBalance)
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
			balanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}
