package main

import (
	"../util"
	"fmt"
	"strings"
	"regexp"
)

type span struct {
	start int
	end int
}

type field struct {
	name string
	spans []span
}

func main() {
	inputs := util.GetInput("16.in")
	fieldRe := regexp.MustCompile(`([\w|\s]+):\s+(\d+)\-(\d+)\s+or\s+(\d+)\-(\d+)`)

	var fieldVals []field = nil
	var yourTicket []int
	var otherTickets [][]int = nil
	var nearByTicketStart int
	for i := 0; i < len(inputs); i++ {
		if inputs[i] == "" {
			continue
		} else if inputs[i] == "your ticket:" {
			i++
			yourTicket = parseTick(strings.Split(inputs[i], ","))
		} else if inputs[i] == "nearby tickets:" {
			nearByTicketStart = i+1
			break
		} else {
			fmt.Println(inputs[i])
			res := fieldRe.FindAllStringSubmatch(inputs[i], -1)[0]
			fieldName := res[1]
			start1 := util.Atoi(res[2])
			end1 := util.Atoi(res[3])
			start2 := util.Atoi(res[4])
			end2 := util.Atoi(res[5])
			fieldVals = append(fieldVals, field{fieldName, []span{span{start1, end1}, span{start2, end2}}})
		}
	}
	for i := nearByTicketStart; i < len(inputs); i++ {
		otherTickets = append(otherTickets, parseTick(strings.Split(inputs[i], ",")))
	}

	fmt.Println(fieldVals)
	fmt.Println(yourTicket)
	fmt.Println(otherTickets)

	err := 0
	var sanTickets [][]int = nil
	for _, e := range(otherTickets) {
		tickValid := true
		for _, comp := range(e) {
			valid := false
			ValidLoop:
			for _, val := range(fieldVals) {
				for _, s := range(val.spans) {
					if comp >= s.start && comp <= s.end {
						valid = true
						break ValidLoop
					}
				}
			}
			if !valid {
				//fmt.Println(comp)
				err += comp
				tickValid = false
				break
			}
		}
		if tickValid {
			sanTickets = append(sanTickets, e)
		}
	}
	sanTickets = append(sanTickets, yourTicket)
	fmt.Println(err)

	var locFieldMap map[int]map[string]bool = make(map[int]map[string]bool)
	for i, _ := range(yourTicket) {
		locFieldMap[i] =  make(map[string]bool)
		for _, e := range(fieldVals) {
			locFieldMap[i][e.name] = true
		}
	}

	fmt.Println(locFieldMap)
	fmt.Println(sanTickets)

	for _, e := range(sanTickets) {
		for i, comp := range(e) {
			for _, val := range(fieldVals) {
				valid := false
				for _, s := range(val.spans) {
					if comp >= s.start && comp <= s.end {
						valid = true
						break
					}
				}
				if !valid {
					locFieldMap[i][val.name] = false
				}
			}
		}
	}
	
	for k1, v1 := range(locFieldMap) {
		count := 0
		for _, v2 := range(v1) {
			if v2 == true {
				count++
				//fmt.Printf("%d: %d\n", k1, k2)
			}
		}
		fmt.Printf("%d: %d\n", k1, count)
	}

	var availName map[string]bool = make(map[string]bool)
	for _, e := range(fieldVals) {
		availName[e.name] = true
	}

	var availLoc map[int]bool = make(map[int]bool)
	for k, _ := range(locFieldMap) {
		availLoc[k] = true
	}

	var departureLocs []int = nil 
	for true {
		keepGoing := false
		for _, v := range(availLoc) {
			if v {
				keepGoing = true
			}
		}
		if !keepGoing {
			break
		}
		for k1, v1 := range(locFieldMap) {
			var possible map[string]bool = make(map[string]bool)
			for k2, v2 := range(v1) {
				if v2 == true {
					possible[k2] = true
				} else {
					possible[k2] = false
				}
			}
			numPossible := 0
			for k3, v3 := range(possible) {
				if v3 && availName[k3] {
					numPossible++
				}
			}
			if numPossible == 1 {
				for k4, v4 := range(possible) {
					if v4 && availName[k4] {
						fmt.Printf("%d: %s\n", k1, k4)
						if strings.Contains(k4, "departure") {
							departureLocs = append(departureLocs, k1)
						}
						availLoc[k1] = false
						availName[k4] = false
					}
				}
			}
		}
	}

	total := 1
	for _, e := range(departureLocs) {
		total *= yourTicket[e]
	}
	fmt.Println(total)
}

func parseTick(in []string) []int{
	var out []int = nil
	for _, e := range(in) {
		out = append(out, util.Atoi(e))
	}
	return out
}