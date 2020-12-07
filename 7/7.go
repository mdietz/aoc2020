package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("7.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var bm map[string]map[string]int = make(map[string]map[string]int)

	for scanner.Scan() {
		line := scanner.Text()
		first := strings.Split(line, " bags contain ")[0]
		second := strings.Split(line, " bags contain ")[1]
		re := regexp.MustCompile(`\s*bag[s]*[,.]*\s*`)
		bags := re.Split(second, -1)
		fmt.Println(bags)
		var bagmap map[string]int = make(map[string]int)
		for _, e := range(bags) {
			if len(e) > 0 {
				if e[:2] == "no" {
					break
				}
				val, _ := strconv.Atoi(e[:1])
				bagmap[e[2:]] = val
			}
		}
		bm[first] = bagmap
	}
	fmt.Println(bm)
	var containsGold []string = nil
	containsGold = append(containsGold, "shiny gold")
	lenBags := 0
	for true {
		for k1, v := range(bm) {
			for k, _ := range(v) {
				found := contains(containsGold, k) && !contains(containsGold, k1)
				if found {
					containsGold = append(containsGold, k1)
				}
			}
		}
		if lenBags == len(containsGold) {
			break
		}
		lenBags = len(containsGold)
	}
	fmt.Println(containsGold)
	fmt.Println(len(containsGold)-1)

	var ws []string = nil
	ws = append(ws, "shiny gold")
	count := 0
	for len(ws) != 0 {
		curr := ws[0]
		ws = ws[1:]
		for k, v := range(bm[curr]) { 
			for i := 0; i < v; i++ {
				ws = append(ws, k)
				count++
			}	
		}
	}
	fmt.Println(count)
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}