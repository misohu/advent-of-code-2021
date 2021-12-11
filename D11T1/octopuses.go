package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point [2]int

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

	counter := 0
	for i := 0; i < 100; i++ {
		isTen := make(map[[2]int]bool)
		for x, line := range grid {
			for y, _ := range line {
				if isTen[[2]int{x, y}] {
					continue
				}
				grid[x][y] += 1
				if grid[x][y] == 10 {
					grid[x][y] = 0
					isTen[[2]int{x, y}] = true
					neighbours := getNeighbours(x, y, len(grid), len(grid[0]))
					var incNeighbours func(nei [][]int)
					incNeighbours = func(nei [][]int) {
						for _, n := range nei {
							if isTen[[2]int{n[0], n[1]}] {
								continue
							}
							grid[n[0]][n[1]]++
							if grid[n[0]][n[1]] == 10 {
								grid[n[0]][n[1]] = 0
								isTen[[2]int{n[0], n[1]}] = true
								temp := getNeighbours(n[0], n[1], len(grid), len(grid[0]))
								incNeighbours(temp)
							}
						}
					}
					incNeighbours(neighbours)
				}
			}
		}
		for _ = range isTen {
			counter++
		}
	}

	fmt.Println(counter)
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

	// upleft
	if x > 0 && y > 0 {
		neigh = append(neigh, []int{x - 1, y - 1})
	}
	// downleft
	if x < gridX-1 && y > 0 {
		neigh = append(neigh, []int{x + 1, y - 1})
	}
	// downright
	if x < gridX-1 && y < gridY-1 {
		neigh = append(neigh, []int{x + 1, y + 1})
	}
	// top right
	if x > 0 && y < gridY-1 {
		neigh = append(neigh, []int{x - 1, y + 1})
	}
	return
}
