package main

import "fmt"

// 抽象工厂方法模式
func main() {
	xiaomiFactory := GetFactory("xiaomi")
	xiaomiPhone := xiaomiFactory.makePhone()
	xiaomiWatch := xiaomiFactory.makeWatch()

	huaweiFactory := GetFactory("huawei")
	huaweiPhone := huaweiFactory.makePhone()
	huaweiWatch := huaweiFactory.makeWatch()

	printPhoneDetails(xiaomiPhone)
	printWatchDetails(xiaomiWatch)
	printPhoneDetails(huaweiPhone)
	printWatchDetails(huaweiWatch)
}

func printPhoneDetails(p IPhone) {
	fmt.Printf("Phone: %s", p.getPhoneName())
	fmt.Println()
	fmt.Printf("Price: %.2f", p.getPhonePrice())
	fmt.Println()
	fmt.Println("---------------------------")
}

func printWatchDetails(w IWatch) {
	fmt.Printf("watch: %s", w.getWatchName())
	fmt.Println()
	fmt.Printf("Price: %.2f", w.getWatchPrice())
	fmt.Println()
	fmt.Println("---------------------------")
}
