package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	x, y int
}

func ReadFromFile(filePath string) ([][]int, []int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	res := [][]int{}
	source := []int{}
	dest := []int{}

	for scanner.Scan() {
		line := scanner.Text()

		intLine := []int{}
		for jdx, char := range line {
			if char == 'S' {
				source = []int{len(res), jdx}
				intLine = append(intLine, 0)
			} else if char == 'E' {
				dest = []int{len(res), jdx}
				intLine = append(intLine, 26)
			} else {
				intLine = append(intLine, int(char-'a'))
			}
		}
		res = append(res, intLine)
	}

	return res, source, dest, nil
}

func ValidCoords(matrix [][]int, x, y int) bool {
	return x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0])
}

func Solve(matrix [][]int, start, end []int) int {
	matrixSol := [][]int{}
	for idx := 0; idx < len(matrix); idx++ {
		line := []int{}
		for jdx := 0; jdx < len(matrix[0]); jdx++ {
			line = append(line, 0)
		}
		matrixSol = append(matrixSol, line)
	}

	queue := []Pair{{x: start[0], y: start[1]}}
	visited := map[Pair]bool{}

	endP := Pair{x: end[0], y: end[1]}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}

	for len(queue) > 0 {
		currPos := queue[0]
		queue = queue[1:]

		if currPos == endP {
			break
		}

		for idx := 0; idx < len(dx); idx++ {
			newX := currPos.x + dx[idx]
			newY := currPos.y + dy[idx]

			newPair := Pair{x: newX, y: newY}
			if _, ok := visited[newPair]; !ok && ValidCoords(matrix, newX, newY) && (matrix[newX][newY] <= matrix[currPos.x][currPos.y] || matrix[newX][newY]-1 == matrix[currPos.x][currPos.y]) {
				visited[newPair] = true
				queue = append(queue, newPair)
				matrixSol[newPair.x][newPair.y] = 1 + matrixSol[currPos.x][currPos.y]
			}
		}
	}

	return matrixSol[endP.x][endP.y]
}

func main() {
	matrix, start, end, err := ReadFromFile("./in-day12.txt")
	if err != nil {
		panic("Could not read from file...")
	}

	fmt.Println(Solve(matrix, start, end))
}
