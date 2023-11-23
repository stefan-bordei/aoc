package main

import (
	"os"

	"github.com/stefan-bordei/goAOC/cmd/engine"
)

func main() {
    allArgs := os.Args[1:]
    var day, year string;

    /* Be able to provide only one  argument to make run or no arguments
    *       Valid:
    *           make run
    *           make run day=1
    *           make run day=1 year=22
    *
    *   If no args are provided it defaults ro day=1 year=22
    */
    switch len(allArgs) {
    case 0:
        day, year = "1", "22"
    case 1:
        day, year = allArgs[0], "22"
    default:
        day, year = allArgs[0], allArgs[1]
    }

    engine.PrintResults(year, day)
}
