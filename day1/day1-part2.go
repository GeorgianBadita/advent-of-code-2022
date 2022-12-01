package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func readFromFile(filePath string) (map[int]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	currElf := 1
	res := map[int]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			currElf++
			continue
		}
		conv, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		res[currElf] += conv
	}

	return res, nil
}

func main() {
	path := "./in-day1.txt"
	data, err := readFromFile(path)
	if err != nil {
		fmt.Println(err)
		panic("Error reading file")
	}

	max1 := math.MinInt32
	max2 := math.MinInt32
	max3 := math.MinInt32

	for _, v := range data {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}

	fmt.Println(max1 + max2 + max3)
}
