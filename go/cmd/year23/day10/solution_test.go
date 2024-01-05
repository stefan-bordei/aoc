package day10

import "testing"

var example = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

var example2 = `.....
.S-7.
.|.|.
.L-J.
.....`

var example3 = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

var example4 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

var example5 = `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`

var example6 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

var example7 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestSolution(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "example1",
			input: example,
			want:  "8",
		},
		{
			name:  "example2",
			input: example2,
			want:  "4",
		},
		{
			name:  "example3",
			input: example3,
			want:  "8",
		},
		{
			name:  "example4",
			input: example4,
			want:  "4",
		},
		{
			name:  "example5",
			input: example5,
			want:  "4",
		},
		{
			name:  "example6",
			input: example6,
			want:  "8",
		},
		{
			name:  "example7",
			input: example7,
			want:  "10",
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
			} else if tt.name == "example3" {
				if got, _, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part1c() = %v, want %v", got, tt.want)
				}
			} else if tt.name == "example4" {
				if _, got, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part2a() = %v, want %v", got, tt.want)
				}
			} else if tt.name == "example5" {
				if _, got, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part2b() = %v, want %v", got, tt.want)
				}
			} else if tt.name == "example6" {
				if _, got, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part2c() = %v, want %v", got, tt.want)
				}
			} else if tt.name == "example7" {
				if _, got, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part2d() = %v, want %v", got, tt.want)
				}
			} else {
				if _, got, _ := Solve(tt.input); got != tt.want {
					t.Errorf("part2e() = %v, want %v", got, tt.want)
				}
			}

		})
	}
}
