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

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labirint: problem with file %v", err)
		os.Exit(1)
	}
	lab, _, end, err := readFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labirint: problem with parsing file %v", err)
		os.Exit(1)
	}
	res := dijkstra(lab, Point{Coord{0, 0}, 0}, end)
	fmt.Println(res)
}

func dijkstra(l Labirint, start, end Point) int {
	visited := []Coord{start.Coord}
	var queue []Point
	queue = append(queue, start)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.Coord == end.Coord {
			return cur.val
		}
		for _, n := range l[cur.Coord] {
			if !contains(visited, n.Coord) {
				queue = append(queue, Point{n.Coord, cur.val + n.val})
				visited = append(visited, n.Coord)
			}
		}
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].val < queue[j].val
		})
	}
	return 0
}

func contains(s []Coord, e Coord) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readFile(f *os.File) (l Labirint, vals map[Coord]int, end Point, err error) {
	input := bufio.NewScanner(f)
	tmp := [][]int{}
	for input.Scan() {
		line := strings.Split(input.Text(), "")
		row := []int{}
		for _, num := range line {
			val, err := strconv.Atoi(num)
			if err != nil {
				e := fmt.Sprintf("labirint: problem with parsing file %v", err)
				return nil, nil, Point{}, fmt.Errorf(e)
			}
			row = append(row, val)
		}
		tmp = append(tmp, row)
	}
	l = make(Labirint)
	vals = make(map[Coord]int)
	for x, row := range tmp {
		for y, val := range row {
			l[Coord{x, y}] = []Point{}
			neigh := findNeighbours(Point{Coord{x, y}, val}, len(tmp)-1, len(row)-1)
			for _, n := range neigh {
				l[Coord{x, y}] = append(l[Coord{x, y}], Point{n, tmp[n.x][n.y]})
			}
			vals[Coord{x, y}] = val
		}
	}
	return l, vals, Point{Coord{len(tmp) - 1, len(tmp[0]) - 1}, tmp[len(tmp)-1][len(tmp[0])-1]}, nil
}

func findNeighbours(p Point, maX, maxY int) []Coord {
	var neighbours []Coord
	if p.x > 0 {
		neighbours = append(neighbours, Coord{p.x - 1, p.y})
	}
	if p.x < maX {
		neighbours = append(neighbours, Coord{p.x + 1, p.y})
	}
	if p.y > 0 {
		neighbours = append(neighbours, Coord{p.x, p.y - 1})
	}
	if p.y < maxY {
		neighbours = append(neighbours, Coord{p.x, p.y + 1})
	}
	return neighbours
}
