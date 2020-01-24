package models

import (
	"reflect"
	"testing"

	"github.com/Flyewzz/tester/interfaces"
)

func TestTestCaseRunner_Run(t *testing.T) {
	type args struct {
		testCase *TestCase
		app      *interfaces.Executable
	}
	tests := []struct {
		name string
		tcr  *TestCaseRunner
		args args
		want *Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tcr.Run(tt.args.testCase, tt.args.app); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestCaseRunner.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
