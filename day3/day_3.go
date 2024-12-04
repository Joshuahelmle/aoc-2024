package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func calculateFromMatch(match []string) int {
	v1, err := strconv.Atoi(match[1])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	v2, err := strconv.Atoi(match[2])
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return v1 * v2
}

func parseFile(regex *regexp.Regexp, file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var data [][]string
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindAllStringSubmatch(line, -1)
		if len(matches) > 0 {
			data = append(data, matches...)

		}
	}
	return data
}

func Part1() int {
	re := regexp.MustCompile(`mul\(([0-9]{4,}),([0-9]{4,})\)`)
	input, err := os.Open("input/day3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()
	data := parseFile(re, input)
	sum := 0
	if len(data) > 0 {
		for _, match := range data {
			if len(match) > 0 {
				sum += calculateFromMatch(match)
			}
		}
	}

	return sum
}
func Part2() int {
	re := regexp.MustCompile(`do\(\)|don\'t\(\)|mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	input, err := os.Open("input/day3.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()
	data := parseFile(re, input)
	sum := 0
	skip := false
	if len(data) > 0 {
		for _, match := range data {
			if len(match) > 0 {
				if match[0] == "do()" {
					skip = false
				} else if match[0] == "don't()" {
					skip = true
				} else if !skip {
					sum += calculateFromMatch(match)
				}
			}
		}
	}
	return sum
}
