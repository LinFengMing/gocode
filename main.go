package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file)
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
