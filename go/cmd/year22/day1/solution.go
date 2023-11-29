package day1

import (
	"sort"
	"strconv"
	"strings"

	"github.com/stefan-bordei/goAOC/cmd/utils"
)

func Solve(d string) (string, string, error) {
    d = strings.TrimSpace(d)
    part_one, err := partOne(d)

    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    part_two, err := partTwo(d)

    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return part_one, part_two, nil
}

func partOne(d string) (string, error) {
    res := utils.SplitByEmptyNewline(d)
    max, sum := 0, 0

    for _, e := range res {
        calories := strings.Split(e, "\n")
        for _, f := range calories {
            cal_int, err := strconv.Atoi(f)
            if err != nil {
                return utils.NOT_DONE, err
            }
            sum += cal_int
        }
        if sum > max {
            max = sum
        }
        sum = 0
    }
    return strconv.Itoa(max), nil
}

func partTwo(d string) (string, error) {
    elfs := utils.SplitByEmptyNewline(d)
    var res []int
    sum := 0

    for _, e := range elfs {
        calories := strings.Split(e, "\n")
        for _, cal := range calories {
            cal_int, err := strconv.Atoi(cal)
            if err != nil {
                return utils.NOT_DONE, err
            }
            sum += cal_int
        }
        res = append(res, sum)
        sum = 0
    }

    sort.Slice(res, func(i, j int) bool {
        return res[i] > res[j]
    })

    return strconv.Itoa(utils.SumUp(res[:3])), nil
}

