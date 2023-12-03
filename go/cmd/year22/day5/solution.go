package day5

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

func Solve(d string) (string, string, error) {
    s := utils.SplitByEmptyNewline(d)

    instructions := utils.SplitLines(strings.TrimSpace(s[1]))
    part_one, err := partOne(s[0], instructions)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    part_two, err := partTwo(s[0], instructions)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return part_one, part_two, nil
}

func partOne(stacks string, d []string) (string, error) {
    c, s := parseCrates(stacks)

    for _, i := range d {
        raw := strings.Split(i, " ")

        move, err := strconv.Atoi(raw[1])
        if err != nil {
            return utils.NOT_DONE, err
        }

        from, to := raw[3], raw[5]

        for p := 0; p < move; p++ {
            tmp := c[from][len(c[from])-1]
            c[to] = append(c[to], tmp)
            c[from] = c[from][:len(c[from])-1]
        }
    }

    res := ""
    for _, i := range s {
        if i > '0' && i <= '9' {
            res += string(c[string(i)][len(c[string(i)])-1])
        }
    }

    return res, nil
}

func partTwo(stacks string, d []string) (string, error) {
    c, s := parseCrates(stacks)

    for _, i := range d {
        raw := strings.Split(i, " ")

        move, err := strconv.Atoi(raw[1])
        if err != nil {
            return utils.NOT_DONE, err
        }

        from, to := raw[3], raw[5]

        breakpoint := len(c[from])-move
        c[to] = append(c[to], c[from][breakpoint:]...)
        c[from] = c[from][:breakpoint]
    }

    res := ""
    for _, i := range s {
        if i > '0' && i <= '9' {
            res += string(c[string(i)][len(c[string(i)])-1])
        }
    }
    return res, nil
}

func parseCrates(s string) (map[string][]byte, string) {
    temp := utils.SplitLines(s)
    stacks := temp[:len(temp)-1]
    stack_num := temp[len(temp)-1]

    m := make(map[string][]byte)

    for pos := 0; pos <= len(stack_num)-1; pos++ {
        if stack_num[pos] >= '0' && stack_num[pos] <= '9' {
            for j := len(stacks)-1; j >= 0; j-- {
                if pos <= len(stacks[j])-1 && stacks[j][pos] >= 'A' && stacks[j][pos] <= 'Z' {
                    m[string(stack_num[pos])] = append(m[string(stack_num[pos])], stacks[j][pos])
                }
            }
        }
    }
    return m, stack_num
}
