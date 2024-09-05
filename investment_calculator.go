package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var year float64
	expectedReturnRate := 5.5
	
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ") 
	fmt.Scan(&year)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, year)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, year)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}