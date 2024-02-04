package main

import (
	"fmt"
	"math"
)

func main() {
	var investimentAmount = 1000
	var expectReturnRate = 5.5
	var years = 10

	var futureValue = float64(investimentAmount) * math.Pow(1+expectReturnRate/100, float64(years))
	fmt.Println(futureValue)
}
