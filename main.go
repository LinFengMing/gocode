package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "tokyo"
	cities["no2"] = "osaka"
	cities["no3"] = "nagoya"
	fmt.Println(cities)
	heroes := map[string]string{
		"no1": "superman",
		"no2": "batman",
		"no3": "spiderman",
	}
	fmt.Println(heroes)
	students := make(map[string]map[string]string)
	students["no1"] = make(map[string]string, 2)
	students["no1"]["name"] = "tom"
	students["no1"]["sex"] = "male"
	students["no2"] = make(map[string]string, 2)
	students["no2"]["name"] = "mary"
	students["no1"]["sex"] = "woman"
	fmt.Println(students)
}
