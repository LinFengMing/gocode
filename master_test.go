package main

import "testing"

func TestStore(t *testing.T) {
	m := &Monster{
		Name:  "紅孩兒",
		Age:   10,
		Skill: "噴火",
	}
	res := m.Store()
	if !res {
		t.Fatalf("測試失敗，希望為 = %v，實際為 = %v", true, res)
	}
	t.Logf("測試成功")
}
func TestReStore(t *testing.T) {
	var m = &Monster{}
	res := m.ReStore()
	if !res {
		t.Fatalf("測試失敗，希望為 = %v，實際為 = %v", true, res)
	}
	if m.Name != "紅孩兒" {
		t.Fatalf("測試失敗，希望為 = %v，實際為 = %v", "紅孩兒", m.Name)
	}
	t.Logf("測試成功")
}
