package main

import (
	"fmt"
)

func main() {
	cities := make(map[string]string)
	cities["no1"] = "tokyo"
	cities["no2"] = "osaka"
	cities["no3"] = "nagoya"
	for k, v := range cities {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}
	fmt.Println("cities 有", len(cities), "組 key-value")
	students := make(map[string]map[string]string)
	students["no1"] = make(map[string]string, 2)
	students["no1"]["name"] = "tom"
	students["no1"]["sex"] = "male"
	students["no2"] = make(map[string]string, 2)
	students["no2"]["name"] = "mary"
	students["no2"]["sex"] = "woman"
	for k1, v1 := range students {
		fmt.Printf("k1 = %v\n", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\tk2 = %v, v2 = %v\n", k2, v2)
		}
		fmt.Println()
	}
}
