package day5

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

type Seed struct {
    Seed int
    Soil int
    Fertilizer int
    Water int
    Light int
    Temp int
    Humidity int
    Location int
}

func Solve(d string) (string, string, error) {
    s := utils.SplitByEmptyNewline(strings.TrimSpace(d))

    seeds := strings.Split(strings.Split(s[0], ": ")[1], " ")
    var sMap []*Seed
    for _, seed := range seeds {
        sTmp, err := strconv.Atoi(seed)
        if err != nil {
            return utils.NOT_DONE, utils.NOT_DONE, err
        }
        sMap = append(sMap, &Seed{Seed: sTmp,
                                Soil: -1,
                                Fertilizer: -1,
                                Water: -1,
                                Light: -1,
                                Temp: -1,
                                Humidity: -1,
                                Location: -1,
                            })
    }

    partOne, err := partOne(sMap, s[1:])
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partTwo, err := partTwo(sMap, s[1:])
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return partOne, partTwo, nil
}

func partOne(s []*Seed, d[]string) (string, error) {
    for move, elems := range d {
        t := utils.SplitLines(elems)[1:]
        for _, e := range t {
            elem := strings.Split(e, " ")
            rDest, err := strconv.Atoi(elem[0])
            if err != nil {
                return utils.NOT_DONE, err
            }

            rSource, err := strconv.Atoi(elem[1])
            if err != nil {
                return utils.NOT_DONE, err
            }

            rLen, err := strconv.Atoi(elem[2])
            if err != nil {
                return utils.NOT_DONE, err
            }

            for _, m := range s {
                updateMapping(rSource, rDest, rLen, m, move)
            }
        }
        for _, m := range s {
            finishMapping(m, move)
        }
    }

    res := s[0].Location
    for i := 1; i < len(s); i++ {
        if s[i].Location < res {
            res = s[i].Location
        }
    }

    return strconv.Itoa(res), nil
}

func finishMapping(o *Seed, check int) {
    switch check {
    case 0:
        if o.Soil == -1 {
            o.Soil = o.Seed
        }
    case 1:
        if o.Fertilizer == -1 {
            o.Fertilizer = o.Soil
        }
    case 2:
        if o.Water == -1 {
            o.Water = o.Fertilizer
        }
    case 3:
        if o.Light == -1 {
            o.Light = o.Water
        }
    case 4:
        if o.Temp == -1 {
            o.Temp = o.Light
        }
    case 5:
        if o.Humidity == -1 {
            o.Humidity = o.Temp
        }
    case 6:
        if o.Location == -1 {
            o.Location = o.Humidity
        }
    }

}

func partTwo(s []*Seed, d []string) (string, error) {
    return utils.NOT_DONE, nil
}

func updateMapping(s, d, l int, o *Seed, check int) {
    switch check {
    case 0:
        if checkInBounds(s, d, l, o.Seed) {
            o.Soil = updateDif(s, d, o.Seed)
        }
    case 1:
        if checkInBounds(s, d, l, o.Soil) {
            o.Fertilizer = updateDif(s, d, o.Soil)
        }
    case 2:
        if checkInBounds(s, d, l, o.Fertilizer) {
            o.Water = updateDif(s, d, o.Fertilizer)
        }
    case 3:
        if checkInBounds(s, d, l, o.Water) {
            o.Light = updateDif(s, d, o.Water)
        }
    case 4:
        if checkInBounds(s, d, l, o.Light) {
            o.Temp = updateDif(s, d, o.Light)
        }
    case 5:
        if checkInBounds(s, d, l, o.Temp) {
            o.Humidity = updateDif(s, d, o.Temp)
        }
    case 6:
        if checkInBounds(s, d, l, o.Humidity) {
            o.Location = updateDif(s, d, o.Humidity)
        }
    }
}

func updateDif(s, d, o int) int {
    return abs(o - s + d)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func checkInBounds(s, d, l, o int) bool {
    if s <= o && o <= s + l {
        return true
    }
    return false
}
