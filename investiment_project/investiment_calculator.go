package main

import (
	"fmt"
	"math"
)

func main() {
	investimentAmount, years := 1000.0, 10.0
	var expectReturnRate float64 = 5.5

	futureValue := investimentAmount * math.Pow(1+expectReturnRate/100, years)
	fmt.Println(futureValue)
}
