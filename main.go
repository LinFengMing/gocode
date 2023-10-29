package main

import (
	"fmt"
)

func main() {
	var heroes [3]string = [3]string{"卡特蓮娜", "雷尼克頓", "伊澤瑞爾"}
	for i, v := range heroes {
		fmt.Printf("i = %v, v = %v\n", i, v)
	}
	for _, v := range heroes {
		fmt.Printf("v = %v\n", v)
	}
}
