package checker

type Verdict struct {
	TestName string `json:"name"`
	Status   string `json:"status"`
	Message  string `json:"message,omitempty"`
}

func NewVerdict(name, status string) *Verdict {
	return &Verdict{
		TestName: name,
		Status:   status,
	}
}
