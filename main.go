package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

func modifyUser(users map[string]map[string]string, name string) {
	// 判断users中是否有name
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string, 2)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "暱稱~" + name
	}
}

func main() {
	users := make(map[string]map[string]string, 10)
	users["smith"] = make(map[string]string, 2)
	users["smith"]["pwd"] = "999999"
	users["smith"]["nickname"] = "小花貓"
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println(users)
}
