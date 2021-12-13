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
		fmt.Fprintf(os.Stderr, "dots: problem with file %v", err)
		os.Exit(1)
	}
	dots, folds, err := readFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dots: problem with parsing file %v", err)
		os.Exit(1)
	}
	for _, fold := range folds {
		dots = foldPaper(dots, fold.ax, fold.pos)
	}
	maxX := 0
	maxY := 0
	for dot, _ := range dots {
		if dot.x > maxX {
			maxX = dot.x
		}
		if dot.y > maxY {
			maxY = dot.y
		}
	}

	grid := [][]int{}
	for i := 0; i <= maxY; i++ {
		grid = append(grid, make([]int, maxX+1))
	}
	for dot, _ := range dots {
		grid[dot.y][dot.x] = 1
	}
	for _, line := range grid {
		for _, val := range line {
			if val == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type Point struct {
	x, y int
}

type Fold struct {
	ax  int
	pos int
}

func readFile(f *os.File) (dots map[Point]bool, folds []Fold, err error) {
	dots = make(map[Point]bool)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break
		}
		newCoors := []int{}
		lineNums := strings.Split(line, ",")
		for _, lineNum := range lineNums {
			num, err := strconv.Atoi(lineNum)
			if err != nil {
				error := fmt.Errorf("readFile: parsing coord %v", err)
				return nil, nil, error
			}
			newCoors = append(newCoors, num)
		}
		dots[Point{newCoors[0], newCoors[1]}] = true
	}
	for input.Scan() {
		line := input.Text()
		segments := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
		value, err := strconv.Atoi(segments[1])
		if err != nil {
			error := fmt.Errorf("readFile: parsing folds %v", err)
			return nil, nil, error
		}
		var ax int
		if segments[0] == "y" {
			ax = 0
		} else {
			ax = 1
		}
		folds = append(folds, Fold{ax: ax, pos: value})
	}
	return
}

func foldPaper(dots map[Point]bool, axes, value int) map[Point]bool {
	for k, _ := range dots {
		if axes == 0 {
			if k.y > value {
				dots[Point{k.x, value - (k.y - value)}] = true
				delete(dots, k)
			}
		} else {
			if k.x > value {
				dots[Point{value - (k.x - value), k.y}] = true
				delete(dots, k)
			}
		}
	}
	return dots
}
