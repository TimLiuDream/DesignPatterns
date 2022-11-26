package main

// 抽象工厂，包含一系列创建产品的抽象
type IFactory interface {
	makePhone() IPhone
	makeWatch() IWatch
}

func GetFactory(tp string) IFactory {
	if tp == "xiaomi" {
		return &Xiaomi{}
	} else if tp == "huawei" {
		return &Huawei{}
	} else {
		panic("not found type")
	}
	return nil
}

type Xiaomi struct{}

func (x *Xiaomi) makePhone() IPhone {
	return &Phone{
		name:  "xiaomi 12",
		price: 3999,
	}
}

func (x *Xiaomi) makeWatch() IWatch {
	return &Watch{
		name:  "miWatch",
		price: 299,
	}
}

type Huawei struct{}

func (h *Huawei) makePhone() IPhone {
	return &Phone{
		name:  "iPhone 14",
		price: 6999,
	}
}

func (h *Huawei) makeWatch() IWatch {
	return &Watch{
		name:  "apple watch",
		price: 2999,
	}
}
