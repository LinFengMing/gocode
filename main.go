package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 10
	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan
	fmt.Printf("num1 = %v, num2 = %v, num3 = %v\n", num1, num2, num3)
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)
	m1 := make(map[string]string, 20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"
	m2 := make(map[string]string, 20)
	m2["hero1"] = "宋江"
	m2["hero2"] = "武松"
	mapChan <- m1
	mapChan <- m2
	m3 := <-mapChan
	m4 := <-mapChan
	fmt.Printf("m3 = %v, m4 = %v\n", m3, m4)
	var catChan chan Cat
	catChan = make(chan Cat, 10)
	cat1 := Cat{"tom", 18}
	cat2 := Cat{"jack", 19}
	catChan <- cat1
	catChan <- cat2
	cat3 := <-catChan
	cat4 := <-catChan
	fmt.Printf("cat3 = %v, cat4 = %v\n", cat3, cat4)
	var cathan2 chan *Cat
	cathan2 = make(chan *Cat, 10)
	cat5 := Cat{"tom", 18}
	cat6 := Cat{"jack", 19}
	cathan2 <- &cat5
	cathan2 <- &cat6
	cat7 := <-cathan2
	cat8 := <-cathan2
	fmt.Printf("cat7 = %v, cat8 = %v\n", cat7, cat8)
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)
	cat9 := Cat{"tom", 18}
	cat10 := Cat{"jack", 19}
	allChan <- cat9
	allChan <- cat10
	allChan <- 10
	allChan <- "tom jack"
	cat11 := <-allChan
	cat12 := <-allChan
	v1 := <-allChan
	v2 := <-allChan
	fmt.Printf("cat11 = %v, cat12 = %v, v1 = %v, v2 = %v\n", cat11, cat12, v1, v2)
	var allChan2 chan interface{}
	allChan2 = make(chan interface{}, 3)
	allChan2 <- 10
	allChan2 <- "tom jack"
	cat13 := Cat{"小花貓", 4}
	allChan2 <- cat13
	<-allChan2
	<-allChan2
	cat14 := <-allChan2
	fmt.Printf("cat14 = %T, cat14 = %v\n", cat14, cat14)
	newCat := cat14.(Cat)
	fmt.Printf("newCat.Name = %v\n", newCat.Name)
}
