package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFromFile(filePath string) ([][]rune, [][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}

	stacks := [][]rune{}
	moves := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "\n" || len(line) == 0 {
			break
		}

		parts := strings.Split(line, " ")

		stackIdx := 0
		for idx := 0; idx < len(parts); idx++ {
			part := parts[idx]
			trimmedPart := strings.Trim(part, " []\n")
			if len(trimmedPart) == 0 {
				if len(stacks) <= stackIdx {
					stacks = append(stacks, []rune{})
				}
				stackIdx++
				idx += 3
				continue
			}
			if len(stacks) <= stackIdx {
				stacks = append(stacks, []rune{})
			}
			if trimmedPart[0] >= '1' && trimmedPart[0] <= '9' {
				break
			}
			stacks[stackIdx] = append(stacks[stackIdx], rune(trimmedPart[0]))
			stackIdx++
		}
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		amountStr := strings.Trim(parts[1], " ")
		fromStr := strings.Trim(parts[3], " ")
		toStr := strings.Trim(parts[5], " ")

		amount, _ := strconv.Atoi(amountStr)
		from, _ := strconv.Atoi(fromStr)
		to, _ := strconv.Atoi(toStr)

		moves = append(moves, []int{from - 1, to - 1, amount})
	}
	for idx := 0; idx < len(stacks); idx++ {
		rev(&stacks[idx])
	}
	return stacks, moves, nil
}

func rev(rn *[]rune) {
	for idx := 0; idx < len(*rn)/2; idx++ {
		(*rn)[idx], (*rn)[len(*rn)-idx-1] = (*rn)[len(*rn)-idx-1], (*rn)[idx]
	}
}

func solve(stacks [][]rune, moves [][]int) string {
	for _, mv := range moves {
		from, to, amount := mv[0], mv[1], mv[2]
		rest := stacks[from][len(stacks[from])-amount:]
		stacks[from] = stacks[from][:len(stacks[from])-amount]
		stacks[to] = append(stacks[to], rest...)
	}
	res := ""
	for _, stack := range stacks {
		res += string(stack[len(stack)-1])
	}
	return res
}

func main() {
	stacks, moves, err := readFromFile("./in-day5.txt")
	if err != nil {
		panic("Error reading input...")
	}

	fmt.Println(solve(stacks, moves))
}
