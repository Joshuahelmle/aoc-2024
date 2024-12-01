package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInputFile(inputFile string) ([]int, []int, error) {
	f, err := os.Open(inputFile)
	if err != nil {
		return nil, nil, err
	}
	scanner := bufio.NewScanner(f)
	l1, l2 := []int{}, []int{}
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		v1, err := strconv.Atoi((words[0]))
		if err != nil {
			return nil, nil, err
		}
		v2, err := strconv.Atoi((words[1]))
		if err != nil {

			return nil, nil, err
		}
		l1 = append(l1, v1)
		l2 = append(l2, v2)
	}
	return l1, l2, nil
}

func day1First() int {
	l1, l2, err := readInputFile("input/day1_1.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	slices.Sort(l1)
	slices.Sort(l2)
	sum := 0
	for i := range l1 {
		diff := l1[i] - l2[i]
		sum += int(math.Abs(float64(diff)))
	}
	return sum

}

func day1Second() int {

	l1, l2, err := readInputFile("input/day1_1.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	occurrences := make(map[int]int)
	for _, v := range l2 {
		occurrences[v]++
	}
	sum := 0
	for _, v := range l1 {
		sum += v * occurrences[v]
	}
	return sum

}
