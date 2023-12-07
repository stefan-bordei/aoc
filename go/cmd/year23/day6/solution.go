package day6

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

func Solve(d string) (string, string, error) {
    s := utils.SplitLines(strings.TrimSpace(d))

    times := []int{}
    timesTemp := strings.Fields(strings.Split(s[0], ": ")[1])
    for _, t := range timesTemp {
        time, err := strconv.Atoi(t)
        if err != nil {
            return utils.NOT_DONE, utils.NOT_DONE, err
        }
        times = append(times, time)
    }

    distances := []int{}
    distancesTemp := strings.Fields(strings.Split(s[1], ": ")[1])
    for _, d := range distancesTemp {
        dist, err := strconv.Atoi(d)
        if err != nil {
            return utils.NOT_DONE, utils.NOT_DONE, err
        }
        distances = append(distances, dist)
    }

    partOne, err := partOne(times, distances)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partTwo, err := partTwo(times, distances)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return partOne, partTwo, nil
}

func partOne(t, d []int) (string, error) {
    res := 1
    for i := 0; i < len(t); i++ {
        m := 0
        for h := 0; h < t[i]; h++ {
            x := t[i] - h
            if h * x > d[i] {
                m += 1
            }
        }
        res *= m
    }
    return strconv.Itoa(res), nil
}

func partTwo(t, d []int) (string, error) {
    tempT, err := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(t), " ", "", -1), "[]"))
    if err != nil {
        return utils.NOT_DONE, err
    }
    tempD, err := strconv.Atoi(strings.Trim(strings.Replace(fmt.Sprint(d), " ", "", -1), "[]"))
    if err != nil {
        return utils.NOT_DONE, err
    }
    time := []int{tempT}
    dist := []int{tempD}
    return partOne(time, dist)
}
