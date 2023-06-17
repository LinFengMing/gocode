package main

import "fmt"

// 定義全域變數
var n8 = 100
var n9 = 200
var name4 = "jack"

// 一次性宣告
var (
	n10   = 300
	n11   = 900
	name5 = "mary"
)

func main() {
	// 第一種：指定變數類型，宣告後若不賦值，使用默認值
	// int的默認值是0
	var i int
	fmt.Println("i =", i)
	// 第二種：根據值自行判斷變數類型(類型推導)
	var num = 10.11
	fmt.Println("num =", num)
	// 第三種：省略 var ，注意 := 左側的變數不應該是已經宣告過的，否則會導致編譯錯誤
	// 等於 var name string  name = "tom"
	// := 的 : 不能省略，否則會錯誤
	name := "tom"
	fmt.Println("name =", name)
	// 第四種：一次性宣告多個變數
	var n1, n2, n3 int
	fmt.Println("n1 =", n1, "n2 =", n2, "n3 =", n3)
	var n4, name2, n5 = 100, "tom", 888
	fmt.Println("n4 =", n4, "name2 =", name2, "n5 =", n5)
	// 一次性宣告多個變數，一樣可以使用類型推導
	n6, name3, n7 := 100, "tom", 888
	fmt.Println("n6 =", n6, "name3 =", name3, "n7 =", n7)
	// 輸出全域變數
	fmt.Println("n8 =", n8, "name4 =", name4, "n9 =", n9)
	fmt.Println("n10 =", n10, "name5 =", name5, "n11 =", n11)
}
