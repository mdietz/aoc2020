package util

import (
	"bufio"
	"os"
	"strconv"
)

func GetInput(filename string) []string {
	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var out []string = nil

	for scanner.Scan() {
		out = append(out, scanner.Text())
	}
	return out
}

func Atoi(s string) int {
	a, _ := strconv.Atoi(s)
	return a
}