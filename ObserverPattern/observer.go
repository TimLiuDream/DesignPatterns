package main

import "fmt"

type Observer interface {
	Answer()
}

type AngelObserver struct{}

func (o *AngelObserver) Answer() {
	fmt.Println("fuck off")
}

type DemonObserver struct{}

func (o *DemonObserver) Answer() {
	fmt.Println("go for it!")
}
