package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, _ := os.Open("2.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := strings.Split(line, ":")[0]
		second := strings.Split(line, ":")[1][1:]

		nums := strings.Split(first, " ")[0]
		a, _ := strconv.Atoi(strings.Split(nums, "-")[0])
		b, _ := strconv.Atoi(strings.Split(nums, "-")[1])
		c := strings.Split(first, " ")[1]

		count := 0
		for loc, e := range second {
			if loc+1 == a || loc+1 == b {
				if string(e) == c {
					count++
				}
			}
		}
		if count == 1 {
			total++
		}
	}

	fmt.Printf("Total: %d\n", total)

	file.Close()
}
