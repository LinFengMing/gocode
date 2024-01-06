package main

import (
	"fmt"
	"os"
)

func main() {
	// read file
	data, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v", string(data))
}
