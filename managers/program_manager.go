package managers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Flyewzz/tester/checker"
	uuid "github.com/satori/go.uuid"
)

type ProgramManager struct {
}

func (this ProgramManager) Create(
	ram, hdd, cpu string,
	timeLimit int, code string) (checker.Program, error) {
	folderPath := "checker/task/" +
		strings.Replace(uuid.NewV4().String(),
			"-", "", -1)

	programPath := filepath.Join(folderPath, "main.cpp")

	err := os.Mkdir(folderPath, 0700)
	if err != nil {
		return nil, err
	}

	programFile, err := os.Create(programPath)
	if err != nil {
		os.RemoveAll(folderPath)
		return nil, err
	}
	_, err = programFile.WriteString(code)
	if err != nil {
		programFile.Close()
		os.RemoveAll(folderPath)
		return nil, err
	}
	return checker.NewCppProgram(
		programPath,
		ram,
		hdd,
		".800",
		timeLimit,
	), nil
}
