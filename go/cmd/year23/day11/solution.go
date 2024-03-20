package day11

import (
	"fmt"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

func Solve(d string) (string, string, error) {
	s := utils.SplitLines(strings.TrimSpace(d))

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

func partOne(s []string) (string, error) {
	return "", fmt.Errorf("not implemented")
}

func partTwo(s []string) (string, error) {
	return "", fmt.Errorf("not implemented")
}
