package main

import "fmt"

// 装饰者模式
func main() {
	var b Beverage = new(Espresso)
	b = &Mocha{b}
	fmt.Println(b.Desc())
	fmt.Println(b.Cost())
}
