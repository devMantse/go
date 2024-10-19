package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate = 6.5
	var investmentAmount float64
	var expectedReturnRate = 5.5
	var years float64

	futureValue, futureRealValue := doCalculation(investmentAmount, expectedReturnRate, inflationRate, years)

	fmt.Print("Entee Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	fmt.Println("Future Value: ", futureValue)
	fmt.Printf("Real Value: %v\f", futureRealValue)
	fmt.Printf("Future value: %.1f\nAdjusted: %.1f", futureValue, futureRealValue)
}

func doCalculation(investmentAmount, expectedReturnRate, inflationRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := investmentAmount * math.Pow(1+inflationRate/100, years)
	return fv, rfv

}
