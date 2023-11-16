package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "tokyo"
	cities["no2"] = "osaka"
	cities["no3"] = "nagoya"
	// key 不存在，就是新增，存在就是修改
	cities["no3"] = "fukuoka"
	// delete 指定的 key 不存在時，删除不會執行，也不會報錯
	delete(cities, "no4")
	fmt.Println(cities)
	// 尋找 map 指定 key 是否存在
	val, ok := cities["no4"]
	if ok {
		fmt.Println("no4 key =", val)
	} else {
		fmt.Println("no4 key 不存在")
	}
	// 刪除 map 全部 key
	cities = make(map[string]string)
	fmt.Println(cities)
}
