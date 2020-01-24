package models

import (
	"github.com/Flyewzz/tester/interfaces"
)

type TestCase struct {
	tests []interfaces.Test
}

func NewTestCase() *TestCase {
	return &TestCase{}
}

func (tc *TestCase) AddTest(t *interfaces.Test) {

}
