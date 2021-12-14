package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var recipe map[string]string = make(map[string]string)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dots: problem with file %v", err)
		os.Exit(1)
	}
	start := readFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dots: problem with parsing file %v", err)
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		start = step(start)
		// fmt.Println(start)
	}
	counts := make(map[rune]int)
	for _, c := range start {
		counts[c]++
	}
	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	fmt.Println(ss[0].val - ss[len(ss)-1].val)
}

type kv struct {
	key rune
	val int
}

func step(start string) (result string) {
	result = ""
	for i := 0; i < len(start)-1; i++ {
		result += recipe[start[i:i+2]]
	}
	result += string(start[len(start)-1])
	return
}

func readFile(f *os.File) (start string) {
	input := bufio.NewScanner(f)
	input.Scan()
	start = input.Text()
	input.Scan()
	for input.Scan() {
		parts := strings.Split(input.Text(), " -> ")
		recipe[parts[0]] = string(parts[0][0]) + parts[1]
	}
	return
}
