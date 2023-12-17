package day7

import "testing"

var example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestSolution(t *testing.T) {
    tests := []struct {
        name string
        input string
        want string
    }{
        {
            name: "example1",
            input: example,
            want: "6440",
        },
        {
            name: "example2",
            input: example,
            want: "71503",
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
