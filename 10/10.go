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

	var adaptersWithRoutes [][]int = nil
	for _, e := range(adapters) {
		adaptersWithRoutes = append(adaptersWithRoutes, []int{e, 0})
	}

	adaptersWithRoutes[0][1]++

	for i, e := range(adaptersWithRoutes) {
		if i == 0 {
			continue
		}
		last := i - 1
		fmt.Println(adaptersWithRoutes)
		for true {
			if last >= 0 && e[0] - adaptersWithRoutes[last][0] <= 3 {
				adaptersWithRoutes[i][1] += adaptersWithRoutes[last][1]
			} else {
				break
			}
			last--
		}
	}
	fmt.Println(adaptersWithRoutes[len(adaptersWithRoutes)-1][1])
}