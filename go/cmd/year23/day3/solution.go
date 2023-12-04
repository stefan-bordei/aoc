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
        nums := extractNums(l)
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
            l = replaceSlice(l, start, end, rune('.'))
        }
    }
    return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
    res := 0
    for posL, l := range d {
        for posC, c := range l {
            if c == '*' {
                res += gearsValue(d, posC, posL)
            }
        }
    }
    return strconv.Itoa(res), nil
}

func gearsValue(s []string, x, y int) int {
    res := 0

    vars := checkGears(s, x, y)
    if len(vars) == 2 {
        tempR := 1
        for _, v := range vars {
            temp, err := strconv.Atoi(v)
            if err != nil {
                panic(err)
            }
            tempR *= temp
        }
        res += tempR
    }
    return res
}

func checkGears(s []string, x, y int) []string {
    res := []string{}
    switch y {
    case 0:
        res = append(res, inlineGears(s[y], x)...)
        res = append(res, inlineGears(s[y + 1], x)...)
    case len(s) - 1:
        res = append(res, inlineGears(s[y - 1], x)...)
        res = append(res, inlineGears(s[y], x)...)
    default:
        res = append(res, inlineGears(s[y - 1], x)...)
        res = append(res, inlineGears(s[y], x)...)
        res = append(res, inlineGears(s[y + 1], x)...)
    }

    return res
}

func inlineGears(s string, x int) []string {
    res := []string{}
    nums := extractNums(s)
    for _, num := range nums {
        start := strings.Index(s, num)
        end := start + len(num) - 1
        if end >= x-1 && start <= x+1 {
            res = append(res, num)
        }
        s = replaceSlice(s, start, end, rune('.'))
    }
    return res
}

func split(r rune) bool {
    return !unicode.IsDigit(r) || r == '.'
}

func replaceSlice(s string, start, end int, r rune) string {
    for i := start; i <= end; i++ {
        out := []rune(s)
        out[i] = r
        s = string(out)
    }
    return s
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

func extractNums(s string) []string {
    return strings.FieldsFunc(s, split)
}

