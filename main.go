package main

import (
	"fmt"
)

func main() {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("請輸入第 %d 個班第 %d 個學生的成績：", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}
	total := 0.0
	sduentNum := 0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sduentNum++
			sum += scores[i][j]
		}
		total += sum
		fmt.Printf("第 %d 個班的總分為 %v, 平均分為 %v\n", i+1, sum,
			sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班的總分為 %v, 平均分為 %v\n", total, total/float64(sduentNum))
}
