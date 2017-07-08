package main

import (
	"reflect"
	"testing"
)

func TestMapUnsafe(t *testing.T) {
	type args struct {
		fn   func(interface{}) interface{}
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapUnsafe(tt.args.fn, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapUnsafe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapUnsafeParallel(t *testing.T) {
	type args struct {
		fn   func(interface{}) interface{}
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapUnsafeParallel(tt.args.fn, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapUnsafeParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
