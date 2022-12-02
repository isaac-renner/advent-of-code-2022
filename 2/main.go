package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var choiceScoreMap = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func main() {
	fmt.Println("Day 2")
	winScore := make(map[int]int)
	winScore[1] = 6
	winScore[0] = 3
	winScore[-1] = 0
	buff, err := ioutil.ReadFile("data.txt")
	if err != nil {
		println(err)
		return
	}
	fileAsString := string(buff)

	totalScore := 0
	for _, round := range strings.Split(fileAsString, "\n") {
		if len(round) != 3 {
			break
		}
		choices := strings.Split(round, " ")
		opChoice := choices[0]
		myChoice := choices[1]
		score, outcome := getChoiceAndOutcome(myChoice, opChoice)
		totalScore += score
		totalScore += scoreFromChoice(outcome)
	}

	fmt.Println(totalScore)
}

func scoreFromChoice(myChoice string) int {
	score := choiceScoreMap[myChoice]

	return score
}

var drawMap = map[string]string{
	"A": "X",
	"B": "Y",
	"C": "Z",
}

var winMap = map[string]string{
	"A": "Y",
	"B": "Z",
	"C": "X",
}

var loseMap = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func getChoiceAndOutcome(myChoice string, opChoice string) (int, string) {
	newDec := getChoice(myChoice, opChoice)

	if myChoice == "X" {
		return 0, newDec
	}

	if myChoice == "Z" {
		return 6, newDec
	}

	return 3, newDec
}

func getChoice(myChoice string, opChoice string) string {
	if myChoice == "Y" {
		return drawMap[opChoice]
	}

	if myChoice == "X" {
		return loseMap[opChoice]
	}

	if myChoice == "Z" {
		return winMap[opChoice]
	}

	panic("something went wrong")
}

func scoreFromOutcome(myChoice string, opChoice string) int {
	opScore := choiceScoreMap[opChoice]
	myScore := choiceScoreMap[myChoice]

	if opScore == myScore {
		return 3
	}

	if opChoice == "A" && myChoice == "Y" {
		return 6
	}

	if opChoice == "B" && myChoice == "Z" {
		return 6
	}

	if opChoice == "C" && myChoice == "X" {
		return 6
	}

	println("lose")
	return 0
}
