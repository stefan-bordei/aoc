package day08

import (
	"errors"
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
	partOne, _ := partOne(instructions, nodes, "AAA", "ZZZ")

	partTwo, _ := partTwo(instructions, nodes, "A", "Z")

	return partOne, partTwo, nil
}

func partOne(i string, n map[string][]string, start, stop string) (string, error) {
	res := 0
	s := start

	for {
		if s == stop {
			break
		}
		for _, c := range i {
			res++
			temp, ok := n[s]
			if !ok {
				return utils.NOT_DONE, errors.New("Node not found.")
			}
			s = string(temp[Instrction[c]])
		}
	}
	return strconv.Itoa(res), nil
}

func partTwo(i string, n map[string][]string, start, stop string) (string, error) {
	res := 0
	startingPoints := []string{}
	// get all the starting nodes
	for k, _ := range n {
		if strings.Contains(k, start) {
			startingPoints = append(startingPoints, k)
		}
	}

	cycles := []int{}
	for _, s := range startingPoints {
		tmp := 0
		for {
			if strings.Contains(s, stop) {
				break
			}
			for _, c := range i {
				tmp++
				temp, ok := n[s]
				if !ok {
					return utils.NOT_DONE, errors.New("Node not found.")
				}
				s = string(temp[Instrction[c]])
			}
		}
		cycles = append(cycles, tmp)
	}
	res = lcm(cycles...)

	return strconv.Itoa(res), nil
}

func lcm(integers ...int) int {
	if len(integers) == 1 {
		return integers[0]
	}
	x := integers[0]
	y := integers[1]
	result := x * y / gcd(x, y)

	if len(integers) == 2 {
		return result
	}
	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}
	return result
}

func gcd(x, y int) int {
	for y != 0 {
		temp := y
		y = x % y
		x = temp
	}
	return x
}
