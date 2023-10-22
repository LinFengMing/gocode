package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 按字元統計字串的長度
	str := "hello"
	fmt.Println("str len =", len(str))
	// 字符串遍歷，同時處理有中文的問題
	str2 := "hello北京"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符 = %c\n", r[i])
	}
	// 字串轉整數
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("轉換錯誤:", err)
	} else {
		fmt.Println("轉換的結果:", n)
	}
	// 整數轉字串
	str3 := strconv.Itoa(12345)
	fmt.Printf("str3 = %v, str3 type = %T\n", str3, str3)
	// 字串轉 byte
	var bytes = []byte("hello go")
	fmt.Printf("bytes = %v\n", bytes)
	// byte 轉字串
	srt4 := string([]byte{97, 98, 99})
	fmt.Printf("srt4 = %v\n", srt4)
	// 10 進位轉 2, 8, 16 進位
	str5 := strconv.FormatInt(123, 2)
	fmt.Printf("字串對應的進位是 %v\n", str5)
	// 字串搜尋
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b = %v\n", b)
	// 統計字串有幾個指定的字串
	num := strings.Count("ceheese", "e")
	fmt.Printf("num = %v\n", num)
	// 不區分大小寫的字串比較
	b2 := strings.EqualFold("abc", "Abc")
	fmt.Printf("b2 = %v\n", b2)
	// 尋找字串第一次出現的位置，沒有回傳 -1
	index := strings.Index("NLT_abc", "abc")
	fmt.Printf("index = %v\n", index)
}
