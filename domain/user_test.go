package domain

import (
	"reflect"
	"testing"
)

func TestGetMatt(t *testing.T) {
	tests := []struct {
		name string
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMatt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetJohn(t *testing.T) {
	tests := []struct {
		name string
		want *User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetJohn(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetJohn() = %v, want %v", got, tt.want)
			}
		})
	}
}
