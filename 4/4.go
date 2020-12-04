package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
)

var reqFields []string = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	file, _ := os.Open("4.in")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var passports []map[string]string = nil

	var cur map[string]string = make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println(cur)
			passports = append(passports, cur)
			cur = make(map[string]string)
		} else {
			ls := strings.Split(line, " ")
			for _, e := range(ls) {
				pair := strings.Split(e, ":")
				cur[pair[0]] = pair[1]
			}
		}
	}

	valid := 0
	for _, p := range(passports) {
		fmt.Println(p)
		if isValid(p) {
			valid++
		}
	}
	fmt.Printf("Valid: %d\n", valid)
}

func isValid(p map[string]string) bool {
	for _, f := range(reqFields) {
		val, prs := p[f]
		fmt.Println(f)
		if !prs {
			return false
		}
		if f == "byr" {
			i, _ := strconv.Atoi(val)
			if i < 1920 || i > 2002 {
				return false
			}
		}
		if f == "iyr" {
			i, _ := strconv.Atoi(val)
			if i < 2010 || i > 2020 {
				return false
			}
		}
		if f == "eyr" {
			i, _ := strconv.Atoi(val)
			if i < 2020 || i > 2030 {
				return false
			}
		}
		if f == "hgt" {
			units := val[len(val)-2:]
			if units != "cm" && units != "in" {
				return false
			}
			i, _ := strconv.Atoi(val[:len(val)-2])
			if units == "cm" {
				if i < 150 || i > 193 {
					return false
				}
			}
			if units == "in" {
				if i < 59 || i > 76 {
					return false
				}
			}
		}
		if f == "hcl" {
			match, _ := regexp.MatchString("^#[0-9a-f]{6}$", val)
			if match == false {
				return false
			}
		}
		if f == "ecl" {
			if val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth" {
				continue
			} else {
				return false
			}
		}
		if f == "pid" {
			match, _ := regexp.MatchString("^[0-9]{9}$", val)
			if match == false {
				return false
			}
		}
	}
	return true
}