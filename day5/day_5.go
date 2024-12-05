package day5

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() ([][]int, [][]int) {
	file, err := os.Open("input/day5_rules.txt")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	scanner := bufio.NewScanner(file)
	rules := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, []int{before, after})
	}
	file.Close()
	file, err = os.Open("input/day5_updates.txt")
	if err != nil {
		fmt.Println(err)
		defer file.Close()
	}
	updates := [][]int{}
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		update := []int{}
		for _, part := range parts {
			page, _ := strconv.Atoi(part)
			update = append(update, page)
		}
		updates = append(updates, update)
	}
	return rules, updates
}

func Part1() int {
	rules, updates := readInput()
	pageAllowedAfter := map[int][]int{}
	for _, rule := range rules {
		pageAllowedAfter[rule[1]] = append(pageAllowedAfter[rule[1]], rule[0])
	}
	pagesInUpdate := []map[int]struct{}{}
	//build the list of pages in each update
	for idx, update := range updates {
		pagesInUpdate = append(pagesInUpdate, map[int]struct{}{})
		for _, page := range update {
			pagesInUpdate[idx][page] = struct{}{}
		}
	}
	validUpdates := [][]int{}
	for _, update := range updates {
		if slices.IsSortedFunc(update, func(i, j int) int {
			if pagesAfterI, ok := pageAllowedAfter[i]; ok {
				for _, page := range pagesAfterI {
					if page == j {
						return 1
					}
				}
			}
			if pagesAfterJ, ok := pageAllowedAfter[j]; ok {
				for _, page := range pagesAfterJ {
					if page == i {
						return -1
					}
				}
			}

			return 0
		}) {
			validUpdates = append(validUpdates, update)
		}
	}
	sum := 0
	for _, update := range validUpdates {
		middle := len(update) / 2
		sum += update[middle]
	}

	return sum
}

func Part2() int {

	rules, updates := readInput()
	pageAllowedAfter := map[int][]int{}
	for _, rule := range rules {
		pageAllowedAfter[rule[1]] = append(pageAllowedAfter[rule[1]], rule[0])
	}
	pagesInUpdate := []map[int]struct{}{}
	//build the list of pages in each update
	for idx, update := range updates {
		pagesInUpdate = append(pagesInUpdate, map[int]struct{}{})
		for _, page := range update {
			pagesInUpdate[idx][page] = struct{}{}
		}
	}
	invalidUpdates := [][]int{}
	for _, update := range updates {
		if !slices.IsSortedFunc(update, func(i, j int) int {
			if pagesAfterI, ok := pageAllowedAfter[i]; ok {
				for _, page := range pagesAfterI {
					if page == j {
						return 1
					}
				}
			}
			if pagesAfterJ, ok := pageAllowedAfter[j]; ok {
				for _, page := range pagesAfterJ {
					if page == i {
						return -1
					}
				}
			}

			return 0
		}) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}
	for _, update := range invalidUpdates {
		slices.SortFunc(update, func(i, j int) int {
			if pagesAfterI, ok := pageAllowedAfter[i]; ok {
				for _, page := range pagesAfterI {
					if page == j {
						return 1
					}
				}
			}
			if pagesAfterJ, ok := pageAllowedAfter[j]; ok {
				for _, page := range pagesAfterJ {
					if page == i {
						return -1
					}
				}
			}

			return 0
		})
	}
	sum := 0
	for _, update := range invalidUpdates {
		middle := len(update) / 2
		sum += update[middle]
	}

	return sum
}
