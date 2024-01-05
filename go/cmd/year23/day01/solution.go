package day01

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

var digitMapping = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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
	res := 0
	for _, l := range d {
		first, err := firstDigit(l)
		if err != nil {
			return utils.NOT_DONE, err
		}
		second, err := lastDigit(l)
		if err != nil {
			return utils.NOT_DONE, err
		}
		res += 10*int(first-'0') + int(second-'0')
	}
	return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
	res := 0
	for _, l := range d {
		first, err := trueFirstDigit(l)
		if err != nil {
			return utils.NOT_DONE, err
		}
		second, err := trueLastDigit(l)
		if err != nil {
			return utils.NOT_DONE, err
		}
		res += 10*int(first-'0') + int(second-'0')
	}
	return strconv.Itoa(res), nil
}

func reverse(s string) string {
	var out string
	for _, r := range s {
		out = string(r) + out
	}
	return out
}

func firstDigit(s string) (rune, error) {
	for _, c := range s {
		if unicode.IsDigit(c) {
			return c, nil
		}
	}
	return rune(0), nil
}

func lastDigit(s string) (rune, error) {
	return firstDigit(reverse(s))
}

func stringDigits(s string, m map[int]rune, dm map[string]rune) {
	for k, v := range dm {
		if strings.Contains(s, k) {
			m[strings.Index(s, k)] = v
			m[strings.LastIndex(s, k)] = v
		}
	}
}

func trueFirstDigit(s string) (rune, error) {
	d := make(map[int]rune)

	stringDigits(s, d, digitMapping)
	for pos, c := range s {
		if unicode.IsDigit(c) {
			d[pos] = c
		}
	}

	f := 0
	fR := '\n'
	for k, v := range d {
		if fR == '\n' {
			f = k
			fR = v
			continue
		}
		if k < f {
			fR = v
			f = k
		}
	}
	return fR, nil
}

func trueLastDigit(s string) (rune, error) {
	d := make(map[int]rune)

	stringDigits(s, d, digitMapping)
	for pos, c := range s {
		if unicode.IsDigit(c) {
			d[pos] = c
		}
	}

	f := 0
	lR := '\n'
	for k, v := range d {
		if lR == '\n' {
			f = k
			lR = v
			continue
		}
		if k > f {
			lR = v
			f = k
		}
	}
	return lR, nil
}
