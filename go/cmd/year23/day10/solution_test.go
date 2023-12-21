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

var example3 =`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func TestSolution(t *testing.T) {
    tests := []struct {
        name string
        input string
        want string
    }{
        {
            name: "example1",
            input: example,
            want: "8",
        },
        {
            name: "example2",
            input: example2,
            want: "4",
        },
        {
            name: "example3",
            input: example3,
            want: "8",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.name == "example1" {
                if got, _, _ := Solve(tt.input); got != tt.want {
                    t.Errorf("part1() = %v, want %v", got, tt.want)
                }
            } else if tt.name == "example2" {
                if got, _, _ := Solve(tt.input); got != tt.want {
                    t.Errorf("part1a() = %v, want %v", got, tt.want)
                }
            } else if tt.name == "example3" {
                if got, _, _ := Solve(tt.input); got != tt.want {
                    t.Errorf("part1b() = %v, want %v", got, tt.want)
                }
            } else {
                if _, got, _ := Solve(tt.input); got != tt.want {
                    t.Errorf("part2() = %v, want %v", got, tt.want)
                }
            }

        })
    }
}
