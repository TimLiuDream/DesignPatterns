package main

// 抽象出来的 phone 基类，子类（xiaomi、huawei）使用组合的方式可替代基类的行为
type Phone struct {
	name  string
	price float64
}

func (g *Phone) setName(name string) {
	g.name = name
}

func (g *Phone) getName() string {
	return g.name
}

func (g *Phone) setPrice(price float64) {
	g.price = price
}

func (g *Phone) getPrice() float64 {
	return g.price
}

type Xiaomi struct {
	Phone
}

func NewXiaomi() IPhone {
	return &Phone{
		name:  "xiaomi",
		price: 1999,
	}
}

type Huawei struct {
	Phone
}

func NewHuawei() IPhone {
	return &Phone{
		name:  "huawei",
		price: 4999,
	}
}
