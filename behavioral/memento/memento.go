package memento

type Originator struct {
	State string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{state: o.State}
}

func (o *Originator) SetMemento(memento *Memento) {
	o.State = memento.GetState()
}

type Memento struct {
	state string
}

func (m *Memento) GetState() string {
	return m.state
}

type Caretaker struct {
	Memento *Memento
}
