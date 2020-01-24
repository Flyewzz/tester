package models

import (
	"reflect"
	"testing"

	"github.com/Flyewzz/tester/interfaces"
)

func TestNewTestCase(t *testing.T) {
	tests := []struct {
		name string
		want *TestCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTestCase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTestCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestCase_AddTest(t *testing.T) {
	type args struct {
		t *interfaces.Test
	}
	tests := []struct {
		name string
		tc   *TestCase
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.tc.AddTest(tt.args.t)
		})
	}
}
