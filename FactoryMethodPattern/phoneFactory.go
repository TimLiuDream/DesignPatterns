package main

import "fmt"

func getPhone(phoneType string) (IPhone, error) {
	if phoneType == "xiaomi" {
		return NewXiaomi(), nil
	}
	if phoneType == "huawei" {
		return NewHuawei(), nil
	}
	return nil, fmt.Errorf("wrong phone type passed")
}
