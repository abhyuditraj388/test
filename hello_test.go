package main

import (
	"math"
	"testing"
)

func TestDiv(t *testing.T) {
	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"simple", 4, 2, 2},
		{"negative", -6, 2, -3},
		{"fraction", 5, 2.5, 2},
		{"divide by zero", 5, 0, math.Inf(1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := div(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("div(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = div(10, 3)
	}
}
