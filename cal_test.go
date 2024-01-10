package main

import "testing"

func TestAddUpper(t *testing.T) {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper 錯誤返回值 = %v 期望值 = %v\n", res, 55)
	}
	t.Logf("AddUpper 正确返回值 = %v 期望值 = %v\n", res, 55)
}
