package checker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type CppProgram struct {
	Path        string
	MemoryLimit string
	DiskLimit   string
	CpuLimit    string
	TimeLimit   int
}

func NewCppProgram(path, memoryLimit,
	diskLimit, cpuLimit string, timeLimit int) *CppProgram {
	return &CppProgram{
		Path:        path,
		MemoryLimit: memoryLimit,
		DiskLimit:   diskLimit,
		CpuLimit:    cpuLimit,
		TimeLimit:   timeLimit,
	}
}

func (p *CppProgram) Run(ctx context.Context, input string) (string, error) {

	// filepath.Dir(p.Path))) - a directory contains the executing program
	cmd := exec.Command("bash", "-c", fmt.Sprintf("docker run --rm -i --memory=%s --memory-swap %s --cpus=%s "+
		"-v %s:/program frolvlad/alpine-gxx "+
		"/bin/sh -c \"g++ program/main.cpp && ./a.out\"",
		p.MemoryLimit, p.DiskLimit, p.CpuLimit, filepath.Dir(p.Path)))
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errCh := make(chan error, 1)
	go func() {
		errCh <- cmd.Run()
	}()
	var err error
	select {
	case <-ctx.Done():
		cmd.Process.Kill()
	case err = <-errCh:
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
			return "", errors.New(fmt.Sprint(err) + ": " + stderr.String())
		}
	}
	return out.String(), nil
}

func (p *CppProgram) Check() []*Verdict {
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
		timer := time.NewTimer(time.Duration(p.TimeLimit+1800) * time.Millisecond)
		ready := make(chan struct{})
		ctx, cancel := context.WithCancel(context.Background())
		go func(test *Test) {
			ready <- struct{}{}
			result, err := p.Run(ctx, test.Input)
			answerCh <- Answer{
				result: result,
				err:    err,
			}
		}(test)
		var verdict *Verdict = nil
		<-ready
		select {
		case <-timer.C:
			cancel() // Kill the process
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
			statusAnswer := "CE"

			if strings.Contains(answer.err.Error(), "137") {
				statusAnswer = "ML" // Memory limit
			}

			verdict = NewVerdict(test.Name, statusAnswer)
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
