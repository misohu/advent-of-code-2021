package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lows: %v", f)
		os.Exit(1)
	}
	grid, err := readFile(f)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "lows: %v", f)
		os.Exit(1)
	}
	basins := make(map[Point]int)
	for x, line := range grid {
		for y := range line {
			visited := [][]bool{}
			for range grid {
				visited = append(visited, make([]bool, len(grid[0])))
			}
			if grid[x][y] != 9 {
				basins[findBasin(x, y, grid, visited)]++
			}
		}
	}
	p := make(PairList, len(basins))
	i := 0
	for k, v := range basins {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	result := 1
	for _, k := range p[:3] {
		fmt.Printf("%v\t%v -> %v\n", k.Key, k.Value, grid[k.Key.x][k.Key.y])
		result *= k.Value
	}
	fmt.Println(result)
}

type Pair struct {
	Key   Point
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func findBasin(x, y int, grid [][]int, visited [][]bool) Point {
	visited[x][y] = true
	if isLowest(x, y, grid) {
		return Point{x, y}
	}
	for _, n := range findLower(x, y, grid) {
		if !visited[n[0]][n[1]] {
			return findBasin(n[0], n[1], grid, visited)
		}
	}
	return Point{}
}

func findLower(x, y int, grid [][]int) (result [][]int) {
	val := grid[x][y]
	gridX := len(grid)
	gridY := len(grid[0])
	for _, n := range getNeighbours(x, y, gridX, gridY) {
		if val >= grid[n[0]][n[1]] {
			result = append(result, []int{n[0], n[1]})
		}
	}
	return
}

func isLowest(x, y int, grid [][]int) (isLowest bool) {
	val := grid[x][y]
	gridX := len(grid)
	gridY := len(grid[0])
	isLowest = true
	for _, n := range getNeighbours(x, y, gridX, gridY) {
		if val >= grid[n[0]][n[1]] {
			isLowest = false
			break
		}
	}
	return
}

func getNeighbours(x, y, gridX, gridY int) (neigh [][]int) {
	// up
	if x > 0 {
		neigh = append(neigh, []int{x - 1, y})
	}
	// down
	if x < gridX-1 {
		neigh = append(neigh, []int{x + 1, y})
	}
	// left
	if y > 0 {
		neigh = append(neigh, []int{x, y - 1})
	}
	// right
	if y < gridY-1 {
		neigh = append(neigh, []int{x, y + 1})
	}
	return
}

func readFile(f *os.File) (result [][]int, err error) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := []int{}
		for _, c := range strings.Split(scanner.Text(), "") {
			num, err := strconv.Atoi(c)
			if err != nil {
				return nil, fmt.Errorf("readFile: parsing err %v", err)
			}
			line = append(line, num)
		}
		result = append(result, line)
	}
	return
}
