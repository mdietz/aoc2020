package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, _ := os.Open("5.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var seatList []int = nil

	maxSeat := 0
	for scanner.Scan() {
		var row [7]bool
		var col [3]bool
		line := scanner.Text()
		for i, e := range(line[:7]) {
			if e == 'B' {
				row[i] = true
			} else {
				row[i] = false
			}
		}
		for i, e := range(line[7:]) {
			if e == 'R' {
				col[i] = true
			} else {
				col[i] = false
			}
		}

		if seatNum(boolArrToInt(row[:]), boolArrToInt(col[:])) > maxSeat {
			maxSeat = seatNum(boolArrToInt(row[:]), boolArrToInt(col[:]))
		}

		seatList = append(seatList, seatNum(boolArrToInt(row[:]), boolArrToInt(col[:])))

	}
	
	sort.Ints(seatList)
	lastSeen := seatList[0]
	for _, e := range(seatList[1:]) {
		fmt.Println(lastSeen, e)
		if e - lastSeen == 1 {
			lastSeen = e
		} else {
			break
		}
	}
	fmt.Println(lastSeen + 1)

}

func seatNum(row int, col int) int {
	return row*8+col;
}

func boolArrToInt(b []bool) int {
	o := 0
	for i := 0; i < len(b)-1; i++ {
		if b[i] {
			o++
		}
		o = o << 1
	}
	if b[len(b)-1] {
		o++
	}
	return o
}