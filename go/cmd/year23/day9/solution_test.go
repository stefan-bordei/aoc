package day9

import (
	"testing"
)

var example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: example,
			want:  "114",
		},
		{
			name:  "example2",
			input: example,
			want:  "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
            if tt.name == "example" {
                if got, _, _ := Solve(tt.input); got != tt.want {
                    t.Errorf("part1() = %v, want %v", got, tt.want)
                }
            } else {
                if _, got, _ := Solve(tt.input); got != tt.want {
                    t.Errorf("part2() = %v, want %v", got, tt.want)
                }
            }
		})
	}
}

