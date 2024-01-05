package mapper23

import (
	"github.com/stefan-bordei/aoc/go/cmd/year23/day01"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day02"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day03"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day04"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day05"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day06"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day07"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day08"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day09"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day10"
)

type fn func(string) (string, string, error)

func Solve(d string, i string) (string, string, error) {
	m := map[string]fn{
		"1":  day01.Solve,
		"2":  day02.Solve,
		"3":  day03.Solve,
		"4":  day04.Solve,
		"5":  day05.Solve,
		"6":  day06.Solve,
		"7":  day07.Solve,
		"8":  day08.Solve,
		"9":  day09.Solve,
		"10": day10.Solve,
	}

	part_one, part_two, err := m[d](i)
	return part_one, part_two, err
}
