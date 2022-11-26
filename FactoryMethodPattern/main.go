package main

import "fmt"

func main() {
	xiaomi, err := getPhone("xiaomi")
	if err != nil {
		panic(err)
	}
	huawei, err := getPhone("huawei")
	if err != nil {
		panic(err)
	}

	printDetails(xiaomi)
	printDetails(huawei)
}

func printDetails(p IPhone) {
	fmt.Printf("Phone: %s", p.getName())
	fmt.Println()
	fmt.Printf("Price: %.2f", p.getPrice())
	fmt.Println()
	fmt.Println("---------------------------")
}
