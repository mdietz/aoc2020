package main

import (
	"../util"
	"fmt"
)

func main () {
	inputs := util.GetInput("17.in")
	size := 30
	var dim [][][][]bool = make([][][][]bool, size)
	start := size/2-1
	for h := 0; h < size; h++ {
		dim[h] = make([][][]bool, size)
		for i := 0; i < size; i++ {
			dim[h][i] = make([][]bool, size)
			for j := 0; j < size; j++ {
				dim[h][i][j] = make([]bool, size)
				for k := 0; k < size; k++ {
					dim[h][i][j][k] = false
				}
			}
		}
	}
	w := start
	z := start
	y := start
	for _, e := range(inputs) {
		x := start
		for _, i := range(e) {
			if i == '#' {
				dim[w][z][y][x] = true
			} else {
				dim[w][z][y][x] = false
			}
			x++
		}
		y++
	}
	next := dim
	for r := 0; r < 6; r++ {
		next = RunCycle2(next)
		
	}
	tot := 0
	for h := 0; h < size; h++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					if next[h][i][j][k] == true {
						tot++
					}
				}
			}
		}
	}
	fmt.Println(tot)
}

func CheckNeighboors2(in [][][][]bool, w int, z int, y int, x int) int {
	max := len(in)-1
	tot := 0
	for h := w-1; h <= w+1; h++ {
		if h < 0 || h > max {
			continue
		}
		for i := z-1; i <= z+1; i++ {
			if i < 0 || i > max {
				continue
			}
			for j := y-1; j <= y+1; j++ {
				if j < 0 || j > max {
					continue
				}
				for k := x-1; k <= x+1; k++ {
					if k < 0 || k > max {
						continue
					}
					if h == w && i == z && j == y && k == x {
						continue
					}
					if in[h][i][j][k] == true {
						tot++
					}
				}
			}
		}
	}
	return tot
}

func CheckNeighboors(in [][][]bool, z int, y int, x int) int {
	max := len(in)-1
	tot := 0
	for i := z-1; i <= z+1; i++ {
		if i < 0 || i > max {
			continue
		}
		for j := y-1; j <= y+1; j++ {
			if j < 0 || j > max {
				continue
			}
			for k := x-1; k <= x+1; k++ {
				if k < 0 || k > max {
					continue
				}
				if i == z && j == y && k == x {
					continue
				}
				if in[i][j][k] == true {
					tot++
				}
			}
		}
	}
	return tot
}

func RunCycle2(in [][][][]bool) [][][][]bool {
	size := len(in)
	var out [][][][]bool = make([][][][]bool, size)
	for h := 0; h < size; h++ {
		out[h] = make([][][]bool, size)
		for i := 0; i < size; i++ {
			out[h][i] = make([][]bool, size)
			for j := 0; j < size; j++ {
				out[h][i][j] = make([]bool, size)
				for k := 0; k < size; k++ {
					out[h][i][j][k] = false
				}
			}
		}
	}
	for h := 0; h < size; h++ {
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				for k := 0; k < size; k++ {
					num := CheckNeighboors2(in, h, i, j, k)
					if in[h][i][j][k] {
						if num == 2 || num == 3 {
							out[h][i][j][k] = true
						} else {
							out[h][i][j][k] = false
						}
					} else {
						if num == 3 {
							out[h][i][j][k] = true
						} else {
							out[h][i][j][k] = false
						}
					}

				}
			}
		}
	}
	return out
}

func RunCycle(in [][][]bool) [][][]bool {
	size := len(in)
	var out [][][]bool = make([][][]bool, size)
	for i := 0; i < size; i++ {
		out[i] = make([][]bool, size)
		for j := 0; j < size; j++ {
			out[i][j] = make([]bool, size)
			for k := 0; k < size; k++ {
				out[i][j][k] = false
			}
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				num := CheckNeighboors(in, i, j, k)
				if in[i][j][k] {
					if num == 2 || num == 3 {
						out[i][j][k] = true
					} else {
						out[i][j][k] = false
					}
				} else {
					if num == 3 {
						out[i][j][k] = true
					} else {
						out[i][j][k] = false
					}
				}

			}
		}
	}
	return out
}

func PrettyPrint2(in [][][][]bool) string{
	out := ""
	size := len(in)
	start := -size/2
	w := start
	z := start
	for _, d := range(in) {
		out += fmt.Sprintf("w=%d\n", w)
		for _, a := range(d) {
			out += fmt.Sprintf("z=%d\n", z)
			for _, b := range(a) {
				for _, c := range(b) {
					if c == true {
						out += fmt.Sprintf("#")
					} else {
						out += fmt.Sprintf(".")
					}
				}
				out += fmt.Sprintf("\n")
			}
			out += fmt.Sprintf("\n")
			z++
		}
		w++
	}
	return out
}

func PrettyPrint(in [][][]bool) string{
	out := ""
	size := len(in)
	start := -size/2
	z := start
	for _, a := range(in) {
		out += fmt.Sprintf("z=%d\n", z)
		for _, b := range(a) {
			for _, c := range(b) {
				if c == true {
					out += fmt.Sprintf("#")
				} else {
					out += fmt.Sprintf(".")
				}
			}
			out += fmt.Sprintf("\n")
		}
		out += fmt.Sprintf("\n")
		z++
	}
	return out
}