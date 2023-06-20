package main

import (
	"fmt"
)

func main() {
	str := "abc\nabc"
	fmt.Println("str =", str)
	str2 := `
	package main

	import (
		"fmt"
		"unsafe"
	)

	func main() {
		// bool 類型只能是 true or false
		var b = false
		fmt.Println("b =", b)
		// bool 類型占用儲存空間是一個字節
		fmt.Println("b 的佔用空間 =", unsafe.Sizeof(b))
	}
	`
	fmt.Println("str2 =", str2)
	// 字串拼接方式
	var str3 = "hello " + "world"
	str3 += " haha!"
	fmt.Println("str3 =", str3)
	var str4 = "hello " + "world" + "hello " + "world" +
		"hello " + "world" +
		"hello " + "world" +
		"hello " + "world"
	fmt.Println("str4 =", str4)
}
