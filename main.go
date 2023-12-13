package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}
type HeroSlice []Hero

// 實作 Interface 介面
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool {
	// 按年齡從小到大排序
	return hs[i].Age < hs[j].Age
	// 按名稱從大到小排序
	// return hs[i].Name > hs[j].Name
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}
	for _, v := range heroes {
		fmt.Println(v)
	}
	// 使用 sort.Sort 排序
	sort.Sort(heroes)
	fmt.Println("排序後")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
