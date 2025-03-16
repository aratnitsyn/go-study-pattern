package chain_of_responsibility

import "testing"

func TestChainOfResponsibility(t *testing.T) {
	expect := "I'm handler 2"
	handlers := &ConreteHandlerA{
		next: &ConcreteHandlerB{
			next: &ConcreteHandlerC{},
		},
	}
	result := handlers.SendRequest(2)
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
