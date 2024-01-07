package main

import (
	"fmt"
	"os"
)

func main() {
	filepath := "abc.txt"
	filepath2 := "kkk.txt"
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = os.WriteFile(filepath2, data, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
}
