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
	dots = fold(dots, folds[0].ax, folds[0].pos)
	fmt.Println(len(dots))
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

func fold(dots map[Point]bool, axes, value int) map[Point]bool {
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
