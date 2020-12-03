package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("3.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var treeMap [500][500]bool

	i := 0
	wrap := 0
	for scanner.Scan() {
		line := scanner.Text()
		for j, e := range(line) {
			if e == '#' {
				treeMap[i][j] = true
			} else {
				treeMap[i][j] = false
			}
		}
		wrap = len(line) - 1
		i = i + 1
	}

	for _, e := range treeMap {
		for _, f := range e {
			if f {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}

	row := 0
	col := 0
	count := 0
	fmt.Println(wrap)
	for row < i - 1 {
		row = row + 2
		col = col + 1

		if col > wrap {
			col = col % wrap - 1
		}

			fmt.Printf("row: %d, col: %d, val: %d\n", row, col, treeMap[row][col])

		if treeMap[row][col] {
			count++
		}
	}

	file.Close()

	fmt.Printf("Result: %d\n", count)
}
