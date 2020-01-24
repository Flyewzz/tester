package models

import (
	"reflect"
	"testing"
)

func TestNewMessage(t *testing.T) {
	type args struct {
		answer string
		passed bool
	}
	tests := []struct {
		name string
		args args
		want *Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMessage(tt.args.answer, tt.args.passed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
