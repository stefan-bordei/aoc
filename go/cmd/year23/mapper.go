package mapper23

import (
	"github.com/stefan-bordei/goAOC/cmd/year23/day1"
	"github.com/stefan-bordei/goAOC/cmd/year23/day2"
)


type fn func (string) (string, string, error)


func Solve(d string, i string) (string, string, error) {
    m := map[string] fn {
        "1": day1.Solve,
        "2": day2.Solve,
    }


    part_one, part_two, err := m[d](i)
    return part_one, part_two, err
}

