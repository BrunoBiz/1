package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount, years, expectedReturnRate float64

	investmentAmount = scanPrintExercise("Investment Amount:")
	expectedReturnRate = scanPrintExercise("Expected return rate:")
	years = scanPrintExercise("Years:")

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future value (adjusted): %.1f\n", futureRealValue)

	//fmt.Println(futureValue)
	//fmt.Println(futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv := fv / math.Pow(1+inflationRate/100, years)

	return fv, rfv
}

func scanPrintExercise(text string) float64 {
	var returnValue float64

	outputText(text)
	fmt.Scan(&returnValue)

	return returnValue
}
