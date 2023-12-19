package day9

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

func Solve(d string) (string, string, error) {
    d = strings.TrimSpace(d)
    t := utils.SplitLines(d)
    s, err := parse(t)

    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

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

func partOne(d [][]int) (string, error) {
    res := 0
    for _, line := range d {
        tv := [][]int{}
        tv = append(tv, line)

        tl := make([]int, len(line))
        copy(tl, line)

        for {
            temp := []int{}

            if allZero(tl) {
                break
            }

            for i := 0; i < len(tl) - 1; i++ {
                temp = append(temp, tl[i + 1] - tl[i])
            }

            tl = make([]int, len(temp))
            copy(tl, temp)

            tv = append(tv, tl)
        }

        for _, e := range tv {
            res += e[len(e) - 1]
        }
    }

    return strconv.Itoa(res), nil
}

func partTwo(d [][]int) (string, error) {
    res := 0
    for _, line := range d {
        tv := [][]int{}
        tv = append(tv, line)

        tl := make([]int, len(line))
        copy(tl, line)

        for {
            temp := []int{}

            if allZero(tl) {
                break
            }

            for i := 0; i < len(tl) - 1; i++ {
                temp = append(temp, tl[i + 1] - tl[i])
            }

            tl = make([]int, len(temp))
            copy(tl, temp)

            tv = append(tv, tl)
        }
        for i := len(tv) - 1; i >= 0; i-- {
            if i == len(tv) - 1 {
                tv[i] = append([]int{0}, tv[i]...)
                continue
            }
            if i == 0 {
                res += tv[i][0] - tv[i + 1][0]
                break
            }
            tv[i] = append([]int{tv[i][0] - tv[i + 1][0]}, tv[i]...)
        }

    }

    return strconv.Itoa(res), nil
}

func allZero(n []int) bool {
    if len(n) < 1 {
        return false
    }
    for _, e := range n {
        if e != 0 {
            return false
        }
    }
    return true
}

func parse(d []string) ([][]int, error) {
    nums := [][]int{}
    for _, seq := range d {
        tnums := []int{}
        ts := strings.Split(seq, " ")
        for _, n := range ts {
            tn, err := strconv.Atoi(n)
            if err != nil {
                return nil, err
            }
            tnums = append(tnums, tn)
        }
        nums = append(nums, tnums)
    }
    return nums, nil
}
