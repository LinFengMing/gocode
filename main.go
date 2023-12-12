package main

import (
	"fmt"
)

type Usb interface {
	Say()
}
type Stu struct {
}

func (stu *Stu) Say() {
	fmt.Println("Say()...")
}
func main() {
	var stu Stu
	var usb Usb = &stu
	usb.Say()
}
