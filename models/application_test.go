package models

import "testing"

func TestApplication_Execute(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		a    *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.Execute(tt.args.input)
		})
	}
}
