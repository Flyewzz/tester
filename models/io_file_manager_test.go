package models

import "testing"

func TestIOFileManager_Read(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		iofm    *IOFileManager
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.iofm.Read(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IOFileManager.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IOFileManager.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIOFileManager_Write(t *testing.T) {
	type args struct {
		text string
		path string
	}
	tests := []struct {
		name    string
		iofm    *IOFileManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.iofm.Write(tt.args.text, tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("IOFileManager.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
