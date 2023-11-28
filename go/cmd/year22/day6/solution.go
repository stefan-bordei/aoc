package day6

import (
	"fmt"

	"github.com/stefan-bordei/goAOC/cmd/utils"
)

func Solve(d string) (string, string, error) {
    s := utils.SplitLines(d)[0]

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

func partOne(s string) (string, error) {
    res, err := findMarkerIndex(s, 4)
    if err != nil {
        return utils.NOT_DONE, err
    }
    return fmt.Sprintf("%d", res), nil
}

func partTwo(s string) (string, error) {
    res, err := findMarkerIndex(s, 14)
    if err != nil {
        return utils.NOT_DONE, err
    }
    return fmt.Sprintf("%d", res), nil
}

func findMarkerIndex(s string, ml int) (int, error) {
    for i := 0; i <= len(s)-ml; i++ {
        if checkUniqueElements(s[i:i+ml]) {
            return i+ml, nil
        }
    }
    return 0, nil
}

func checkUniqueElements(s string) bool {
    var Empty struct{}
    m := make(map[rune]struct{})
    for _, v := range s {
        if _, ok := m[v]; ok {
           return false
        }
        m[v] = Empty
    }
    return true
}

