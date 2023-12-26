package desafio_testes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// implemententando um caso de testes que possa cubrir uma ampla gama de casos de testes
func TestSumTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "should sum correctly", input: []int{1, 2, 3}, want: 6},
		{name: "should sum correctly", input: []int{1, 2, 4}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := sum(tt.input...)
			assert.Equal(t, tt.want, total)
		})
	}
}

func TestSubstractionTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "should substract correctly", input: []int{1, 2, 3}, want: -6},
		{name: "should substract correctly", input: []int{1, 2, 4}, want: -7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := substraction(tt.input...)
			assert.Equal(t, tt.want, total)
		})
	}
}

func TestMultiplicationTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "should multiply correctly", input: []int{1, 2, 3}, want: 6},
		{name: "should multiply correctly", input: []int{1, 2, 14}, want: 28},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := multiplication(tt.input...)
			assert.Equal(t, tt.want, total)
		})
	}
}

func TestDivideTableDriven(t *testing.T) {
	tests := []struct {
		name  string
		input []float64
		want  float64
	}{
		{name: "should divide correctly", input: []float64{1, 2}, want: 0.5},
		{name: "should divide correctly", input: []float64{14, 2}, want: 7.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total := division(tt.input[0], tt.input[1])
			assert.Equal(t, tt.want, total)
		})
	}
}
