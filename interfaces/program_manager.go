package interfaces

import (
	"github.com/Flyewzz/tester/checker"
)

type ProgramManager interface {
	Create(ram, hdd, cpu string,
		timeLimit int, code string) (checker.Program, error)
}
