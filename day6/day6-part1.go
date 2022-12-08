package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", nil
}

func solve(str string) int {
	occs := map[rune]int{}
	length := 0
	idx := 0
	for idx < len(str) {
		chr := rune(str[idx])
		if val, ok := occs[chr]; ok {
			length = 0
			idx = val
			occs = map[rune]int{}
		} else {
			occs[chr] = idx
			length++
			if length == 4 {
				return idx
			}
		}
		idx++
	}
	return -1
}

func main() {
	str, err := readFile("./in-day6.txt")
	if err != nil {
		panic("Error reading input...")
	}

	res := solve(str)
	fmt.Println(res + 1)
}
