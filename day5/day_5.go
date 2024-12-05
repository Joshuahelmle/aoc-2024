package day5

import (
	"bufio"
	"fmt"
	"os"
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
	for idx, update := range updates {
		pagesSeen := map[int]struct{}{}
		isValidUpdate := true
		for _, page := range update {
			//check if the page needs to be after another page
			if _, ok := pageAllowedAfter[page]; ok {
				isPageAllowed := true
				for _, allowedAfter := range pageAllowedAfter[page] {
					//if the allowed page isn't in the same update, we don't need to check if it is before the current page
					if _, ok := pagesInUpdate[idx][allowedAfter]; !ok {
						continue
					}
					//if the allowed page is in the same update, we need to check if it is before the current page
					if _, ok := pagesSeen[allowedAfter]; !ok {
						isPageAllowed = false
						break
					}
				}
				if !isPageAllowed {
					isValidUpdate = false
					break
				}
			}
			//mark the page as seen
			pagesSeen[page] = struct{}{}
		}
		if isValidUpdate {
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
