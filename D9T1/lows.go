package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
	res := foundLowest(grid)
	fmt.Println(res)
}

func foundLowest(grid [][]int) (result int) {
	gridX := len(grid[0])
	gridY := len(grid)
	for y, line := range grid {
		for x, val := range line {
			next := false
			for _, n := range getNeighbours(x, y, gridX, gridY) {
				if val >= grid[n[1]][n[0]] {
					next = true
					break
				}
			}
			if !next {
				result += 1 + val
			}
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
