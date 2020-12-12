package main

import (
	"../util"
	"fmt"
)

func main() {
	boatDir := 'E'
	x := 0
	y := 0
	wX := 0
	wY := 0
	inputs := util.GetInput("12.in")
	fmt.Println(inputs)
	for _, e := range(inputs) {
		if e[0] == 'F' {
			x, y = moveBoat(x, y, boatDir, util.Atoi(e[1:]))
			continue
		}
		if e[0] == 'R' {
			deg := util.Atoi(e[1:])
			boatDir = turnBoat(boatDir, deg)
			continue
		}
		if e[0] == 'L' {
			deg := util.Atoi(e[1:])
			boatDir = turnBoat(boatDir, -deg)
			continue
		}
		if e[0] == 'N' || e[0] == 'S' || e[0] == 'E' || e[0] == 'W' {
			x, y = moveBoat(x, y, rune(e[0]), util.Atoi(e[1:]))
			continue
		}
	}
	fmt.Println(Abs(x)+Abs(y))

	fmt.Println("=====================")

	boatDir = 'E'
	x = 0
	y = 0
	wX = 10
	wY = 1
	for _, e := range(inputs) {
		if e[0] == 'F' {
			x, y, wX, wY = moveBoat2(x, y, wX, wY, util.Atoi(e[1:]))
			continue
		}
		if e[0] == 'R' {
			deg := util.Atoi(e[1:])
			wX, wY = turnBoat2(wX, wY, deg)
			continue
		}
		if e[0] == 'L' {
			deg := util.Atoi(e[1:])
			wX, wY = turnBoat2(wX, wY, -deg)
			continue
		}
		if e[0] == 'N' || e[0] == 'S' || e[0] == 'E' || e[0] == 'W' {
			wX, wY = moveWaypoint(wX, wY, rune(e[0]), util.Atoi(e[1:]))
			continue
		}
	}
	fmt.Println(Abs(x)+Abs(y))
}

func moveWaypoint(wX int, wY int, d rune, l int) (int, int) {
	switch d {
	case 'N':
		return wX, wY + l
	case 'S':
		return wX, wY - l
	case 'E':
		return wX + l, wY
	case 'W':
		return wX - l, wY
	}
	return -1, -1


}

func moveBoat2(x int, y int, wX int, wY int, n int) (int, int, int, int) {
	for i := 0; i < n; i++ {
		x += wX
		y += wY
	}
	return x, y, wX, wY
}

func turnBoat2(wX int, wY int, deg int) (int, int) {
	if deg < 0 {
		deg += 360
	}
	if deg == 90 {
		return wY, -wX
	}
	if deg == 180 {
		return -wX, -wY
	}
	if deg == 270 {
		return -wY, wX
	}
	return -1, -1
}

func moveBoat(x int, y int, d rune, l int) (int, int) {
	switch d {
	case 'E':
		x += l
	case 'W':
		x -= l
	case 'N':
		y += l
	case 'S':
		y -= l
	}
	return x, y
}

func turnBoat(d rune, deg int) rune {
	if deg < 0 {
			deg += 360
		}
	switch d {
	case 'E':
		if deg == 0 {
			return 'E'
		}
		if deg == 90 {
			return 'S'
		}
		if deg == 180 {
			return 'W'
		}
		if deg == 270 {
			return 'N'
		}
	case 'W':
		if deg == 0 {
			return 'W'
		}
		if deg == 90 {
			return 'N'
		}
		if deg == 180 {
			return 'E'
		}
		if deg == 270 {
			return 'S'
		}
	case 'N':
		if deg == 0 {
			return 'N'
		}
		if deg == 90 {
			return 'E'
		}
		if deg == 180 {
			return 'S'
		}
		if deg == 270 {
			return 'W'
		}
	case 'S':
		if deg == 0 {
			return 'S'
		}
		if deg == 90 {
			return 'W'
		}
		if deg == 180 {
			return 'N'
		}
		if deg == 270 {
			return 'E'
		}
	}
	return 'A'
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}