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
func (p Phone) Call() {
	fmt.Println("手機打電話")
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

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	// 如果 usb 指向的是 Phone 結構體變量，則調用 Call 方法
	if phone, flag := usb.(Phone); flag {
		phone.Call()
	}
	usb.Stop()
}
func main() {
	// 定義多態數組
	var usbArr [3]Usb
	usbArr[0] = Phone{"小米手機"}
	usbArr[1] = Phone{"華為手機"}
	usbArr[2] = Camera{"索尼相機"}
	var computer Computer
	for _, value := range usbArr {
		computer.Working(value)
		fmt.Println()
	}
}
