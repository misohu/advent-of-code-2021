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
	field := make(map[Point]int)
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "hydrotermal: %v", err)
		os.Exit(1)
	}

	coords := readFile(f)
	for _, line := range coords {
		isTarget, d := line.isHorOrVert()
		if isTarget {
			for _, point := range line.pointsBetween(d) {
				field[point]++
			}
		}
	}

	counter := 0
	for _, v := range field {
		if v >= 2 {
			counter++
		}
	}
	fmt.Println(counter)
	f.Close()
}

type Point struct {
	x int
	y int
}

type Line struct {
	start Point
	end   Point
}

func (l *Line) isHorOrVert() (bool, string) {
	if l.start.x == l.end.x {
		return true, "y"
	} else if l.start.y == l.end.y {
		return true, "x"
	} else {
		return false, ""
	}
}

func (l *Line) pointsBetween(direction string) []Point {
	points := []Point{}
	var start, finish, freeze int
	if direction == "x" {
		freeze = l.start.y
		start = l.start.x
		finish = l.end.x
	} else if direction == "y" {
		freeze = l.start.x
		start = l.start.y
		finish = l.end.y
	}
	for start != finish {
		var newPoint Point
		if direction == "x" {
			newPoint = Point{
				x: start,
				y: freeze,
			}
		} else if direction == "y" {
			newPoint = Point{
				x: freeze,
				y: start,
			}
		}
		points = append(points, newPoint)
		if start < finish {
			start += 1
		} else {
			start -= 1
		}
	}
	points = append(points, l.end)
	return points
}

func readFile(f *os.File) []Line {
	input := bufio.NewScanner(f)
	lines := []Line{}
	for input.Scan() {
		coords := strings.Split(input.Text(), " -> ")
		lines = append(lines, Line{
			start: parseCoordinates(coords[0]),
			end:   parseCoordinates(coords[1]),
		})
	}
	return lines
}

func parseCoordinates(s string) Point {
	coords := strings.Split(s, ",")
	x, err := strconv.Atoi(coords[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parseCoordinates: x parsing %v\n", err)
		os.Exit(1)
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "parseCoordinates: y parsing %v\n", err)
		os.Exit(1)
	}
	return Point{x, y}
}
