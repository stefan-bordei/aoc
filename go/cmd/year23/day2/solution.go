package day2

import (
	"strings"

	"github.com/stefan-bordei/goAOC/cmd/utils"
)

func Solve(d string) (string, string, error) {
    s := strings.TrimSpace(d)

    partOne, err := partOne(s)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partTwo, err := partTwo(s)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return partOne, partTwo, nil
}

func partOne(d string) (string, error) {
    return utils.NOT_DONE, nil
}

func partTwo(d string) (string, error) {
    return utils.NOT_DONE, nil
}
