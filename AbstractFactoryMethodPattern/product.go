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

type Xiaomi12 struct {
	Phone
}

func NewXiaomi12() IPhone {
	return &Phone{
		name:  "xiaomi 12",
		price: 3999,
	}
}

type MiWatch struct {
	Watch
}

func NewMiWatch() IWatch {
	return &Watch{
		name:  "miWatch",
		price: 299,
	}
}

type HuaweiMeta60 struct {
	Phone
}

func NewHuaweiMeta60() IPhone {
	return &Phone{
		name:  "Meta 60",
		price: 6999,
	}
}

type HuaweiWatch struct {
	Watch
}

func NewHuaweiWatch() IWatch {
	return &Watch{
		name:  "huawei watch",
		price: 2999,
	}
}
