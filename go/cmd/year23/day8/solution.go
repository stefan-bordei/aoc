package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

var Instrction = map[rune]int{
    'R': 1,
    'L': 0,
}

func Solve(d string) (string, string, error) {
    s := utils.SplitLines(strings.TrimSpace(d))

    instructions := s[0]
    temp := s[2:]
    nodes := map[string][]string{}
    for _, n := range temp {
        t := strings.Split(n, " = ")
        v := strings.Split(t[1][1:len(t[1])-1], ", ")
        nodes[t[0]] = v
    }
    partOne, err := partOne(instructions, nodes)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partTwo, err := partTwo(s)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return partOne, partTwo, nil
}

func partOne(i string, n map[string][]string) (string, error) {
    res := 0
    fmt.Printf("I: %s\n\nNodes: %s\n", i, n)
    start := "AAA"

    for {
        if start == "ZZZ" { break }
        for _, c := range i {
            res++
            temp := n[start][Instrction[c]]
            start = string(temp)
        }
    }
    return strconv.Itoa(res), nil
}

func partTwo(d []string) (string, error) {
    res := 0
    return strconv.Itoa(res), nil
}
