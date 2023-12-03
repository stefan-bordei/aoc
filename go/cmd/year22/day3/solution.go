package day3

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

type backpackError struct {
    s string
}

func (e *backpackError) Error() string {
    return e.s
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

    for _, i := range d {
        if len(i[:len(i) / 2]) != len(i[len(i) / 2:]) {
            return utils.NOT_DONE, &backpackError{"Backpack pockets not equal in len."}
        }

        common := commonRunes(i[:len(i) / 2], i[len(i) / 2:])
        p := prepRune(common[0]) // always take the first one in case of duplicates

        res += p
    }

    return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
    res := 0

    if len(d) % 3 != 0 {
        return utils.NOT_DONE, &backpackError{"Backpacks length not modulo of 3."}
    }

    i := 0
    for i <= len(d) - 2 {
        first := commonRunes(d[i], d[i + 1])
        g := commonRunes(string(first), d[i + 2])

        res += prepRune(g[0])
        i += 3
    }

    return strconv.Itoa(res), nil
}

func prepRune(r rune) int {
    i := int(r)
    if i > 90 {
        i -= 96
    } else {
        i -= 38
    }
    return i
}

func commonRunes(a, b string) []rune {
    var res []rune
    for _, c := range a {
        if strings.ContainsRune(b, c) {
            res = append(res, c)
        }
    }
    return res
}

