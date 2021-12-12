package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point [2]int

type Cave struct {
	name   string
	next   []string
	isBig  bool
	isEdge bool
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lows: %v", f)
		os.Exit(1)
	}
	caves := readFile(f)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "lows: %v", f)
		os.Exit(1)
	}
	findEnd("", "start", caves, []string{"start"}, true)
	fmt.Println(counter)
}

var counter = 0

func findEnd(path string, cave string, caves map[string]*Cave, visited []string, oneSmall bool) {
	if cave == "end" {
		// fmt.Println(path + "," + "end")
		counter++
		return
	}
	for _, c := range (*caves[cave]).next {
		if !caves[c].isBig {
			if isIn, _ := isInSlice(c, visited); !isIn {
				findEnd(path+","+cave, c, caves, append(visited, c), oneSmall)
			} else if oneSmall && !caves[c].isEdge {
				findEnd(path+","+cave, c, caves, visited, false)
			}
		} else {
			findEnd(path+","+cave, c, caves, visited, oneSmall)
		}
	}
}

func isInSlice(target string, slice []string) (bool, int) {
	for i, s := range slice {
		if s == target {
			return true, i
		}
	}
	return false, 0
}

func readFile(f *os.File) map[string]*Cave {
	caves := make(map[string]*Cave)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		from, to := line[0], line[1]
		parseCave(from, to, caves)
		parseCave(to, from, caves)
	}
	return caves
}

func parseCave(cave string, to string, caves map[string]*Cave) {
	if target, ok := caves[cave]; !ok {
		caves[cave] = &Cave{
			name:   cave,
			next:   []string{to},
			isBig:  strings.ToUpper(cave) == cave,
			isEdge: cave == "start",
		}
	} else {
		target.next = append(target.next, to)
	}

}
