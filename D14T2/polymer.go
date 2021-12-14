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
	start, counts := readFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dots: problem with parsing file %v", err)
		os.Exit(1)
	}

	for i := 0; i < 40; i++ {
		start = step(start, counts)
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
	key string
	val int
}

func step(start map[string]int, counts map[string]int) (result map[string]int) {
	result = make(map[string]int)
	for k, v := range start {
		tmp := recipe[k]
		counts[tmp] += v
		result[string(k[0])+tmp] += v
		result[tmp+string(k[1])] += v
	}
	return
}

func readFile(f *os.File) (start map[string]int, counts map[string]int) {
	start = make(map[string]int)
	counts = make(map[string]int)
	input := bufio.NewScanner(f)
	input.Scan()
	tmp := input.Text()
	for i := 0; i < len(tmp)-1; i++ {
		counts[string(tmp[i])]++
		start[tmp[i:i+2]]++
	}
	counts[string(tmp[len(tmp)-1])]++
	input.Scan()
	for input.Scan() {
		parts := strings.Split(input.Text(), " -> ")
		recipe[parts[0]] = parts[1]
	}
	return
}
