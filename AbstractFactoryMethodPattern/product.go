package main

// 抽象产品
type IPhone interface {
	setPhoneName(name string)
	setPhonePrice(price float64)
	getPhoneName() string
	getPhonePrice() float64
}

type IWatch interface {
	setWatchName(name string)
	setWatchPrice(price float64)
	getWatchName() string
	getWatchPrice() float64
}

type Phone struct {
	name  string
	price float64
}

func (g *Phone) setPhoneName(name string) {
	g.name = name
}

func (g *Phone) getPhoneName() string {
	return g.name
}

func (g *Phone) setPhonePrice(price float64) {
	g.price = price
}

func (g *Phone) getPhonePrice() float64 {
	return g.price
}

type Watch struct {
	name  string
	price float64
}

func (w *Watch) setWatchName(name string) {
	w.name = name
}

func (w *Watch) getWatchName() string {
	return w.name
}

func (w *Watch) setWatchPrice(price float64) {
	w.price = price
}

func (w *Watch) getWatchPrice() float64 {
	return w.price
}
