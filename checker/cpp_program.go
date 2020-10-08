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

	"github.com/Flyewzz/tester/models"
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

func (p *CppProgram) Run(ctx context.Context, input string) (models.Result, error) {
	// filepath.Dir(p.Path))) - a directory contains the executing program
	path, _ := filepath.Abs(filepath.Dir(p.Path))
	cmd := exec.Command("bash", "-c", fmt.Sprintf("timeout 15 docker run --rm -i --memory=%s --memory-swap %s --cpus=%s "+
		"-v %s:/program frolvlad/alpine-gxx "+
		"/bin/sh -c \"g++ program/main.cpp && ./a.out\"",
		p.MemoryLimit+"m", p.DiskLimit+"m", p.CpuLimit, path))
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	errCh := make(chan error, 1)
	ready := make(chan struct{})
	go func() {
		ready <- struct{}{}
		errCh <- cmd.Run()
	}()
	<-ready
	var err error
	select {
	case <-ctx.Done():
		if err := cmd.Process.Kill(); err != nil {
			log.Printf("failed to kill process: %v\n", err)
		} else {
			log.Println("Correct closing")
		}
	case err = <-errCh:
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + stderr.String())

			return models.Result{
				Out:      "",
				ExitCode: cmd.ProcessState.ExitCode(),
			}, errors.New(fmt.Sprint(err) + ": " + stderr.String())
		}
	}
	return models.Result{
		Out:      out.String(),
		ExitCode: cmd.ProcessState.ExitCode(),
	}, nil
}

func (p *CppProgram) Check(tests []*Test) []*Verdict {
	fmt.Println("----------------------")
	var verdicts []*Verdict
	type Answer struct {
		result models.Result
		err    error
	}
TESTLOOP:
	for _, test := range tests {
		answerCh := make(chan Answer, 1)
		answer := Answer{}
		timer := time.NewTimer(time.Duration(p.TimeLimit) * time.Millisecond)
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

			if answer.result.ExitCode == 139 || answer.result.ExitCode == 137 {
				statusAnswer = "ML" // Memory limit
			}

			verdict = NewVerdict(test.Name, statusAnswer)
			verdict.Message = answer.err.Error()
			verdicts = append(verdicts, verdict)
			break
		}
		if test.Output != answer.result.Out {
			verdict = NewVerdict(test.Name, "WA")
		} else {
			verdict = NewVerdict(test.Name, "OK")
		}
		fmt.Printf("Result for %s: %s\n", test.Name, verdict.Status)
		verdicts = append(verdicts, verdict)
		if verdict.Status == "WA" {
			break
		}
	}
	return verdicts
}
