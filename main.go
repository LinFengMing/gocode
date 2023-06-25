package main

import (
	"fmt"
	// 為了能使用 utils.go 文件的變數或函式，需要先引入該 model 包
	"gocode/model"
)

func main() {
	// 使用 utils.go 的 HeroName package.標識符
	fmt.Println(model.HeroName)
}
