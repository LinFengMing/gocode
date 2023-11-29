package main

import (
	"encoding/json"
	"fmt"
)

type Master struct {
	Name  string `json:"name"` // struct tag
	Agr   int    `json:"agr"`
	Skill string `json:"skill"`
}

func main() {
	master := Master{"牛魔王", 500, "芭蕉扇"}
	jsonStr, err := json.Marshal(master)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonStr))
}
