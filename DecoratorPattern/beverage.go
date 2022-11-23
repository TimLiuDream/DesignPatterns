package main

import "fmt"

// Beverage 定义接口，让之后的装饰者以及被装饰者都实现此接口
type Beverage interface {
	Desc() string
	Cost() float64
}

type Espresso struct {
	Beverage
}

func (e *Espresso) Desc() string {
	return "Espresso!"
}

func (e *Espresso) Cost() float64 {
	return 7.00
}

type Mocha struct {
	Beverage
}

func (m *Mocha) Desc() string {
	return fmt.Sprintf("%s,%s", "Mocha", m.Beverage.Desc())
}

func (m *Mocha) Cost() float64 {
	return 1.99 + m.Beverage.Cost()
}
