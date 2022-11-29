package main

import "fmt"

type single struct{}

var normalSingleInstance *single

func getNormalSingle() *single {
	if normalSingleInstance == nil {
		fmt.Println("creating single instance")
		normalSingleInstance = new(single)
	} else {
		fmt.Println("single instance already created")
	}
	return normalSingleInstance
}

func normal() {
	for i := 0; i < 10; i++ {
		getNormalSingle()
	}
}
