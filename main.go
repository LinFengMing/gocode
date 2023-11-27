package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	// r1 記憶體位址是連續的
	fmt.Printf("r1.leftUp.x 的位址 = %p, r1.leftUp.y 的位址 = %p, r1.rightDown.x 的位址 = %p, r1.rightDown.y 的位址 = %p\n",
		&r1.leftUp.x, &r1.leftUp.y, &r1.rightDown.x, &r1.rightDown.y)
	// r2 有兩個 *Point 類型，這兩個 *Point 的記憶體位址也是連續的，但指向的記憶體位址不一定是連續的
	r2 := Rect2{&Point{10, 20}, &Point{30, 40}}
	fmt.Printf("r2.leftUp 的位址 = %p, r2.rightDown 的位址 = %p\n", &r2.leftUp, &r2.rightDown)
	fmt.Printf("r2.leftUp 指向的位址 = %p, r2.rightDown 指向的位址 = %p\n", r2.leftUp, r2.rightDown)
}
