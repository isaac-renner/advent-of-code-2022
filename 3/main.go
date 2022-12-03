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

		middle := len(round) / 2
		first := round[:middle]
		second := round[middle:]
		if round == "" {
			continue
		}

		var matchingLetter rune
		for _, letter := range first {
			if strings.IndexRune(second, letter) != -1 {
				matchingLetter = letter
				break
			}
		}
		values += int(getScore(matchingLetter))
	}
	fmt.Println(values)

}

func getScore(r rune) rune {
	if r < 91 {
		return r - 38
	}

	return r - 96
}
