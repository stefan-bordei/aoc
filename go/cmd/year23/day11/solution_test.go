package day11

import "testing"

var example = `....1........
.........2...
3............
.............
.............
........4....
.5...........
.##.........6
..##.........
...##........
....##...7...
8....9.......`

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example1",
			input: example,
			want:  "374",
		},
		{
			name:  "example2",
			input: example,
			want:  "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "example1" {
				if got, _, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part1a() = %v, want %v", got, tt.want)
				}
			} else if tt.name == "example2" {
				if got, _, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part1b() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
