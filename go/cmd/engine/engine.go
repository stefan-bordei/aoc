package engine

import (
	"fmt"
	"log"

	"github.com/stefan-bordei/goAOC/cmd/utils"
	"github.com/stefan-bordei/goAOC/cmd/year22"
)


type fn func (string, string) (string, string, error)


func GetResults(y string, d string) (string, string, error) {
    m := map[string]fn {
        "22": mapper22.Solve,
    }

    puzzle_input, err := utils.GetInput(y, d)

    if err != nil {
        log.Fatal("Unable to get input data.")
    }

    p1, p2, err := m[y](d, puzzle_input)

    return p1, p2, err
}

func PrintResults(y string, d string) {
    fmt.Printf("\t----- AoC -----\nday: %s\nyear: %s\n\n", d, y)

    p1, p2, err := GetResults(y, d)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("\t~RESULT~\n\nPart1: %s\nPart2: %s\n", p1, p2)
}

