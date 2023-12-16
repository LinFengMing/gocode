package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	name string
}

func (p Phone) Start() {
	fmt.Println("手機開始工作")
}
func (p Phone) Stop() {
	fmt.Println("手機停止工作")
}

type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相機開始工作")
}
func (c Camera) Stop() {
	fmt.Println("相機停止工作")
}
func main() {
	// 定義多態數組
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米手機"}
	usbArr[1] = Phone{"華為手機"}
	usbArr[2] = Camera{"索尼相機"}
	fmt.Println(usbArr)
}
