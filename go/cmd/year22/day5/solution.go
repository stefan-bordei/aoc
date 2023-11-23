package day5

import (
	"fmt"
	"strings"

	"github.com/stefan-bordei/goAOC/cmd/utils"
)

func Solve(d string) (string, string, error) {
    s:= utils.SplitLines(d)
    var stacks int = 9
    var crates int = 8

    part_one, err := partOne(s[:crates-1], s[crates+1:], stacks)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    part_two, err := partTwo(s)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return part_one, part_two, nil
}

func partOne(s, d []string, i int) (string, error) {
    for _, i := range s {
        fmt.Println(strings.SplitN(i, " ", 9))
    }
    return utils.NOT_DONE, nil
}

func partTwo(d []string) (string, error) {
    return utils.NOT_DONE, nil
}

