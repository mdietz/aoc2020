package main

import (
	"../util"
	"fmt"
)

type seatType string

const (
	Empty    seatType = "Empty"
	Floor             = "Floor"
	Occupied          = "Occupied"
)

type direction string

const (
	Above      direction = "Above"
	Below                = "Below"
	Left                 = "Left"
	Right                = "Right"
	AboveLeft            = "AboveLeft"
	AboveRight           = "AboveRight"
	BelowLeft            = "BelowLeft"
	BelowRight           = "BelowRight"
)

type seat struct {
	seatType                                                                seatType
	above, below, left, right, aboveLeft, aboveRight, belowLeft, belowRight *seat
}

func main() {
	inputs := util.GetInput("11.in")
	seatMap := makeMap(inputs)
	for true {
		changes := 0
		dup := makeMap(toString(seatMap))
		for i, r := range seatMap {
			for j, c := range r {
				if becomeOccupied2(c) {
					dup[i][j].seatType = Occupied
					changes++
					continue
				}
				if becomeEmpty2(c) {
					dup[i][j].seatType = Empty
					changes++
					continue
				}
			}
		}
		seatMap = dup
		if changes == 0 {
			break
		}
	}

	occupied := 0
	for _, r := range seatMap {
		for _, c := range r {
			if c.seatType == Occupied {
				occupied++
			}
		}
	}
	fmt.Println(occupied)
}

func firstSeen(s *seat, dir direction) seatType {
	if s == nil {
		return Empty
	}
	if s.seatType == Floor {
		switch dir {
		case Above:
			return firstSeen(s.above, dir)
		case Below:
			return firstSeen(s.below, dir)
		case Left:
			return firstSeen(s.left, dir)
		case Right:
			return firstSeen(s.right, dir)
		case AboveLeft:
			return firstSeen(s.aboveLeft, dir)
		case AboveRight:
			return firstSeen(s.aboveRight, dir)
		case BelowLeft:
			return firstSeen(s.belowLeft, dir)
		case BelowRight:
			return firstSeen(s.belowRight, dir)
		}
	}
	return s.seatType
}

func becomeOccupied2(s *seat) bool {
	if s.seatType == Empty {
		return firstSeen(s.above, Above) == Empty && firstSeen(s.below, Below) == Empty && firstSeen(s.left, Left) == Empty && firstSeen(s.right, Right) == Empty && firstSeen(s.aboveLeft, AboveLeft) == Empty && firstSeen(s.aboveRight, AboveRight) == Empty && firstSeen(s.belowLeft, BelowLeft) == Empty && firstSeen(s.belowRight, BelowRight) == Empty
	}
	return false
}

func becomeEmpty2(s *seat) bool {
	if s.seatType == Occupied {
		count := 0
		if firstSeen(s.above, Above) == Occupied {
			count++
		}
		if firstSeen(s.below, Below) == Occupied {
			count++
		}
		if firstSeen(s.left, Left) == Occupied {
			count++
		}
		if firstSeen(s.right, Right) == Occupied {
			count++
		}
		if firstSeen(s.aboveLeft, AboveLeft) == Occupied {
			count++
		}
		if firstSeen(s.aboveRight, AboveRight) == Occupied {
			count++
		}
		if firstSeen(s.belowLeft, BelowLeft) == Occupied {
			count++
		}
		if firstSeen(s.belowRight, BelowRight) == Occupied {
			count++
		}
		return count >= 5
	}
	return false
}

func becomeOccupied(s *seat) bool {
	if s.seatType == Empty {
		return emptyOrFloor(s.above) && emptyOrFloor(s.below) && emptyOrFloor(s.left) && emptyOrFloor(s.right) && emptyOrFloor(s.aboveLeft) && emptyOrFloor(s.aboveRight) && emptyOrFloor(s.belowLeft) && emptyOrFloor(s.belowRight)
	}
	return false
}

func becomeEmpty(s *seat) bool {
	if s.seatType == Occupied {
		count := 0
		if isOccupied(s.above) {
			count++
		}
		if isOccupied(s.below) {
			count++
		}
		if isOccupied(s.left) {
			count++
		}
		if isOccupied(s.right) {
			count++
		}
		if isOccupied(s.aboveLeft) {
			count++
		}
		if isOccupied(s.aboveRight) {
			count++
		}
		if isOccupied(s.belowLeft) {
			count++
		}
		if isOccupied(s.belowRight) {
			count++
		}
		return count >= 4
	}
	return false
}

func isOccupied(s *seat) bool {
	if s == nil {
		return false
	}
	return s.seatType == Occupied
}

func emptyOrFloor(s *seat) bool {
	if s == nil {
		return true
	}
	return s.seatType == Empty || s.seatType == Floor
}

func prettyPrint(seatMap [][]*seat) {
	for _, r := range seatMap {
		for _, c := range r {
			if c.seatType == Floor {
				fmt.Print(".")
			}
			if c.seatType == Occupied {
				fmt.Print("#")
			}
			if c.seatType == Empty {
				fmt.Print("L")
			}
		}
		fmt.Print("\n")
	}
}

func toString(seatMap [][]*seat) []string {
	var out []string = nil
	for _, r := range seatMap {
		curr := ""
		for _, c := range r {
			if c.seatType == Floor {
				curr += "."
			}
			if c.seatType == Occupied {
				curr += "#"
			}
			if c.seatType == Empty {
				curr += "L"
			}
		}
		out = append(out, curr)
	}
	return out
}

func makeMap(inputs []string) [][]*seat {
	var seatMap [][]*seat = nil

	for _, e := range inputs {
		var curr []*seat = nil
		for range e {
			curr = append(curr, &seat{})
		}
		seatMap = append(seatMap, curr)
	}

	for i, e := range inputs {
		for j, c := range e {
			currSeat := seatMap[i][j]
			if i != 0 {
				currSeat.above = seatMap[i-1][j]
				if j != 0 {
					currSeat.aboveLeft = seatMap[i-1][j-1]
				}
				if j != len(inputs[0])-1 {
					currSeat.aboveRight = seatMap[i-1][j+1]
				}
			}
			if i != len(inputs)-1 {
				currSeat.below = seatMap[i+1][j]
				if j != 0 {
					currSeat.belowLeft = seatMap[i+1][j-1]
				}
				if j != len(inputs[0])-1 {
					currSeat.belowRight = seatMap[i+1][j+1]
				}
			}
			if j != 0 {
				currSeat.left = seatMap[i][j-1]
			}
			if j != len(inputs[0])-1 {
				currSeat.right = seatMap[i][j+1]
			}
			if c == '.' {
				currSeat.seatType = Floor
			}
			if c == 'L' {
				currSeat.seatType = Empty
			}
			if c == '#' {
				currSeat.seatType = Occupied
			}
		}
	}
	return seatMap
}
