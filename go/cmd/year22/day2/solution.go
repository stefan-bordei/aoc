package day2

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/goAOC/cmd/utils"
)

const (
    Win = 6
    Draw = 3
    Loss = 0

    Rock = 1
    Paper = 2
    Scissors = 3
)

var choice map[string]int = map[string]int {
    "X": Rock,
    "Y": Paper,
    "Z": Scissors,

    "A": Rock,
    "B": Paper,
    "C": Scissors,
}

var win map[string]string = map[string]string {
    "X": "C",
    "Y": "A",
    "Z": "B",
}

var lose map[string]string = map[string]string {
    "X": "B",
    "Y": "C",
    "Z": "A",
}

var secondStrategy map[string]int = map[string]int {
    "X": Loss,
    "Y": Draw,
    "Z": Win,
}

type choiceError struct {
    s string
}

func (e *choiceError) Error() string {
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
    score := 0

    for _, i := range d {
        s := strings.Fields(i)
        l, r := s[0], s[1]

        _, ok := choice[r]
        if !ok {
            return "0", &choiceError{"Unsupported player choice."}
        }

        score += choice[r]
        switch {
        case choice[l] == choice[r]:
            score += Draw
        case win[r] == l:
            score += Win
        default:
            score += Loss
        }
    }

    return strconv.Itoa(score), nil
}

func partTwo(d []string) (string, error) {
    score := 0
    var w, f map[string]string = reverseMap(lose), reverseMap(win)

    for _, i := range d {
        s := strings.Fields(i)
        l, r := s[0], s[1]

        _, ok := choice[r]
        if !ok {
            return "0", &choiceError{"Unsupported player choice."}
        }

        score += secondStrategy[r]
        switch secondStrategy[r] {
        case 0:
            score += choice[w[l]]
        case 6:
            score += choice[f[l]]
        default:
            score += choice[l]
        }
    }

    return strconv.Itoa(score), nil
}

func reverseMap(m map[string]string) map[string]string {
    n := make(map[string]string, len(m))

    for k, v := range m {
        n[v] = k
    }

    return n
}
