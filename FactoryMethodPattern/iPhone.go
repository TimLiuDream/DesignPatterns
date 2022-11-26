package main

type IPhone interface {
	setName(name string)
	setPrice(price float64)
	getName() string
	getPrice() float64
}
