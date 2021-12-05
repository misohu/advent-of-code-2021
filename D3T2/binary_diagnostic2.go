package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary_diagnostic2: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	codes := []string{}
	candidates_ox := []int{}
	candidates_co := []int{}

	counter := 0
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
		candidates_ox = append(candidates_ox, counter)
		candidates_co = append(candidates_co, counter)
		counter++
	}

	targetBit := 0
	for {
		if len(candidates_ox) <= 1 || targetBit >= len(codes[0]) {
			break
		}
		candidates_ox = topOnTargetPos(codes, candidates_ox, targetBit, "ox")
		targetBit += 1
	}

	targetBit = 0
	for {
		if len(candidates_co) <= 1 || targetBit >= len(codes[0]) {
			break
		}
		candidates_co = topOnTargetPos(codes, candidates_co, targetBit, "co")
		targetBit += 1
	}

	ox, err := strconv.ParseInt(codes[candidates_ox[0]], 2, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary_diagnostic2: %v", err)
	}

	co, err := strconv.ParseInt(codes[candidates_co[0]], 2, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "binary_diagnostic2: %v", err)
	}
	fmt.Printf("ox: %s, co: %s\n", codes[candidates_ox[0]], codes[candidates_co[0]])
	fmt.Printf("ox: %v, co: %v\n", ox, co)
	fmt.Println(ox * co)
}

func topOnTargetPos(codes []string, candidates []int, target int, retType string) []int {
	counters := map[rune]int{
		'1': 0,
		'0': 0,
	}
	splits := map[rune][]int{
		'1': {},
		'0': {},
	}
	for _, candidate := range candidates {
		if codes[candidate][target] == '1' {
			counters['1']++
			splits['1'] = append(splits['1'], candidate)
		} else {
			counters['0']++
			splits['0'] = append(splits['0'], candidate)
		}
	}
	if retType == "ox" {
		if counters['1'] >= counters['0'] {
			return splits['1']
		} else {
			return splits['0']
		}
	} else if retType == "co" {
		if counters['0'] <= counters['1'] {
			return splits['0']
		} else {
			return splits['1']
		}
	} else {
		return []int{}
	}
}
