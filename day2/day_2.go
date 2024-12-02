package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func readfile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var data [][]int
	for scanner.Scan() {
		var line []int
		for _, v := range strings.Fields(scanner.Text()) {
			num, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			line = append(line, num)
		}
		data = append(data, line)
	}
	return data, nil
}

func checkCondition(row []int) bool {

	isAscending := true
	isDescending := true
	for i := 0; i < len(row)-1; i++ {
		diff := math.Abs(float64(row[i] - row[i+1]))
		if diff < 1 || diff > 3 {
			return false
		}

		if row[i] > row[i+1] {
			isAscending = false
		}

		if row[i] < row[i+1] {

			isDescending = false
		}
		//early return
		if !isAscending && !isDescending {
			return false
		}
	}

	return isAscending || isDescending

}

func Part1() int {
	data, err := readfile("input/day_2.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	count := 0
	for _, row := range data {
		if checkCondition(row) {
			count++
		}
	}
	return count
}

func Part2() int {
	data, err := readfile("input/day_2.txt")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	count := 0
	//naive solution, could be optimized
	for _, row := range data {
		if checkCondition(row) {
			count++
		} else {
			for i := 0; i < len(row); i++ {
				newRow := make([]int, 0, len(row)-1)
				newRow = append(newRow, row[:i]...)
				newRow = append(newRow, row[i+1:]...)
				if checkCondition(newRow) {
					count++
					break
				}
			}
		}
	}
	return count
}
