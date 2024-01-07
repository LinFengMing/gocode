package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	defer srcFile.Close()
	reader := bufio.NewReader(srcFile)
	desFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}
	writer := bufio.NewWriter(desFile)
	defer desFile.Close()
	return io.Copy(writer, reader)
}
func main() {
	srcFile := "gril.jpg"
	dstFile := "gril2.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("拷貝完成")
	} else {
		fmt.Printf("拷貝錯誤 err = %v\n", err)
	}
}
