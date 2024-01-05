package day04

import (
	"strconv"
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

func partOne(d []string) (string, error) {
	res := 0
	for _, card := range d {
		val := 0
		temp := strings.Split(card, ": ")
		nums := strings.Split(temp[1], " | ")
		w := strings.Split(strings.Join(strings.Fields(nums[0]), " "), " ")
		m := strings.Split(strings.Join(strings.Fields(nums[1]), " "), " ")
		for _, num := range w {
			if contains(m, num) {
				if val == 0 {
					val += 1
				} else {
					val *= 2
				}
			}
		}
		res += val
	}
	return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
	res := 0
	c := map[int]int{}
	for posC, card := range d {
		if _, ok := c[posC]; ok {
			c[posC] += 1
		} else {
			c[posC] = 1
		}
		val := 0
		temp := strings.Split(card, ": ")
		nums := strings.Split(temp[1], " | ")
		w := strings.Split(strings.Join(strings.Fields(nums[0]), " "), " ")
		m := strings.Split(strings.Join(strings.Fields(nums[1]), " "), " ")
		for _, num := range w {
			if contains(m, num) {
				val += 1
			}
		}

		for j := 0; j < c[posC]; j++ {
			for i := posC + 1; i < posC+1+val; i++ {
				if _, ok := c[i]; ok {
					c[i] += 1
				} else {
					c[i] = 1
				}
			}
		}
	}
	for _, val := range c {
		res += val
	}

	return strconv.Itoa(res), nil
}

func contains(s []string, x string) bool {
	for _, e := range s {
		if e == x {
			return true
		}
	}
	return false
}
