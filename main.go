package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"score"`
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("白骨精")
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值為=%v\n", i, val.Elem().Field(i).Kind())
	}
	fmt.Printf("struct has %d Field\n", num)
	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag = %v\n", tag)
	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d method\n", numOfMethod)
	val.Elem().Method(0).Call(nil)
}
func main() {
	var a Monster = Monster{
		Name:  "牛魔王",
		Age:   500,
		Score: 98.8,
	}
	result, _ := json.Marshal(a)
	fmt.Println("json result =", string(result))
	TestStruct(&a)
	fmt.Println(a)
}
