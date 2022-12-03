package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func readFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	res := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		first := text[:len(text)/2]
		second := text[len(text)/2:]
		res = append(res, []string{first, second})
	}

	return res, nil
}

func priorityLowercase(a rune) int {
	return int(a-'a') + 1
}

func priority(a rune) int {
	if a >= 'a' && a <= 'z' {
		return priorityLowercase(a)
	}
	return 26 + priorityLowercase(unicode.ToLower(a))
}

func solve(lines [][]string) int {
	res := 0

	commonString := func(str1, str2 string) rune {
		frq1 := map[rune]int{}

		for _, rune := range str1 {
			frq1[rune]++
		}

		for _, rune := range str2 {
			if _, ok := frq1[rune]; ok {
				return rune
			}
		}
		panic("No common char found")
	}

	for _, line := range lines {
		first := line[0]
		second := line[1]
		res += priority(commonString(first, second))
	}

	return res
}

func main() {
	filePath := "./in-day3.txt"
	lines, err := readFile(filePath)
	if err != nil {
		panic("Could not open file")
	}

	for _, line := range lines {
		fmt.Printf("%v\n", line)
	}
}


