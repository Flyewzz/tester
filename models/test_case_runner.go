package models

import (
	"github.com/Flyewzz/tester/interfaces"
)

type TestCaseRunner struct {
}

func (tcr *TestCaseRunner) Run(testCase *TestCase, app *interfaces.Executable) *Message {
	return NewMessage("", true)
}
