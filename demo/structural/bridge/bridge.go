package bridge

import "fmt"

type AbstructMessage interface {
	SendMessage(text, to string)
}

type MessageTmplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func ViaSMS() MessageTmplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text string, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

type MessageEmail struct {
}

func ViaEmail() MessageTmplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text string, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

type CommonMessage struct {
	method MessageTmplementer
}

func NewCommonMesage(method MessageTmplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageTmplementer
}

func NewUrgencyMessage(method MessageTmplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency %s]", text), to)
}
