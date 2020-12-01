package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var inputs []int

	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, line)
	}

	file.Close()

	for _, first := range inputs {
		for _, second := range inputs {
			for _, third := range inputs {
				if first+second+third == 2020 {
					fmt.Printf("First: %d, second: %d, third: %d, mult: %d\n",
						first, second, third, first*second*third)
				}
			}
		}
	}
}
