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

	fmt.Print("investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Returned Rate: ")
	fmt.Scan(&expectedReturnedRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnedRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
