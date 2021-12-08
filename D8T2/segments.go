package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var lengths = map[int]int{
	0: 6,
	1: 2,
	2: 5,
	3: 5,
	4: 4,
	5: 5,
	6: 6,
	7: 3,
	8: 7,
	9: 6,
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "segments: wrong input file %v", err)
		os.Exit(1)
	}
	inputs := readFile(f)

	s := 0
	for _, input := range inputs {
		encoding := findSetup(input["start"])
		monitor := ""
		for _, s := range input["result"] {
			monitor += encoding[SortStringByCharacter(s)]
		}
		val, err := strconv.Atoi(monitor)
		if err != nil {
			fmt.Fprintf(os.Stderr, "segments: value parse %v", err)
		}
		s += val
	}
	fmt.Println(s)
}

func findSetup(inputs []string) map[string]string {
	tempLen := make(map[int]string)
	tempLen[1] = findLen(inputs, 2)[0]
	tempLen[4] = findLen(inputs, 4)[0]
	tempLen[7] = findLen(inputs, 3)[0]
	tempLen[8] = findLen(inputs, 7)[0]

	// 9
	for _, s := range findLen(inputs, 6) {
		if len(stringDiff(s, tempLen[4])) == 2 {
			tempLen[9] = s
		}
	}

	// 6
	for _, s := range findLen(inputs, 6) {
		if len(stringDiff(s, tempLen[7])) == 4 {
			tempLen[6] = s
		}
	}

	// 5
	for _, s := range findLen(inputs, 5) {
		if len(stringDiff(tempLen[6], s)) == 1 {
			tempLen[5] = s
		}
	}

	// 3
	for _, s := range findLen(inputs, 5) {
		if len(stringDiff(tempLen[9], s)) == 1 && s != tempLen[5] {
			tempLen[3] = s
		}
	}

	for _, s := range findLen(inputs, 6) {
		if len(stringDiff(tempLen[8], s)) == 1 && s != tempLen[9] && s != tempLen[6] {
			tempLen[0] = s
		}
	}

	for _, s := range inputs {
		hit := false
		for _, v := range tempLen {
			if s == v {
				hit = true
			}
		}
		if !hit {
			tempLen[2] = s
		}
	}
	res := make(map[string]string)
	for k, v := range tempLen {
		res[SortStringByCharacter(v)] = strconv.Itoa(k)
	}
	return res
}

func stringDiff(s1, s2 string) string {
	result := ""
	for _, c1 := range s1 {
		in := false
		for _, c2 := range s2 {
			if c1 == c2 {
				in = true
			}
		}
		if !in {
			result += string(c1)
		}
	}
	return result
}

func findLen(inputs []string, l int) []string {
	res := []string{}
	for _, i := range inputs {
		if len(i) == l {
			res = append(res, i)
		}
	}
	return res
}

func readFile(f *os.File) []map[string][]string {
	inputs := []map[string][]string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), " | ")
		input := map[string][]string{
			"start":  strings.Split(lineParts[0], " "),
			"result": strings.Split(lineParts[1], " "),
		}
		inputs = append(inputs, input)
	}
	return inputs
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
