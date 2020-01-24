package models

import (
	"reflect"
	"testing"
)

func TestNewFileTest(t *testing.T) {
	type args struct {
		exp string
	}
	tests := []struct {
		name string
		args args
		want *FileTest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileTest(tt.args.exp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileTest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileTest_Run(t *testing.T) {
	tests := []struct {
		name string
		ft   *FileTest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ft.Run()
		})
	}
}
