package core

import (
	"fmt"
	"math"
)
//NorthWest method transport
func NorthWest(fromProvider []float64, toConsumer []float64) [][]float64 {
	var result [][]float64
	for fromIndex, fromVal := range fromProvider {
		var locResult []float64
		for toIndex, toVal := range toConsumer {
			r := math.Min(fromVal, toVal)
			fmt.Printf("[%d(%.f)][%d(%.f)] = %.2f\t", fromIndex, fromVal, toIndex, toVal, r)
			fromVal = fromVal - r
			fromProvider[fromIndex] = fromVal
			toConsumer[toIndex] = toVal - r
			locResult = append(locResult, r)
			
			

		}
		result = append(result, locResult)
		fmt.Println("\n-----------Result--------")
		fmt.Printf("\nResult: %v\n", locResult)
	}
	return result
}
