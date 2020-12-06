package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("6.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var qs []map[rune]int = nil

	m := make(map[rune]int)
	gsize := 0
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			qs = append(qs, m)
			for _, v := range(m) {
				if v == gsize {
					count++
				}
			}
			m = make(map[rune]int)
			gsize = 0
			continue
		}
		for _, e := range(line) {
			m[e]++
		}
		gsize++
	}
	fmt.Println(count)
}