package day10

import (
	"strconv"
	"strings"

	"github.com/stefan-bordei/aoc/go/cmd/utils"
)

type Queue struct {
    q []*Point
}

func (q *Queue) len() int {
    return len(q.q)
}

func (p *Queue) Enqueue(x *Point) {
    p.q = append(p.q, x)
}

func (p *Queue) Dequeue() *Point {
    if len(p.q) == 0 {
        return &Point{}
    }
    e := p.q[0]
    if len(p.q) > 1 {
        p.q = p.q[1:]
    } else {
        p.q = []*Point{}
    }
    return e
}

var north = "|7FS"
var south = "|LJS"
var east = "J-7S"
var west = "L-FS"

type Grid struct {
    Contents [][]*Point
    Ground []*Point
    Visited map[*Point]bool
}

func (g *Grid) rows() int {
    return len(g.Contents)-1
}

func (g *Grid) cols() int {
    // Assuming all rows have the same number of cols
    return len(g.Contents[0])-1
}

type Point struct {
    Id rune
    X int
    Y int
}

func (p *Point) Scan(g *Grid, q *Queue, s *[]rune) {
    if p.X != 0 && strings.ContainsRune(north, g.Contents[p.X-1][p.Y].Id) && strings.ContainsRune(south, p.Id) {
        if _, ok := g.Visited[g.Contents[p.X-1][p.Y]]; !ok {
            g.Visited[g.Contents[p.X-1][p.Y]] = true
            q.Enqueue(g.Contents[p.X-1][p.Y])
            if p.Id == 'S' {
               intersect(s, south)
            }
        }
    }
    if p.X != g.rows() && strings.ContainsRune(south, g.Contents[p.X+1][p.Y].Id) && strings.ContainsRune(north, p.Id) {
        if _, ok := g.Visited[g.Contents[p.X+1][p.Y]]; !ok {
            g.Visited[g.Contents[p.X+1][p.Y]] = true
            q.Enqueue(g.Contents[p.X+1][p.Y])
            if p.Id == 'S' {
               intersect(s, north)
            }
        }
    }
    if p.Y != g.cols() && strings.ContainsRune(east, g.Contents[p.X][p.Y+1].Id) && strings.ContainsRune(west, p.Id) {
        if _, ok := g.Visited[g.Contents[p.X][p.Y+1]]; !ok {
            g.Visited[g.Contents[p.X][p.Y+1]] = true
            q.Enqueue(g.Contents[p.X][p.Y+1])
            if p.Id == 'S' {
               intersect(s, west)
            }
        }
    }
    if p.Y != 0 && strings.ContainsRune(west, g.Contents[p.X][p.Y-1].Id) && strings.ContainsRune(east, p.Id) {
        if _, ok := g.Visited[g.Contents[p.X][p.Y-1]]; !ok {
            g.Visited[g.Contents[p.X][p.Y-1]] = true
            q.Enqueue(g.Contents[p.X][p.Y-1])
            if p.Id == 'S' {
               intersect(s, east)
            }
        }
    }
}

func Solve(d string) (string, string, error) {
    s := utils.SplitLines(strings.TrimSpace(d))
    sp := startingPoint(s)
    ss := []rune{'|', '-', '7', 'F', 'J', 'L'}
    g := Grid{Visited: map[*Point]bool{}}
    for r, rc := range s {
        m := []*Point{}
        for c, cc := range rc {
            m = append(m, &Point{Id: cc, X: r, Y: c})
        }
        g.Contents = append(g.Contents, m)
    }

    q := Queue{}
    q.Enqueue(sp)
    g.Visited[sp] = true

    for q.len() > 0 {
        tp := q.Dequeue()
        tp.Scan(&g, &q, &ss)
    }

    // Update 'S' to actual value
    for _, r := range g.Contents {
        for _, c := range r {
            if c.Id == 'S' {
                c.Id = ss[0]
            }
        }
    }

    partOne, err := partOne(&g)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    partTwo, err := partTwo(&g)
    if err != nil {
        return utils.NOT_DONE, utils.NOT_DONE, err
    }

    return partOne, partTwo, nil
}

func partOne(g *Grid) (string, error) {
    res := len(g.Visited) / 2

    return strconv.Itoa(res), nil
}

func partTwo(g *Grid) (string, error) {
    res := 0
    ng := ReplaceOutside(g)

    for _, gr := range ng.Ground {
        count := 0
        for i := gr.Y; i >= 0; i-- {
            if gr.Y == len(ng.Contents[0])-1 {
                continue
            }
            if ng.Contents[gr.X][i].Id == '|' {
                count++
            }
            if ng.Contents[gr.X][i].Id == 'J' {
                count++
            }
            if ng.Contents[gr.X][i].Id == 'L' {
                count++
            }
        }
        if count % 2 != 0 && count > 0{
            res ++
        }
    }
    return strconv.Itoa(res), nil
}

func intersect(r *[]rune, l string) {
    ts := []rune{}
    for _, rc := range *r {
        for _, lc := range l {
            if rc == lc {
                ts = append(ts, rc)
            }
        }
    }
    *r = ts
}

func ReplaceOutside(g *Grid) Grid {
    rg := Grid{}
    for r, row := range g.Contents {
        m := []*Point{}
        for c, ch := range row {
            if _, ok := g.Visited[ch]; !ok {
                rg.Ground = append(rg.Ground, ch)
                m = append(m, &Point{Id: '.', X: r, Y: c})
            } else {
                m = append(m, ch)
            }
        }
        rg.Contents = append(rg.Contents, m)
    }
    return rg
}

func startingPoint(d []string) *Point {
    for r, rowch := range d {
        for c, colch := range rowch {
            if colch == 'S' {
                return &Point{Id:'S', X: r, Y: c}
            }
        }
    }
    return &Point{}
}
