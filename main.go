package main

import (
	"fmt"
	"sort"
)

func main() {
	map1 := make(map[int]int)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	fmt.Println(map1)
	// map 的 key 放入 slice
	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}
	// slice 排序
	sort.Ints(keys)
	fmt.Println(keys)
	for _, k := range keys {
		fmt.Printf("map1[%v] = %v\n", k, map1[k])
	}
}
