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
	return NewXiaomi12()
}

func (x *Xiaomi) makeWatch() IWatch {
	return NewMiWatch()
}

type Huawei struct{}

func (h *Huawei) makePhone() IPhone {
	return NewHuaweiMeta60()
}

func (h *Huawei) makeWatch() IWatch {
	return NewHuaweiWatch()
}
