package main

import (
	"fmt"
)

func main() {
	inputs := []int{0,6,1,7,2,19,20}
	var out []int = nil 
	var seenMap map[int]int = make(map[int]int) 
	idx := 1
	lastNum := inputs[0]
	for _, e := range(inputs[1:]) {
		out = append(out, lastNum)
		seenMap[lastNum] = idx
		lastNum = e
		idx++
	}
	fmt.Println(seenMap)
	fmt.Println(lastNum)
	for true {
		if idx == 30000000 {
			fmt.Println(lastNum)
			return
		}
		val, ok := seenMap[lastNum]
		if !ok {
			//out = append(out, lastNum)
			seenMap[lastNum] = idx
			//fmt.Printf("idx:%d spoken:%d val:%d\n", idx, lastNum, 0)
			lastNum = 0
		} else {
			//out = append(out, val)
			seenMap[lastNum] = idx
			//fmt.Printf("idx:%d spoken:%d val:%d\n", idx, lastNum, idx-val)
			lastNum = idx-val
		}
		idx++
	}
}