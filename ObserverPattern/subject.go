package main

type Subject interface {
	RegisterObserver(name string, o Observer)
	RemoveObserve(name string)
	Notify()
}

type PeopleSubject struct {
	Observers map[string]Observer
	Question  string
}

func (s *PeopleSubject) RegisterObserver(name string, o Observer) {
	if s.Observers == nil {
		s.Observers = make(map[string]Observer)
	}
	s.Observers[name] = o
}

func (s *PeopleSubject) RemoveObserve(name string) {
	delete(s.Observers, name)
}

func (s *PeopleSubject) Notify() {
	for _, o := range s.Observers {
		o.Answer()
	}
}

func (s *PeopleSubject) SetQuestion(question string) {
	s.Question = question
	s.Notify()
}
