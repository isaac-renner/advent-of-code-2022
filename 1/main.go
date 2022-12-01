package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	println("day 1")
	buff, err := ioutil.ReadFile("cals.txt")
	if err != nil {
		println(err)
		return
	}
	fileAsString := string(buff)
	mealsSlice := strings.Split(fileAsString, "\n")
	elfSlice := [][]int{{}}
	for _, mealOrEmpty := range mealsSlice {
		if mealOrEmpty == "" {
			elfSlice = append(elfSlice, []int{})
			continue
		}

		cals, err := strconv.Atoi(mealOrEmpty)

		if err != nil {
			return
		}

		elfSlice[len(elfSlice)-1] = append(elfSlice[len(elfSlice)-1], cals)
	}

	mostCal := 0
	elfWithMostCalIndex := 0
	for index, elf := range elfSlice {
		cal := countAllCal(elf)
		if cal > mostCal {
			mostCal = cal
			elfWithMostCalIndex = index
		}
	}
	println(mostCal)
	println(elfWithMostCalIndex)
}

func countAllCal(elf []int) int {
	cals := 0
	for _, meal := range elf {
		cals += meal
	}

	return cals
}
