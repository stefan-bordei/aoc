package mapper23

import (
	"github.com/stefan-bordei/aoc/go/cmd/year23/day1"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day2"
	"github.com/stefan-bordei/aoc/go/cmd/year23/day3"
)


type fn func (string) (string, string, error)


func Solve(d string, i string) (string, string, error) {
    m := map[string] fn {
        "1": day1.Solve,
        "2": day2.Solve,
        "3": day3.Solve,
    }


    part_one, part_two, err := m[d](i)
    return part_one, part_two, err
}

