package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFromFile(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	res := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		tmp := []int{}
		for _, ch := range line {
			num, _ := strconv.Atoi(string(ch))
			tmp = append(tmp, num)
		}
		res = append(res, tmp)
	}
	return res, nil
}

func Solve(matrix [][]int) int {
	maxScenicScore := 0

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[0])-1; col++ {
			cRow := row - 1
			cCol := col
			scenicScore := 1
			inc := 1

			// Try up
			for cRow >= 0 {
				if matrix[cRow][cCol] >= matrix[row][col] {
					break
				}
				cRow--
				inc++
			}
			if cRow < 0 {
				inc--
			}

			scenicScore *= inc

			// Try down
			cRow = row + 1
			inc = 1
			for cRow < len(matrix) {
				if matrix[cRow][cCol] >= matrix[row][col] {
					break
				}
				cRow++
				inc++
			}
			if cRow >= len(matrix) {
				inc--
			}

			scenicScore *= inc

			// Try left
			cRow = row
			cCol = col - 1
			inc = 1
			for cCol >= 0 {
				if matrix[cRow][cCol] >= matrix[row][col] {
					break
				}
				cCol--
				inc++
			}
			if cCol < 0 {
				inc--
			}

			scenicScore *= inc

			// Try right
			cCol = col + 1
			inc = 1
			for cCol < len(matrix[0]) {
				if matrix[cRow][cCol] >= matrix[row][col] {
					break
				}
				cCol++
				inc++
			}
			if cCol >= len(matrix[0]) {
				inc--
			}

			scenicScore *= inc

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}

	return maxScenicScore
}

func main() {
	matrix, err := readFromFile("./in-day8.txt")
	if err != nil {
		panic("Error reading from file...")
	}

	fmt.Printf("%v\n", Solve(matrix))
}
