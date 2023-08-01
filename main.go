package main

import (
	"fmt"
)

func main() {
	var classNum int = 3
	var stuNum int = 5
	var totalSum float64 = 0.0
	var passCount int = 0
	for j := 1; j <= classNum; j++ {
		var sum float64 = 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("輸入第%d班第%d個學生的成績\n", j, i)
			fmt.Scanln(&score)
			sum += score
			if score >= 60 {
				passCount++
			}
		}
		fmt.Printf("第%d個班級的平均分是%v\n", j, sum/float64(stuNum))
		totalSum += sum
	}
	fmt.Printf("各個班級的總成績是%v，所有班級平均分是%v\n", totalSum, totalSum/float64(classNum*stuNum))
	fmt.Printf("及格人數為%v\n", passCount)
}
