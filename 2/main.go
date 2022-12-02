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
		totalScore += scoreFromOutcome(myChoice, opChoice)
		totalScore += scoreFromChoice(myChoice)
	}

	fmt.Println(totalScore)
}

func scoreFromChoice(myChoice string) int {
	score := choiceScoreMap[myChoice]

	return score
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
