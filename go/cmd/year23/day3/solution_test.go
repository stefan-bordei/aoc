package day3

import "testing"

var example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestSolution(t *testing.T) {
    tests := []struct {
        name string
        input string
        want string
    }{
        {
            name: "example1",
            input: example,
            want: "4361",
        },
        {
            name: "example2",
            input: example,
            want: "467835",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if tt.name == "example1" {
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
