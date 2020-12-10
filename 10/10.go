package main

import (
	"../util"
	"fmt"
	"sort"
)

func main() {
	inputs := util.GetInput("10.in")
	var adapters []int = nil
	adapters = append(adapters, 0)
	for _, e := range(inputs) {
		adapters = append(adapters, util.Atoi(e))
	}
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	fmt.Println(adapters)

	count_1 := 0
	count_3 := 0
	for i, e := range(adapters) {
		if i == len(adapters)-1 {
			break
		}
		if adapters[i+1] - e == 1 {
			count_1++
		}
		if adapters[i+1] - e == 3 {
			count_3++
		}
	}
	fmt.Println(count_1*count_3)

	var workingSet [][]int = nil
	var finalSet [][]int = nil
	var first []int = nil
	first = append(first, 0)
	workingSet = append(workingSet, first)
	for len(workingSet) != 0 {
		c := workingSet[0]
		workingSet = workingSet[1:]
		ex := c[len(c)-1]
		currAdapter := adapters[ex]
		if currAdapter == adapters[len(adapters)-1] {
			finalSet = append(finalSet, c)
		}
		next := ex + 1
		for true {
			if next < len(adapters) &&  adapters[next] - currAdapter <= 3 {
				cop := make([]int, len(c))
				copy(cop, c)
				cop = append(cop, next)
				workingSet = append(workingSet, cop)
			} else {
				break
			}
			next++
		}
	}
	fmt.Println(len(finalSet))
}