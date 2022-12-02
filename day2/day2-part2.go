package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Option int64

const (
	Rock     Option = 1
	Paper    Option = 2
	Scissors Option = 3
)

type Round struct {
	Player1Option, Player2Option Option
}

func readFile(filePath string) ([]Round, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	mapping := map[string]Option{
		"A": Rock,
		"B": Paper,
		"C": Scissors,
	}

	rounds := []Round{}

	for scanner.Scan() {
		line := scanner.Text()
		elems := strings.Split(line, " ")
		first := elems[0]
		second := elems[1]
		player1Option := mapping[first]
		player2Option := Rock
		if second == "Y" {
			player2Option = player1Option
		} else if second == "X" {
			if player1Option == Rock {
				player2Option = Scissors
			} else if player1Option == Paper {
				player2Option = Rock
			} else if player1Option == Scissors {
				player2Option = Paper
			}
		} else if second == "Z" {
			if player1Option == Rock {
				player2Option = Paper
			} else if player1Option == Paper {
				player2Option = Scissors
			} else if player1Option == Scissors {
				player2Option = Rock
			}
		}
		rounds = append(rounds, Round{Player1Option: player1Option, Player2Option: player2Option})
	}

	return rounds, nil
}

func roundScore(round Round) int {
	score := 1
	if round.Player2Option == Paper {
		score = 2
	} else if round.Player2Option == Scissors {
		score = 3
	}

	roundRes := result(round.Player1Option, round.Player2Option)

	if roundRes == 0 {
		return score + 3
	} else if roundRes == 1 {
		return score
	}
	return score + 6
}

// Returns 1 if pl1 wins, 0 if draw, -1 if player2 wins
func result(pl1 Option, pl2 Option) int {
	if pl1 == pl2 {
		return 0
	}
	if (pl1 == Rock && pl2 == Scissors) || (pl1 == Paper && pl2 == Rock) || (pl1 == Scissors && pl2 == Paper) {
		return 1
	}
	return -1
}

func main() {
	path := "./in-day2.txt"
	rounds, err := readFile(path)
	if err != nil {
		panic("")
	}

	score := 0
	for _, round := range rounds {
		score += roundScore(round)
	}
	fmt.Printf("Score: %v\n", score)
}
