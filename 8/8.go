package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type inst struct {
	op string
	val int
	run bool
}

func main() {
	file, _ := os.Open("8.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var prog []inst = nil

	for scanner.Scan() {
		line := scanner.Text()
		val, _ := strconv.Atoi(line[4:])
		prog = append(prog, inst{line[0:3], val, false})
	}
	cpy := make([]inst, len(prog))
	copy(cpy, prog)
	fmt.Println(run(prog, 0))
	
	for i := range(cpy) {
		copy2 := make([]inst, len(cpy))
		copy(copy2, cpy)
		if copy2[i].op == "nop" {
			copy2[i].op = "jmp"
			copy3 := make([]inst, len(copy2))
			copy(copy3, copy2)
			if test(copy2, 0) {
				fmt.Println(run(copy3, 0))
				break
			}
		}
		if copy2[i].op == "jmp" {
			copy2[i].op = "nop"
			copy3 := make([]inst, len(copy2))
			copy(copy3, copy2)
			if test(copy2, 0) {
				fmt.Println(run(copy3, 0))
				break
			}
		}
		
	}
}

func run(prog []inst, pc int) int {
	if pc == len(prog) {
		return 0
	}
	if prog[pc].run {
		return 0
	}
	prog[pc].run = true
	if prog[pc].op == "nop" {
		return run(prog, pc+1)
	}
	if prog[pc].op == "acc" {
		return prog[pc].val + run(prog, pc+1)
	}
	if prog[pc].op == "jmp" {
		return run(prog, pc+prog[pc].val)
	}
	return 0
}

func test(prog []inst, pc int) bool {
	if pc == len(prog) {
		return true
	}
	if prog[pc].run {
		return false
	}
	prog[pc].run = true
	if prog[pc].op == "nop" {
		return test(prog, pc+1)
	}
	if prog[pc].op == "acc" {
		return test(prog, pc+1)
	}
	if prog[pc].op == "jmp" {
		return test(prog, pc+prog[pc].val)
	}
	return false
}