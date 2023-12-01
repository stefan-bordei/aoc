package day1

import (
	"testing"
)

var example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchets`

var example2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example",
			input: example,
			want:  "142",
		},
		{
			name:  "example2",
			input: example2,
			want:  "281",
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

