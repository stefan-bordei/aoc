package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const NOT_DONE = "In Progress..."

func readFileInput(p string) (string, error) {
    filePath, _ := filepath.Abs(p)
    contents, err := os.ReadFile(filePath)
    return string(contents), err
}

func GetInput(y string, d string) (string, error) {
    filePath := fmt.Sprintf("./data/%s/day%s", y, d)
    contents, err := readFileInput(filePath)
    return contents, err
}

func SumUp(s []int) int {
    sum := 0
    for _, e := range s {
        sum += e
    }
    return sum
}

func SplitByEmptyNewline(s string) []string {
    //res := regexp.MustCompile(`\n\s*\n`).Split(s, -1)
    res := strings.Split(s, "\n\n")
    return res

}

func SplitLines(s string) []string {
    res := strings.Split(s, "\n")
    return res
}
