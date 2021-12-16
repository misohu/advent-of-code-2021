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
		fmt.Fprintf(os.Stderr, "labirint: problem with file %v", err)
		os.Exit(1)
	}
	i, err := readFile(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labirint: problem with file %v", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", i)
	version, _, result, err := parseInput(i, true)
	if err != nil {
		fmt.Fprintf(os.Stderr, "labirint: problem with file %v", err)
		os.Exit(1)
	}
	fmt.Printf("%d\n", result)
	fmt.Printf("%d\n", version)
}

func parseInput(input string, pad bool) (int, string, int, error) {
	fmt.Println("PARSING: ", input)
	original := input
	version, err := parseInt(input[:3])
	result := 0
	if err != nil {
		return 0, "", 0, fmt.Errorf("parseInput: %v", err)
	}
	input = input[3:]
	packageType, err := parseInt(input[:3])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parseInput: %v", err)
	}
	// var val string
	var versions int
	if packageType == 4 {
		_, input, result, err = parseLiteral(input[3:])
		if err != nil {
			return 0, "", 0, fmt.Errorf("parseInput: %v", err)
		}
	} else {
		versions, input, result, err = parseOperator(input[3:], packageType)
		if err != nil {
			return 0, "", 0, fmt.Errorf("parseInput: %v", err)
		}
		version += versions
	}
	if pad {
		currLen := len(original) - len(input)
		input = input[(4-currLen%4)%4:]
		for len(input) > 0 {
			testNum, err := parseInt(input[:4])
			if err != nil {
				return 0, "", 0, fmt.Errorf("parseInput: %v", err)
			}
			if testNum > 0 {
				break
			}
			input = input[4:]
		}
	}

	fmt.Println("VERSION: ", version)
	// fmt.Println(version, packageType, len(input), val)
	return version, input, result, nil
}

var operatorMap = map[int]func([]int) int{
	0: customSum,
	1: customProduct,
	2: customMin,
	3: customMax,
	5: customGreater,
	6: customLess,
	7: customEqual,
}

func parseOperator(s string, operatorType int) (int, string, int, error) {
	fmt.Println("PARSING OPERATOR: ", s)
	res := 0
	lengthType := string(s[0])
	s = s[1:]
	var vals []int
	if lengthType == "0" {
		nextLength, err := parseInt(s[:15])
		if err != nil {
			return 0, "", 0, fmt.Errorf("parseOperator: %v", err)
		}
		s = s[15:]
		tmp := s[:nextLength]
		s = s[nextLength:]
		for len(tmp) > 0 {
			version, input, val, err := parseInput(tmp, false)
			if err != nil {
				return 0, "", 0, fmt.Errorf("parseOperator: %v", err)
			}
			res += version
			tmp = input
			vals = append(vals, val)
		}
	} else {
		nextPackets, err := parseInt(s[:11])
		if err != nil {
			return 0, "", 0, fmt.Errorf("parseOperator: %v", err)
		}
		s = s[11:]
		fmt.Printf("nextPackets -> %d\n", nextPackets)
		for i := 0; i < nextPackets; i++ {
			version, input, val, err := parseInput(s, false)
			if err != nil {
				return 0, "", 0, fmt.Errorf("parseOperator: %v", err)
			}
			res += version
			vals = append(vals, val)
			s = input
		}
	}
	fmt.Println("VALUES: ", vals)
	return res, s, operatorMap[operatorType](vals), nil
}

func customSum(vals []int) int {
	res := 0
	for _, val := range vals {
		res += val
	}
	return res
}

func customProduct(vals []int) int {
	res := 1
	for _, val := range vals {
		res *= val
	}
	return res
}

func customMin(vals []int) int {
	res := vals[0]
	for _, val := range vals {
		if val < res {
			res = val
		}
	}
	return res
}

func customMax(vals []int) int {
	res := vals[0]
	for _, val := range vals {
		if val > res {
			res = val
		}
	}
	return res
}

func customGreater(vals []int) int {
	if vals[0] > vals[1] {
		return 1
	}
	return 0
}

func customLess(vals []int) int {
	if vals[0] < vals[1] {
		return 1
	}
	return 0
}

func customEqual(vals []int) int {
	if vals[0] == vals[1] {
		return 1
	}
	return 0
}

func parseLiteral(s string) (string, string, int, error) {
	fmt.Println("PARSING LITERAL: ", s)
	val := ""
	for string(s[0]) == "1" {
		// fmt.Println(s, val)
		val += s[1:5]
		s = s[5:]
	}
	val += s[1:5]
	s = s[5:]
	valInt, err := strconv.ParseInt(val, 2, 64)
	if err != nil {
		return "", "", 0, fmt.Errorf("parseLiteral: %v", err)
	}
	fmt.Printf("valInt -> %d\n", valInt)
	return val, s, int(valInt), nil
}

func parseInt(s string) (int, error) {
	if i, err := strconv.ParseInt(s, 2, 64); err != nil {
		return 0, fmt.Errorf("parseInt: %v", err)
	} else {
		return int(i), nil
	}
}

func readFile(f *os.File) (string, error) {
	help := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	reader := bufio.NewScanner(f)
	reader.Scan()
	data := strings.Split(reader.Text(), "")
	res := ""
	for _, d := range data {
		res += help[d]
	}

	return res, nil
}
