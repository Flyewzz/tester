package checker

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type PythonProgram struct {
	Path        string
	MemoryLimit string
	DiskLimit   string
	CpuLimit    string
	TimeLimit   int
}

func NewPythonProgram(path, memoryLimit,
	diskLimit, cpuLimit string, timeLimit int) *PythonProgram {
	return &PythonProgram{
		Path:        path,
		MemoryLimit: memoryLimit,
		DiskLimit:   diskLimit,
		CpuLimit:    cpuLimit,
		TimeLimit:   timeLimit,
	}
}

func (p *PythonProgram) Run(input string) (string, error) {

	// filepath.Dir(p.Path))) - a directory contains the executing program
	cmd := exec.Command("bash", "-c", fmt.Sprintf("docker run --rm -i --memory=%s --memory-swap %s --cpus=%s "+
		"-v %s:/program python "+
		"/bin/sh -c \"python program/main.py\"",
		p.MemoryLimit, p.DiskLimit, p.CpuLimit, filepath.Dir(p.Path)))
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.String(), nil
}

func (p *PythonProgram) Check() []*Verdict {
	const testsPath string = "checker/tests/"

	names, _ := GetTestsNames(testsPath)
	var tests []*Test
	for _, name := range names {
		test, _ := GetTest(testsPath + name)
		tests = append(tests, test)
	}

	fmt.Println("----------------------")
	var verdicts []*Verdict
	type Answer struct {
		result string
		err    error
	}
TESTLOOP:
	for _, test := range tests {
		answerCh := make(chan Answer, 1)
		answer := Answer{}
		var result string
		timer := time.NewTimer(time.Duration(p.TimeLimit+800) * time.Millisecond)
		ready := make(chan struct{})
		go func(test *Test) {
			ready <- struct{}{}
			result, err := p.Run(test.Input)
			answerCh <- Answer{
				result: result,
				err:    err,
			}
		}(test)
		var verdict *Verdict = nil
		<-ready
		select {
		case <-timer.C:
			log.Printf("Time limit exceeded for %s:\n", test.Name)
			verdict = NewVerdict(test.Name, "TL")
			verdict.Message = "Time limit exceeded"
			verdicts = append(verdicts, verdict)
			break TESTLOOP
		case answer = <-answerCh:
			timer.Stop()
		}
		if answer.err != nil {
			log.Printf("Error for %s:\n %s\n", test.Name, answer.err.Error())
			verdict = NewVerdict(test.Name, "CE")
			verdict.Message = answer.err.Error()
			verdicts = append(verdicts, verdict)
			break
		}
		fmt.Printf("Result for %s: \n %s\n", test.Name, result)
		if test.Output != answer.result {
			verdict = NewVerdict(test.Name, "WA")
		} else {
			verdict = NewVerdict(test.Name, "OK")
		}
		verdicts = append(verdicts, verdict)
		if verdict.Status == "WA" {
			break
		}
	}
	return verdicts
}
