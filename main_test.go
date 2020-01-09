package main

import "testing"

func Test_bonus(t *testing.T) {

	tests := []struct {
		name string
		moneyBonus int
		want int
	}{
		// TODO: Add test cases.
		{"bonus received", 12_000, 100},
		{"no bonus", 8_000, 0},
		{"min bonus", 10_020, 1 },

	}
	for _, tt := range tests {
		got :=bonus(tt.moneyBonus)

			if got!= tt.want {
				t.Error("for bonus test", tt.name, "got", got, "want", tt.want)
			}
		}
	}
