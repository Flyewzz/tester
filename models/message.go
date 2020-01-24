package models

type Message struct {
	Answer string
	Passed bool
}

func NewMessage(answer string, passed bool) *Message {
	return &Message{
		Answer: answer,
		Passed: passed,
	}
}
