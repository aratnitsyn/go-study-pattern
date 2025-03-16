package chain_of_responsibility

type Handler interface {
	SendRequest(message int) string
}

type ConreteHandlerA struct {
	next Handler
}

func (h *ConreteHandlerA) SendRequest(message int) (result string) {
	if message == 1 {
		result = "I'm handler 1"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) SendRequest(message int) (result string) {
	if message == 2 {
		result = "I'm handler 2"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}

type ConcreteHandlerC struct {
	next Handler
}

func (h *ConcreteHandlerC) SendRequest(message int) (result string) {
	if message == 3 {
		result = "I'm handler 3"
	} else if h.next != nil {
		result = h.next.SendRequest(message)
	}
	return
}
