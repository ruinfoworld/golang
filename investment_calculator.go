package main

import (
	"fmt"
	"math"
)
const inflationRate = 2.5
func main() {
	var investmentAmount float64
	var year float64
	expectedReturnRate := 5.5
	
	// fmt.Print("Investment Amount: ")
	getUserInput("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	getUserInput("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ") 
	getUserInput("Years: ")
	fmt.Scan(&year)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, year)

	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, year)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, year)

	formatedFV := fmt.Sprintf("Future value: %.2f\n",futureValue)
	formatedRFV := fmt.Sprintf("Future value (adjusted for Inflation): %.2f\n", futureRealValue)

	//%.2f will display float with 2 decimal points.
	// fmt.Printf("Future value: %.2f\nFuture value (adjusted for Inflation): %.2f",futureValue, futureRealValue)
    // fmt.Printf(`Future value: %.2f
	// Future value (adjusted for Inflation): %.2f`,futureValue, futureRealValue)
	// %v will display value as it is
	// fmt.Println("Future value: ",futureValue)
	// fmt.Println("Future value (adjusted for Inflation):",futureRealValue)

	fmt.Print(formatedFV, formatedRFV)
}

func getUserInput(text string){
	fmt.Print(text);
}

func calculateFutureValue(investmentAmount, expectedReturnRate, year float64) (fv float64, rfv float64){
	// fv := investmentAmount * math.Pow(1 + expectedReturnRate/100, year)
	// frv := fv / math.Pow(1 + inflationRate / 100, year)
	fv = investmentAmount * math.Pow(1 + expectedReturnRate/100, year)
	rfv = fv / math.Pow(1 + inflationRate / 100, year)
	return fv,rfv
	// return
}