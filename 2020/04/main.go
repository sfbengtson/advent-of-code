package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	oneStar()
	//twoStar()
}

func oneStar() {
	//prepPuzzle
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	mapArr := []map[string]string{}
	curPas := map[string]string{}
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			mapArr = append(mapArr, curPas)
			curPas = map[string]string{}
		} else {
			arr := strings.Split(scanner.Text(), " ")
			for _, s := range arr {
				kv := strings.Split(s, ":")
				curPas[kv[0]] = kv[1]
			}
		}
	}
	//add last one
	mapArr = append(mapArr, curPas)

	//ok we got some stuff, now to check..........
	a := checkList(mapArr)
	fmt.Println("number of valid passports is: ", a)
}

func checkList(input []map[string]string) int {
	count := 0
	list := []string{"cid", "byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for a, passport := range input {
		fmt.Println(passport)
		valid := true
		for _, entry := range list {
			val, ok := passport[entry]
			if !validateEntry(entry, val, ok) && entry != "cid" {
				fmt.Println("passport", a, "has invalid field ", entry)
				valid = false
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func validateEntry(e string, val string, ok bool) bool {
	if !ok {
		return false
	}
	switch e {
	case "cid":
		//lol we don't check this nonsense
		return true
	case "byr":
		return intRangeValidate(val, 1920, 2002)
	case "iyr":
		return intRangeValidate(val, 2010, 2020)
	case "eyr":
		return intRangeValidate(val, 2020, 2030)
	case "hgt":
		return heightValidate(val)
	case "hcl":
		matched, err := regexp.Match(`^#{1}[0-9a-f]{6}$`, []byte(val))
		if err != nil || !matched {
			return false
		}
		return true
	case "ecl":
		var colList = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		if contains(colList, val) {
			return true
		}
		return false
	case "pid":
		matched, err := regexp.Match(`^([0-9]{9})$`, []byte(val))
		if err != nil || !matched {
			return false
		}
		return true
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func heightValidate(s string) bool {
	// -3 = -2 (2 from end) - 1 (arr starts at 0)
	if len(s) < 3 {
		return false
	}
	if s[len(s)-2:] == "in" {
		fmt.Println(s[:len(s)-2])
		return intRangeValidate(s[:len(s)-2], 59, 76)
	} else if s[len(s)-2:] == "cm" {
		//cm
		return intRangeValidate(s[:len(s)-2], 150, 193)
	}
	return false
}

func intRangeValidate(val string, min int, max int) bool {
	d, err := strconv.Atoi(val)
	if err != nil {
		return false
	}
	if d >= min && d <= max {
		return true
	}
	return false
}
