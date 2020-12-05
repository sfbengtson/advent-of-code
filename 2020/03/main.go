package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//oneStar()
	twoStar()
}

func twoStar() {
	a := slope(1, 1)
	b := slope(3, 1)
	c := slope(5, 1)
	d := slope(7, 1)
	e := slope(1, 2)

	fmt.Println("Trees hit multiplied together is:", a*b*c*d*e)
}

func slope(r int, d int) int {
	arr := prepPuzzle("input.txt")
	pos, count := 0, 0
	//down 1
	for i := 0; i < len(arr); i += d {
		if arr[i][pos] == '#' {
			count++
		}
		pos += r
		if pos >= len(arr[0]) {
			//loop around to the beginning.
			pos = pos - len(arr[0])
		}
	}
	fmt.Println("Number of trees hit going right ", r, " and down", d, "is ", count)
	return count
}

func prepPuzzle(f string) []string {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var arr []string

	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}
