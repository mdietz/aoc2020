package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var preambleLen int = 25

func main() {
	file, _ := os.Open("9.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var inputs []int = nil

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, num)
	}

	/*for i, e := range(inputs) {
		if i < preambleLen {
			continue
		} else {
			if test(inputs, i, e) {
				fmt.Printf("%d: valid\n", e)
			} else {
				fmt.Printf("%d: invalid\n", e)
				return
			}
		}
	}*/

	fmt.Println(inputs)
	fmt.Println(check(inputs))
}

func test(inputs []int, i int, e int) bool {
	for j, t := range(inputs[i-preambleLen:i]) {
		for k, s := range(inputs[i-preambleLen:i]) {
			if j == k {
				continue
			}
			if e == t+s {
				return true
			}
		}
	}
	return false
}

func check(inputs []int) int {
	val := 552655238
	size := 2
	for true {
		for i := 0; i < len(inputs)-size; i++ {
			sum := 0
			for j := 0; j < size; j++ {
				sum += inputs[i+j]
			}
			if sum == val {
				fmt.Println(size)
				for j := 0; j < size; j++ {
					fmt.Println(inputs[i+j])
				}
				final := inputs[i:i+size]
				sort.Ints(final)
				fmt.Println(final)
				return final[0] + final[len(final)-1]
			}
		}
		size++
	}
	return -1
}