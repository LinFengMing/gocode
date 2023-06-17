package main

import "fmt"

func main() {
	fmt.Println("tom\tjack")
	fmt.Println("hello\nworld")
	fmt.Println("C:\\Users\\unuse\\Desktop\\Golang\\gocode")
	fmt.Println("tom說\"i love you\"")
	// \r，以當前行的最前面開始輸出，覆蓋掉以前內容
	fmt.Println("天龍八赴雪山飛\r張飛")
	// 練習
	fmt.Println("姓名\t年齡\t籍貫\t住址\njohn\t12\t河北\t北京")
}
