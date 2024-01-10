package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Birthday string  `json:"birthday"`
	Sal      float64 `json:"sal"`
	Skill    string  `json:"skill"`
}

func testMap() string {
	// 定义一个map
	var a map[string]interface{}
	// 使用map，需要make
	a = make(map[string]interface{})
	a["name"] = "红孩兒"
	a["age"] = 30
	a["address"] = "洪崖洞"
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("marshal err = %v\n", err)
	}
	return string(data)
}
func unmarshalStruct() {
	str := `{"name":"牛魔王","age":500,"birthday":"2011-11-11","sal":8000,"skill":"牛魔拳"}`
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("monster = %v, monster.Name = %v\n", monster, monster.Name)
}
func unmarshalMap() {
	// str := `{"address":"洪崖洞","age":30,"name":"红孩兒"}`
	str := testMap()
	var a map[string]interface{}
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("map = %v\n", a)
}
func unmarshalSlice() {
	str := `[{"address":"北京","age":"7","name":"jack"},{"address":["墨西哥","夏威夷"],"age":"20","name":"tom"}]`
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}
	fmt.Printf("slice = %v\n", slice)
}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
