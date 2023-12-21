package mapper23

import (
	"github.com/stefan-bordei/aoc/go/cmd/year23/day1"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day2"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day3"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day4"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day5"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day6"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day7"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day8"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day9"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day10"
)


type fn func (string) (string, string, error)


func Solve(d string, i string) (string, string, error) {
    m := map[string] fn {
        "1": day1.Solve,
        "2": day2.Solve,
        "3": day3.Solve,
        "4": day4.Solve,
        "5": day5.Solve,
        "6": day6.Solve,
        "7": day7.Solve,
        "8": day8.Solve,
        "9": day9.Solve,
        "10": day10.Solve,
    }


    part_one, part_two, err := m[d](i)
    return part_one, part_two, err
}

