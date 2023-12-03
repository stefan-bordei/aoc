package day2

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

var Colors = map[string]int{
    "red": 12,
    "green": 13,
    "blue": 14,
}

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
    for game, c := range d {
        valid := true
        g := strings.Split(c, ": ")[1]
        p := strings.FieldsFunc(g, splitOnTwo)

        for _, x := range p {
            got := strings.Split(strings.TrimSpace(x), " ")
            count, err := strconv.Atoi(got[0])
            if err != nil {
                return utils.NOT_DONE, err
            }
            if count > Colors[got[1]] {
                valid = false
            }
        }
        if valid {
            res += game + 1
        }
    }
    return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
    res := 0
    for _, c := range d {
        count := map[string]int{
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        game := 1
        g := strings.Split(c, ": ")[1]
        p := strings.FieldsFunc(g, splitOnTwo)

        for _, x := range p {
            got := strings.Split(strings.TrimSpace(x), " ")
            num, err := strconv.Atoi(got[0])
            if err != nil {
                return utils.NOT_DONE, err
            }
            if count[got[1]] < num {
                count[got[1]] = num
            }
        }

        for _, v := range count {
            game *= v
        }
        res += game
    }
    return strconv.Itoa(res), nil
}

func splitOnTwo(r rune) bool {
    return r == ',' || r == ';'
}
