package main

import (
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	type args struct {
		s string
		m *sync.Mutex
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateMessage(tt.args.s, tt.args.m)
		})
	}
}
