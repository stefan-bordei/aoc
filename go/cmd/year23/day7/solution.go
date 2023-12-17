package day7

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

var Strength = map[rune]int{
    '2': 2,
    '3': 3,
    '4': 4,
    '5': 5,
    '6': 6,
    '7': 7,
    '8': 8,
    '9': 9,
    'T': 10,
    'J': 11,
    'Q': 12,
    'K': 13,
    'A': 14,
}

type Hand struct {
    In string
    Cards []int
    Score int
    Type int
}

func Solve(d string) (string, string, error) {
    s := utils.SplitLines(strings.TrimSpace(d))

    c, err := parse(s, Strength, false)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partOne, err := partOne(c)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    Strength['J'] = 1
    c2, err := parse(s, Strength, true)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partTwo, err := partTwo(c2)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return partOne, partTwo, nil
}

func partOne(d []*Hand) (string, error) {
    res := 0
    sort.Slice(d[:], func(i, j int) bool {
        if d[i].Type == d[j].Type {
            for k := range d[i].Cards {
                if d[i].Cards[k] == d[j].Cards[k] { continue }
                return d[i].Cards[k] < d[j].Cards[k]
            }
        }
        return d[i].Type < d[j].Type
    })
    for i, h := range d {
        res += h.Score * (i + 1)
    }
    return strconv.Itoa(res), nil
}

func partTwo(d []*Hand) (string, error) {
    return partOne(d)
}

func parse(s []string, m map[rune]int, j bool) ([]*Hand, error) {
    res := []*Hand{}
    for _, h := range s {
        t := strings.Split(h, " ")
        score, err := strconv.Atoi(t[1])
        if err != nil {
            return nil, err
        }

        cards := []int{}
        for _, c := range t[0] {
            v, ok := m[c]
            if !ok {
                return nil, errors.New("Invalid rune found while parsing hand.")
            }
            cards = append(cards, v)
        }
        ty := cathegorize(cards, j)
        res = append(res, &Hand{In: t[0], Cards: cards, Score: score, Type: ty})
    }
    return res, nil
}

func cathegorize(c []int, j bool) int {
    temp := map[int]int{}
    if !j {
        for _, x := range c {
            _, ok := temp[x]
            if !ok {
                temp[x] = 1
            } else {
                temp[x] += 1
            }
        }
    } else {
        jkr := 0
        for _, v := range c {
            if v == 1 { jkr++; continue }
            _, ok := temp[v]
            if !ok {
                temp[v] = 1
            } else {
                temp[v] += 1
            }
        }

        // get max group
        r, m := 0, 0
        for k, v := range temp {
            if v > m { r, m = k, v }
        }

        temp[r] += jkr
    }
    switch len(temp) {
        case 1:
            return 7
        case 2:
            for _, v := range temp {
                if v == 1 || v == 4 { return 6 }
                if v == 2 || v == 3 { return 5 }
            }
        case 3:
            for _, v := range temp {
                if v == 3 { return 4 }
                if v == 2 {return 3 }
            }
        case 4:
            return 2
        case 5:
            return 1
    }
    return 0
}

