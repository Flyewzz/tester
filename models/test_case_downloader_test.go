package models

import (
	"reflect"
	"testing"
)

func TestTestCaseDownloader_Download(t *testing.T) {
	type args struct {
		dirPath string
	}
	tests := []struct {
		name string
		tcd  *TestCaseDownloader
		args args
		want *TestCase
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tcd.Download(tt.args.dirPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TestCaseDownloader.Download() = %v, want %v", got, tt.want)
			}
		})
	}
}
