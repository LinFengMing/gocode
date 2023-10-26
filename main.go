package main

import (
	"errors"
	"fmt"
)

func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("讀取文件錯誤")
	}
}

func test() {
	err := readConf("config.ini")
	if err != nil {
		panic(err)
	}
	fmt.Println("test")
}

func main() {
	test()
	fmt.Println("main")
}
