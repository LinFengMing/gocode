package main

import (
	"fmt"
)

func main() {
	names := [4]string{"白眉鷹王", "金毛獅王", "紫衫龍王", "青翼蝠王"}
	var heroName = ""
	fmt.Println("請輸入要查詢的人名：")
	fmt.Scanln(&heroName)
	for i := range names {
		if heroName == names[i] {
			fmt.Printf("找到了，這個人是 %v，編號 %v\n", heroName, i)
			break
		} else if i == len(names)-1 {
			fmt.Println("沒找到")
		}
	}
	index := -1
	for i := range names {
		if heroName == names[i] {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Printf("找到了，這個人是 %v，編號 %v\n", heroName, index)
	} else {
		fmt.Println("沒找到")
	}
}
