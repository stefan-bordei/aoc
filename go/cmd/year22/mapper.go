package mapper22

import (
	"github.com/stefan-bordei/aoc/go/cmd/year22/day1"
	"github.com/stefan-bordei/aoc/go/cmd/year22/day2"
	"github.com/stefan-bordei/aoc/go/cmd/year22/day3"
	"github.com/stefan-bordei/aoc/go/cmd/year22/day4"
	"github.com/stefan-bordei/aoc/go/cmd/year22/day5"
	"github.com/stefan-bordei/aoc/go/cmd/year22/day6"
)

type fn func(string) (string, string, error)

func Solve(d string, i string) (string, string, error) {
	m := map[string]fn{
		"1": day1.Solve,
		"2": day2.Solve,
		"3": day3.Solve,
		"4": day4.Solve,
		"5": day5.Solve,
		"6": day6.Solve,
	}

	part_one, part_two, err := m[d](i)
	return part_one, part_two, err
}
