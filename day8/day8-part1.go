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
	numVisible := 2*len(matrix) + 2*len(matrix[0]) - 4

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[0])-1; col++ {
			cRow := row - 1
			cCol := col

			// Try up
			visible := true
			for cRow >= 0 {
				if matrix[cRow][cCol] >= matrix[row][col] {
					visible = false
					break
				}
				cRow--
			}

			if visible {
				numVisible++
				continue
			}

			// Try down
			cRow = row + 1
			visible = true
			for cRow < len(matrix) {
				if matrix[cRow][cCol] >= matrix[row][col] {
					visible = false
					break
				}
				cRow++
			}

			if visible {
				numVisible++
				continue
			}

			// Try left
			cRow = row
			cCol = col - 1
			visible = true

			for cCol >= 0 {
				if matrix[cRow][cCol] >= matrix[row][col] {
					visible = false
					break
				}
				cCol--
			}

			if visible {
				numVisible++
				continue
			}

			// Try right
			cCol = col + 1
			visible = true

			for cCol < len(matrix[0]) {
				if matrix[cRow][cCol] >= matrix[row][col] {
					visible = false
					break
				}
				cCol++
			}

			if visible {
				numVisible++
			}
		}
	}

	return numVisible
}

func main() {
	matrix, err := readFromFile("./in-day8.txt")
	if err != nil {
		panic("Error reading from file...")
	}

	fmt.Printf("%v\n", Solve(matrix))
}
