package memento

import "testing"

func TestMemento(t *testing.T) {
	originator := new(Originator)
	caretaker := new(Caretaker)
	originator.State = "On"
	caretaker.Memento = originator.CreateMemento()
	originator.State = "Off"
	originator.SetMemento(caretaker.Memento)
	if originator.State != "On" {
		t.Errorf("Expect State to %s, but %s.", originator.State, "On")
	}
}
