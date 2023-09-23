package main

var (
	// PeopleSubject 实现 Subject 接口
	_ Subject = (*PeopleSubject)(nil)
)

// 观察者模式
func main() {
	s := new(PeopleSubject)
	s.RegisterObserver("angle", new(AngelObserver))
	s.RegisterObserver("demon", new(DemonObserver))
	s.SetQuestion("what should i do?")
}
