package day1

import (
	"strings"

	"github.com/stefan-bordei/goAOC/cmd/utils"
)

func Solve(d string) (string, string, error) {
    d = strings.TrimSpace(d)
    part_one, err := partOne(d)

    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    part_two, err := partTwo(d)

    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return part_one, part_two, nil
}

func partOne(d string) (string, error) {
    return utils.NOT_DONE, nil
}

func partTwo(d string) (string, error) {
    return utils.NOT_DONE, nil
}


