package day4

import (
	"bufio"
	"os"
)

func readInput() [][]rune {
	file, err := os.Open("input/day4.txt")
	if err != nil {
		defer file.Close()
	}
	scanner := bufio.NewScanner(file)
	grid := [][]rune{}
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}
	return grid
}

func checkWord(word []rune, x, y int, grid [][]rune, direction func(int, int) (int, int)) bool {
	if len(word) == 0 {
		return true
	} else if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return false
	}
	nextX, nextY := direction(x, y)
	return word[0] == grid[x][y] && checkWord(word[1:], nextX, nextY, grid, direction)
}

func Part1() int {

	sum := 0
	grid := readInput()
	directions := []func(int, int) (int, int){
		func(x, y int) (int, int) { return x + 1, y },
		func(x, y int) (int, int) { return x - 1, y },
		func(x, y int) (int, int) { return x, y + 1 },
		func(x, y int) (int, int) { return x, y - 1 },
		func(x, y int) (int, int) { return x + 1, y + 1 },
		func(x, y int) (int, int) { return x - 1, y - 1 },
		func(x, y int) (int, int) { return x + 1, y - 1 },
		func(x, y int) (int, int) { return x - 1, y + 1 },
	}
	for y, row := range grid {
		for x := range row {
			for _, direction := range directions {
				if checkWord([]rune("XMAS"), x, y, grid, direction) {
					sum += 1
				}
			}
		}
	}
	return sum

}
