package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 覆蓋寫入
	file, err := os.OpenFile("abc.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	str := "您好, 林先生!\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	// 追加寫入
	file, err = os.OpenFile("abc.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	str = "ABC, ENGLISG!\r\n"
	writer = bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
	// 讀取後輸出，追加寫入
	file, err = os.OpenFile("abc.txt", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	str = "hello, 台北!\r\n"
	writer = bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
