package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
var scoresIncom = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
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
	scores := []int{}
	for _, line := range ground {
		mem, is := isIncomplete(line)
		if is {
			score := 0
			for i := len(mem) - 1; i >= 0; i-- {
				score *= 5
				score += scoresIncom[pairs[mem[i]]]
			}
			scores = append(scores, score)
		}
	}
	s := Scores(scores)
	sort.Sort(s)
	fmt.Println(scores[len(scores)/2])
}

type Scores []int

func (s Scores) Less(i, j int) bool { return s[i] < s[j] }
func (s Scores) Len() int           { return len(s) }
func (s Scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func isIncomplete(line []string) ([]string, bool) {
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
			return nil, false
		}
	}
	if len(mem) > 0 {
		return mem, true
	}
	return nil, false
}

func readFile(f *os.File) (result [][]string, err error) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, strings.Split(scanner.Text(), ""))
	}
	return
}
