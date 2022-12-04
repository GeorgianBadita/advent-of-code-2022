package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	x, y int
}

type Assignment struct {
	p1, p2 Pair
}

func readFromFile(filePath string) ([]Assignment, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	res := []Assignment{}

	for scanner.Scan() {
		line := scanner.Text()
		assignments := strings.Split(line, ",")
		firstPair := strings.Split(assignments[0], "-")
		secondPair := strings.Split(assignments[1], "-")
		pair11, _ := strconv.Atoi(firstPair[0])
		pair12, _ := strconv.Atoi(firstPair[1])
		pair21, _ := strconv.Atoi(secondPair[0])
		pair22, _ := strconv.Atoi(secondPair[1])
		if pair11 > pair21 {
			pair11, pair21 = pair21, pair11
			pair12, pair22 = pair22, pair12
		}
		res = append(res, Assignment{p1: Pair{pair11, pair12}, p2: Pair{pair21, pair22}})
	}
	return res, nil
}

func fullInclusion(assignment Assignment) bool {
	firstPair := assignment.p1
	secondPair := assignment.p2
	return (firstPair.x >= secondPair.x && firstPair.y <= secondPair.y) || (secondPair.x >= firstPair.x && secondPair.y <= firstPair.y)
}

func solve(assignemnts []Assignment) int {
	res := 0
	for _, ass := range assignemnts {
		if fullInclusion(ass) {
			res++
		}
	}
	return res
}

func main() {
	assignments, err := readFromFile("./in-day4.txt")
	if err != nil {
		panic("Could not read from file")
	}

	fmt.Println(solve(assignments))
}
