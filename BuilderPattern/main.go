package main

import "fmt"

// 生成器模式
func main() {
	normalBuilder := getBuilder("normal")
	villaBuilder := getBuilder("villa")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(villaBuilder)
	villa := director.buildHouse()

	fmt.Printf("\nIgloo House Door Type: %s\n", villa.doorType)
	fmt.Printf("Igloo House Window Type: %s\n", villa.windowType)
	fmt.Printf("Igloo House Num Floor: %d\n", villa.floor)

}
