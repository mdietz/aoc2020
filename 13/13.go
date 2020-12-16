package main

import (
	"../util"
	"fmt"
	"strings"
)

func main() {
	inputs := util.GetInput("13.in")
	for _, e := range(inputs) {
		fmt.Println(e)
	}
	earliest := util.Atoi(inputs[0])
	fmt.Println(earliest)
	busIds := strings.Split(strings.ReplaceAll(inputs[1], "x,", ""), ",")
	fmt.Println(busIds)
	closest := 1000000000
	closBus := 0
	for _, e := range(busIds) {
		fmt.Println(earliest % util.Atoi(e))
		nextDep := earliest - earliest % util.Atoi(e) + util.Atoi(e)
		if nextDep - earliest < closest {
			closest = nextDep - earliest
			closBus = util.Atoi(e)
		}
	}
	fmt.Println((closest)*closBus)

	fmt.Println("====================")

	busList := strings.Split(inputs[1], ",")
	var busInts []int = nil
	var idLoc map[int]int = make(map[int]int)
	var xLoc []int = nil

	for i, e := range(busList) {
		if (e == "x") {
			xLoc = append(xLoc, i)
		} else {
			busInts = append(busInts, util.Atoi(e))
			idLoc[util.Atoi(e)] = i
		}
	}
	fmt.Println(xLoc)
	fmt.Println(busInts)
	fmt.Println(idLoc)

	//var init []int = nil
	//var increment []int = nil
	accum := 1
	for i := 0; i < len(busInts); i++ {
		red := 0
		if idLoc[busInts[i]] > busInts[i] {
			red = idLoc[busInts[i]] % busInts[i]
		} else {
			red = idLoc[busInts[i]]
		}
		first, second := findDistance(busInts[0], busInts[i], red)
		if first == 1 {
			accum = lcm(accum, second)
		}
		fmt.Printf("Init: %d, increment: %d\n", first, second)
	}
	fmt.Printf("Accum: %d\n", accum*7)

	val := accum*7
	for true {
		count := 0
		for k, v := range(idLoc) {
			if mod(val+v, k) != 0 {
				break
			}
			if count == len(idLoc)-1 {
				fmt.Printf("Res: %d\n", val)
			}
		}
		val += accum*7
	}

	/*possible := busInts[0] - idLoc[busInts[0]]
	count := 0
	for true {
		if mod(count, 100000000) == 0 {
			fmt.Println(possible)
		}
		for i := 1; i < len(busInts); i++ {
			if mod(possible, busInts[i]) != busInts[i] - idLoc[busInts[i]] + idLoc[busInts[0]] {
				break
			}
			if i == len(busInts)-1 {
				fmt.Printf("Done: %d\n", possible)
				return
			}
		}
		possible += busInts[0]
		count++
	}*/

	/*
	val := busInts[0] - idLoc[busInts[0]]
	for _, e := range(busInts[1:]) {
		fmt.Println(val)
		val = lcm(val, e)
	}
	fmt.Println(val)
	*/

	/*
	possible := busInts[len(busInts)-1] - idLoc[busInts[len(busInts)-1]] - busInts[len(busInts)-1]
	count := 0
	for true {
		if count % 100000000 == 1 {
			fmt.Println(possible)
		}
		count++
		possible += busInts[len(busInts)-1]
		for i := len(busInts)-1; i >= 0; i-- {
			if i != 0 {
				if busInts[i] - possible % busInts[i] != idLoc[busInts[i]] {
					//fmt.Printf("possible: %d, id: %d, loc: %d, remainder: %d\n", possible, busInts[i], idLoc[busInts[i]], busInts[i] - possible % busInts[i])
					break
				}
			} else {
				if busInts[i] - possible % busInts[i] != busInts[i] {
					//fmt.Println("breaking")
					break
				}
			}
			if i == 0 {
				fmt.Println(possible)
				return
			}
		}
	}*/
}

func findDistance(a int, b int, c int) (int, int) {
	first := 0
	second := 0
	x := 0
	for true {
		//fmt.Printf("%d, %d, %d, %d, %d\n", a, b, c, mod(a*x, b), a*x)
		if first != 0 && mod(a*x, b) == c {
			second = x
			return first, second - first
		}
		if first == 0 && mod(a*x, b) == c {
			first = x
		}
		x++
	}
	return -1, -1
}

func lcm(a int, b int) int  {
	return a * b / gcd(a, b)
}

func gcd(a int, b int) int {
	//fmt.Printf("%d, %d\n", a, b)
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	r := mod(a, b)
	//fmt.Printf("Remainder: %d\n", r)
	return gcd(b, r)
}

func mod(a, b int) int {
    m := a % b
    if a < 0 && b < 0 {
        m -= b
    }
    if a < 0 && b > 0 {
        m += b
    }
    return m
}