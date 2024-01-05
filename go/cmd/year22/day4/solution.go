package day4

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

func Solve(d string) (string, string, error) {
	d = strings.TrimSpace(d)
	s := utils.SplitLines(d)

	part_one, err := partOne(s)
	if err != nil {
		return utils.NOT_DONE, utils.NOT_DONE, err
	}

	part_two, err := partTwo(s)
	if err != nil {
		return utils.NOT_DONE, utils.NOT_DONE, err
	}

	return part_one, part_two, nil
}

func partOne(d []string) (string, error) {
	var res int

	for _, s := range d {
		splits := strings.Split(s, ",")
		l, r := splits[0], splits[1]

		l_split, err := toInt(strings.Split(l, "-"))
		if err != nil {
			return utils.NOT_DONE, err
		}

		r_split, err := toInt(strings.Split(r, "-"))
		if err != nil {
			return utils.NOT_DONE, err
		}

		if l_split[0] >= r_split[0] &&
			l_split[1] <= r_split[1] ||
			r_split[0] >= l_split[0] &&
				r_split[1] <= l_split[1] {
			res += 1
		}
	}

	return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
	var res int

	for _, s := range d {
		splits := strings.Split(s, ",")
		l, r := splits[0], splits[1]

		l_split, err := toInt(strings.Split(l, "-"))
		if err != nil {
			return utils.NOT_DONE, err
		}

		r_split, err := toInt(strings.Split(r, "-"))
		if err != nil {
			return utils.NOT_DONE, err
		}

		// check left
		if l_split[0] >= r_split[0] &&
			l_split[0] <= r_split[1] ||
			l_split[1] >= r_split[0] &&
				l_split[1] <= r_split[1] ||
			r_split[0] >= l_split[0] &&
				r_split[0] <= l_split[1] ||
			r_split[1] >= l_split[0] &&
				r_split[1] <= l_split[1] {
			res += 1
		}
	}
	return strconv.Itoa(res), nil
}

func toInt(s []string) ([]int, error) {
	var res []int
	for _, i := range s {
		tmp, err := strconv.Atoi(i)
		if err != nil {
			return []int{}, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
