package main

import (
	"fmt"
)

type Student struct {
}

func TypeJudge(items ...interface{}) {
	for index, x := range items {
		index++
		switch x.(type) {
		case bool:
			fmt.Printf("第 %v 個參數是 bool 型別，值是 %v\n", index, x)
		case float32:
			fmt.Printf("第 %v 個參數是 float32 型別，值是 %v\n", index, x)
		case float64:
			fmt.Printf("第 %v 個參數是 float64 型別，值是 %v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第 %v 個參數是 整數 型別，值是 %v\n", index, x)
		case string:
			fmt.Printf("第 %v 個參數是 字串 型別，值是 %v\n", index, x)
		case Student:
			fmt.Printf("第 %v 個參數是 Student 型別，值是 %v\n", index, x)
		case *Student:
			fmt.Printf("第 %v 個參數是 *Student 型別，值是 %v\n", index, x)
		default:
			fmt.Printf("第 %v 個參數是 不確定 型別，值是 %v\n", index, x)
		}
	}
}
func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int32 = 30
	var name string = "tom"
	address := "北京"
	n4 := 300
	stu1 := Student{}
	stu2 := &Student{}
	TypeJudge(n1, n2, n3, name, address, n4, stu1, stu2)
}
