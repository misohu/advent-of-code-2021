package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var opens = []string{"(", "[", "{", "<"}
var closes = []string{")", "]", "}", ">"}
var pairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}
var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func contains(s []string, target string) bool {
	for _, e := range s {
		if target == e {
			return true
		}
	}
	return false
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "lows: %v", f)
		os.Exit(1)
	}
	ground, err := readFile(f)
	defer f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "lows: %v", f)
		os.Exit(1)
	}
	score := 0
	for _, line := range ground {
		prob, ok := parseLine(line)
		if !ok && prob != "inc" {
			score += scores[prob]
		}
	}
	fmt.Println(score)
}

func parseLine(line []string) (string, bool) {
	mem := []string{}
	for _, s := range line {
		if len(opens) == 0 || contains(opens, s) {
			mem = append(mem, s)
			continue
		}
		// problem when starts with closing
		if contains(closes, s) && pairs[mem[len(mem)-1]] == s {
			mem = mem[:len(mem)-1]
			continue
		} else if contains(closes, s) {
			return s, false
		}
	}
	if len(mem) > 0 {
		return "inc", false
	}
	return "", true
}

func readFile(f *os.File) (result [][]string, err error) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}
	return
}
