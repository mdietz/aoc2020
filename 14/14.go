package main

import (
	"../util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputs := util.GetInput("14.in")
	fmt.Println(inputs)
	var res map[int]string  = make(map[int]string)

	maskRe := regexp.MustCompile(`mask\s+=\s+([X|\d]+)`)
	re := regexp.MustCompile(`mem\[(\d+)\]\s+=\s+(\d+)`)
	bitmask := maskRe.FindAllStringSubmatch(inputs[0], -1)[0][1]
	for _, e := range(inputs[1:]) {
		if e[0:3] == "mem" {
			matches := re.FindAllStringSubmatch(e, -1)[0]
			s := fmt.Sprintf("%036s", strconv.FormatInt(int64(util.Atoi(matches[2])), 2))
			res[util.Atoi(matches[1])] = applyMask(s, bitmask)
		} else {
			bitmask = maskRe.FindAllStringSubmatch(e, -1)[0][1]
		}
	}

	sum := int64(0)
	for _, v := range(res) {
		i, _ := strconv.ParseInt(v, 2, 64)
		sum += i
	}
	fmt.Println(sum)

	fmt.Println("===============")

	type update struct {
		loc string
		val int64
	}

	var memUps []update = nil 

	bitmask = maskRe.FindAllStringSubmatch(inputs[0], -1)[0][1]
	for _, e := range(inputs[1:]) {
		if e[0:3] == "mem" {
			matches := re.FindAllStringSubmatch(e, -1)[0]
			s := fmt.Sprintf("%036s", strconv.FormatInt(int64(util.Atoi(matches[1])), 2))
			memUps = append(memUps, update{applyMask2(s, bitmask), int64(util.Atoi(matches[2]))})
		} else {
			bitmask = maskRe.FindAllStringSubmatch(e, -1)[0][1]
		}
	}

	sum = 0
	var runningMemSet map[int64]bool = make(map[int64]bool)
	for i := len(memUps)-1; i >= 0; i-- {
		var currMemSet map[int64]bool = make(map[int64]bool)
		curr := memUps[i]
		memSet(curr.loc, currMemSet)
		for k, _ := range(currMemSet) {
			if runningMemSet[k] == false {
				sum += curr.val
				runningMemSet[k] = true
			}
		}
	}
	fmt.Println(sum)
}

func memSet(s string, accum map[int64]bool) {
	if !strings.ContainsRune(s, 'X') {
		i, _ := strconv.ParseInt(s, 2, 64)
		accum[i] = true
		return
	}
	s0 := strings.Replace(s, "X", "0", 1)
	s1 := strings.Replace(s, "X", "1", 1)
	memSet(s0, accum)
	memSet(s1, accum)
}

func applyMask(in string, mask string) string {
	out := ""
	for i, e := range(mask) {
		if e != 'X' {
			out += string(e)
		} else {
			out += string(in[i])
		}
	}
	return out
}

func applyMask2(in string, mask string) string {
	out := ""
	for i, e := range(mask) {
		if e == '1' {
			out += "1"
		}
		if e == '0' {
			out += string(in[i])
		}
		if e == 'X' {
			out += "X"
		}
	}
	return out
}