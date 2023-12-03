package day3

import (
	"strconv"
	"strings"
	"unicode"

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
    for posL, l := range d {
        nums := strings.FieldsFunc(l, split)
        for _, num := range nums {
            valid := false
            start := strings.Index(l, num)
            end := start + len(num)-1
            switch posL {
            case 0:
                valid = inlineCheck(l, start, end) || boundryCheck(d[posL + 1], start, end)
            case len(l)-1:
                valid = inlineCheck(l, start, end) || boundryCheck(d[posL - 1], start, end)
            default:
                valid = inlineCheck(l, start, end) || boundryCheck(d[posL - 1], start, end) || boundryCheck(d[posL + 1], start, end)
            }
            if valid == true {
                t, err := strconv.Atoi(num)
                if err != nil {
                    return utils.NOT_DONE, err
                }
                res += t
            }
            for i := start; i <= end; i++ {
                out := []rune(l)
                out[i] = rune('.')
                l = string(out)
            }
        }
    }
    return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
    return utils.NOT_DONE, nil
}

func split(r rune) bool {
    return !unicode.IsDigit(r) || r == '.'
}

func inlineCheck(s string, start, end int) bool {
    if start == 0 {
        return !unicode.IsDigit(rune(s[end + 1])) && s[end + 1] != '.'
    }

    if end >= len(s) - 1 {
        return !unicode.IsDigit(rune(s[start - 1])) && s[start - 1] != '.'
    }
    return !unicode.IsDigit(rune(s[start - 1])) && s[start - 1] != '.' ||
            !unicode.IsDigit(rune(s[end + 1])) && s[end + 1] != '.'
}

func boundryCheck(s string, start, end int) bool {
    temp := inlineCheck(s, start, end)
    if temp == false {
        for _, c := range s[start:end+1] {
            if !unicode.IsDigit(c) && c != '.' {
                return true
            }
        }
    }
    return temp
}

