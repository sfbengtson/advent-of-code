package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//oneStar()
	twoStar()
}

func oneStar() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		//rule [0] = amt, rule[1] = letter, rule[2] = pw
		rule := strings.Split(scanner.Text(), " ")
		if checkRuleOneStar(rule[0], rule[1], rule[2]) {
			count++
		}
	}
	fmt.Println("Number of valid passwords is: ", count)
}

func twoStar() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		//rule [0] = amt, rule[1] = letter, rule[2] = pw
		rule := strings.Split(scanner.Text(), " ")
		if checkRuleTwoStar(rule[0], rule[1], rule[2]) {
			count++
		}
	}
	fmt.Println("Number of valid passwords is: ", count)
}

func checkRuleOneStar(amt string, letters string, pw string) bool {
	letter := letters[0]
	amtRule := strings.Split(amt, "-")
	min, _ := strconv.Atoi(amtRule[0])
	max, _ := strconv.Atoi(amtRule[1])
	ct := strings.Count(pw, string(letter))
	if min <= ct && ct <= max {
		return true
	}
	return false
}

func checkRuleTwoStar(amt string, letters string, pw string) bool {
	letter := letters[0]
	amtRule := strings.Split(amt, "-")
	location1, _ := strconv.Atoi(amtRule[0])
	location1 = location1 - 1
	location2, _ := strconv.Atoi(amtRule[1])
	location2 = location2 - 1

	fmt.Println(pw, len(pw), location1, location2)
	if pw[location1] == letter {
		if pw[location2] == letter {
			return false
		}
		return true
	} else if pw[location2] == letter {
		return true
	}
	return false
}
