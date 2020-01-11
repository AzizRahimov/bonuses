package main

import "testing"

func Test_bonus(t *testing.T) {

	tests := []struct {
		name  string
		sales []int
		want  int
	}{
		// TODO: Add test cases.
		{"bonus received", []int{12_000, 4_000}, 100},
		{"no bonus", []int{10_000, 10_000}, 0},
		{"for all values", []int{12_000, 15_000}, 350},
	}
	for _, tt := range tests {
		got := bonus(tt.sales)

		if got != tt.want {
			t.Error("for bonus test", tt.name, "got", got, "want", tt.want)
		}
	}
}
