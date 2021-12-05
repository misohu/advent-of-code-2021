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
		fmt.Fprintf(os.Stderr, "bingo: %v", err)
		os.Exit(1)
	}

	numbers, boards := readFile(f)
	border := 5

	for border <= len(numbers) {
		numbersMap := make(map[string]bool)
		for _, val := range numbers[:border] {
			numbersMap[val] = true
		}
		filtered := [][][]string{}
		if len(boards) == 1 {
			if evalBoard(boards[0], numbersMap) {
				sum := getNotChosenSum(boards[0], numbersMap)
				borderVal, err := strconv.Atoi(numbers[border-1])

				if err != nil {
					fmt.Fprintf(os.Stderr, "bingo: %v", err)
				}
				fmt.Printf("SUM: %v\nBORDERVAL: %v\nRESULT: %v\n", sum, borderVal, sum*borderVal)
				f.Close()
				os.Exit(0)
			}
		} else {
			for _, board := range boards {
				if !evalBoard(board, numbersMap) {
					filtered = append(filtered, board)
				}
			}
			boards = filtered
		}
		border += 1
	}
	f.Close()
}

func readFile(f *os.File) ([]string, [][][]string) {
	boards := [][][]string{}

	input := bufio.NewScanner(f)
	input.Scan()
	numbers := strings.Split(input.Text(), ",")
	input.Scan()

	for input.Scan() {
		new_board := [][]string{}
		for i := 0; i < 5; i++ {
			line := strings.Split(standardizeSpaces(input.Text()), " ")
			new_board = append(new_board, line)
			input.Scan()
		}
		boards = append(boards, new_board)
	}
	return numbers, boards
}

func evalBoard(board [][]string, numbers map[string]bool) bool {
	for _, row := range board {
		if sliceInNumbers(row, numbers) {
			return true
		}
	}
	for index, _ := range board[0] {
		column := []string{}
		for _, row := range board {
			column = append(column, row[index])
		}
		if sliceInNumbers(column, numbers) {
			return true
		}
	}
	return false
}

func sliceInNumbers(s []string, numbers map[string]bool) bool {
	for _, e := range s {
		if !numbers[e] {
			return false
		}
	}
	return true
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func getNotChosenSum(board [][]string, numbers map[string]bool) int {
	sum := 0
	for _, row := range board {
		for _, e := range row {
			if !numbers[e] {
				value, err := strconv.Atoi(e)
				if err != nil {
					fmt.Fprintf(os.Stderr, "getNotChosenSum: %v", err)
				}
				sum += value
			}
		}
	}
	return sum
}
