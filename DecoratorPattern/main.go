package main

import "fmt"

func main() {
	var b Beverage = new(Espresso)
	b = &Mocha{b}
	fmt.Println(b.Desc())
	fmt.Println(b.Cost())
}
