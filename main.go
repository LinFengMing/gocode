package main

import (
	"fmt"
)

func main() {
	hen1 := 3.0
	hen2 := 5.0
	hen3 := 1.0
	hen4 := 3.4
	hen5 := 2.0
	hen6 := 50.0
	totalWeight := hen1 + hen2 + hen3 + hen4 + hen5 + hen6
	avgWeight := fmt.Sprintf("%.2f", totalWeight/6)
	fmt.Printf("totalWeight = %v, avgWeight = %v", totalWeight, avgWeight)
}
