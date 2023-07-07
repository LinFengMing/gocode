package main

import (
	"fmt"
)

func main() {
	var name string
	var age byte
	var sal float32
	var isPass bool
	// 方式一
	fmt.Println("請輸入姓名")
	fmt.Scanln(&name)
	fmt.Println("請輸入年齡")
	fmt.Scanln(&age)
	fmt.Println("請輸入薪水")
	fmt.Scanln(&sal)
	fmt.Println("請輸入是否通過考試")
	fmt.Scanln(&isPass)
	fmt.Printf("姓名是 %v \n年齡是 %v \n薪水是 %v \n是否通過考試 %v \n", name, age, sal, isPass)
	// 方式二
	fmt.Println("請輸入你的姓名，年齡，薪水，是否通過考試，請用空格隔開")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名是 %v \n年齡是 %v \n薪水是 %v \n是否通過考試 %v \n", name, age, sal, isPass)
}
