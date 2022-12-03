package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	buff, err := ioutil.ReadFile("data.txt")
	if err != nil {
		println(err)
		return
	}
	fileAsString := string(buff)

	values := 0
	agg := []string{}
	for _, round := range strings.Split(fileAsString, "\n") {
		agg = append(agg, round)

		if len(agg) < 3 {
			continue
		}

		first := agg[0]
		second := agg[1]
		third := agg[2]
		if round == "" {
			continue
		}

		var matchingLetter rune
		for _, letter := range first {
			if strings.IndexRune(second, letter) != -1 {
				if strings.IndexRune(third, letter) != -1 {
					matchingLetter = letter
				}
			}
		}

		fmt.Println(int(getScore(matchingLetter)), agg)
		values += int(getScore(matchingLetter))
		agg = []string{}
	}
	fmt.Println(values)

}

func getScore(r rune) rune {
	if r < 91 {
		return r - 38
	}

	return r - 96
}
