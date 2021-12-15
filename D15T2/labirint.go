package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

type Point struct {
	Coord
	val int
}

type Labirint map[Coord][]Point

var multiplier = 5

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labirint: problem with file %v", err)
		os.Exit(1)
	}
	vals, dim, err := readFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labirint: problem with parsing file %v", err)
		os.Exit(1)
	}
	res := dijkstra(vals, dim*multiplier, Point{Coord{0, 0}, 0}, Point{Coord{dim*multiplier - 1, dim*multiplier - 1}, 0})
	fmt.Println(res)
}

func dijkstra(vals map[Coord]int, dim int, start, end Point) int {
	visited := []Coord{start.Coord}
	var queue []Point
	queue = append(queue, start)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Coord == end.Coord {
			return cur.val
		}
		for _, n := range findNeighbours(cur, dim, dim) {
			if !contains(visited, n) {
				queue = append(queue, Point{n, cur.val + realValue(n, dim/multiplier, dim/multiplier, vals)})
				visited = append(visited, n)
			}
		}
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].val < queue[j].val
		})
	}
	return 0
}

func realValue(c Coord, labX, labY int, vals map[Coord]int) int {
	gridX, gridY := c.x/labX, c.y/labY
	realX, realY := c.x%labX, c.y%labY

	newVal := (gridX + gridY) + vals[Coord{realX, realY}]

	if newVal > 9 {
		return (newVal - 9)
	} else {
		return newVal
	}
}

func contains(s []Coord, e Coord) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readFile(f *os.File) (vals map[Coord]int, dim int, err error) {
	input := bufio.NewScanner(f)
	vals = make(map[Coord]int)
	x := 0
	for input.Scan() {
		line := strings.Split(input.Text(), "")
		for y, num := range line {
			val, err := strconv.Atoi(num)
			if err != nil {
				e := fmt.Sprintf("labirint: problem with parsing file %v", err)
				return nil, 0, fmt.Errorf(e)
			}
			vals[Coord{x, y}] = val
		}
		x++
	}
	return vals, x, err
}

func findNeighbours(p Point, maX, maxY int) []Coord {
	var neighbours []Coord
	if p.x > 0 {
		neighbours = append(neighbours, Coord{p.x - 1, p.y})
	}
	if p.x < maX-1 {
		neighbours = append(neighbours, Coord{p.x + 1, p.y})
	}
	if p.y > 0 {
		neighbours = append(neighbours, Coord{p.x, p.y - 1})
	}
	if p.y < maxY-1 {
		neighbours = append(neighbours, Coord{p.x, p.y + 1})
	}
	return neighbours
}
