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
	group := []string{}
	idx := 1
	for scanner.Scan() {
		group = append(group, scanner.Text())
		if idx%3 == 0 {
			res = append(res, group)
			group = []string{}
		}
		idx++
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

	commonString := func(str1, str2 string) string {
		frq1 := map[rune]int{}

		for _, rune := range str1 {
			frq1[rune]++
		}

		res := ""
		for _, rune := range str2 {
			if _, ok := frq1[rune]; ok {
				res += string(rune)
			}
		}
		return res
	}

	for _, line := range lines {
		first := line[0]
		second := line[1]
		third := line[2]
		res += priority(rune(commonString(commonString(first, second), third)[0]))
	}

	return res
}

func main() {
	filePath := "./in-day3.txt"
	lines, err := readFile(filePath)
	if err != nil {
		panic("Could not open file")
	}
	fmt.Println(solve(lines))
}
