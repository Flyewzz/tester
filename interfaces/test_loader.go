package interfaces

import "github.com/Flyewzz/tester/checker"

type TestLoader interface {
	Load(id int) []*checker.Test
}