package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	totalElfCal := []int{}
	for _, elf := range elfSlice {
		cal := countAllCal(elf)
		totalElfCal = append(totalElfCal, cal)
	}

	sort.Ints(totalElfCal)
	fmt.Println(totalElfCal[len(totalElfCal)-3:])
}

func countAllCal(elf []int) int {
	cals := 0
	for _, meal := range elf {
		cals += meal
	}

	return cals
}
