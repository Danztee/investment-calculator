package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5

	var investmentAmount float64
	var years float64
	expectedReturnedRate := 5.5

	// fmt.Print("Investment amount: ")
	outputText("Investment amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Returned Rate: ")
	outputText("Expected Returned Rate: ")
	fmt.Scan(&expectedReturnedRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnedRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

	// output information
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)

	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}
