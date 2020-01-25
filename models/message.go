package models

type Message struct {
	Status string
	Passed bool
}

func NewMessage(Status string, passed bool) *Message {
	return &Message{
		Status: Status,
		Passed: passed,
	}
}
