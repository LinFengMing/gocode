package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum =", cpuNum)
	// go1.8 前設定 go 程式執行的最大的 cpu 核數
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
