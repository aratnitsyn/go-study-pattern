package observer

type Publisher interface {
	Attach(observer Observer)
	SetState(state string)
	Notify()
}

type Observer interface {
	Update(state string)
}

type ConcretePublisher struct {
	observers []Observer
	state     string
}

func (s *ConcretePublisher) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcretePublisher) SetState(state string) {
	s.state = state
}

func (s *ConcretePublisher) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.state)
	}
}

type ConcreteObserver struct {
	state string
}

func (s *ConcreteObserver) Update(state string) {
	s.state = state
}
