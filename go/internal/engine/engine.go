package engine

import (
	"log"

	"github.com/stefan-bordei/goAOC/cmd/year22/day1"
	"github.com/stefan-bordei/goAOC/cmd/year22/day2"
	"github.com/stefan-bordei/goAOC/cmd/year22/day3"
	"github.com/stefan-bordei/goAOC/cmd/year22/day4"
	"github.com/stefan-bordei/goAOC/cmd/year22/day5"
	"github.com/stefan-bordei/goAOC/internal/utils"
)

func Solve(y string, d string) (string, string, error) {
    puzzle_input, err := utils.GetInput(y, d)

    if err != nil {
        log.Fatal("Unable to get input data.")
    }

    switch y {
    case "22":
        switch d {
        case "1":
            part_one, part_two, err := day1.Solve(puzzle_input)
            return part_one, part_two, err

        case "2":
            part_one, part_two, err := day2.Solve(puzzle_input)
            return part_one, part_two, err

        case "3":
            part_one, part_two, err := day3.Solve(puzzle_input)
            return part_one, part_two, err

        case "4":
            part_one, part_two, err := day4.Solve(puzzle_input)
            return part_one, part_two, err

        case "5":
            part_one, part_two, err := day5.Solve(puzzle_input)
            return part_one, part_two, err
        }
    }

    return utils.NOT_DONE, utils.NOT_DONE, nil
}

