package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//oneStar()
	twoStar()
}

func twoStar() {
	arr := prepProblem("input.txt")

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			for k := j; k < len(arr); k++ {
				if arr[i]+arr[j]+arr[k] == 2020 {
					fmt.Println("i is, ", arr[i], " and j is", arr[j], "and k is", arr[k])
					fmt.Println("Answer is; ", arr[i]*arr[j]*arr[k])
				}
			}
		}
	}
}

// func oneStar() {
// 	return
// 	fmt.Println("TEST")
// 	arr := prepProblem("input.txt")

// 	for i := 0; i < len(arr); i++ {
// 		for j := i; j < len(arr); j++ {
// 			if arr[i]+arr[j] == 2020 {
// 				fmt.Println("i is, ", arr[i], " and j is", arr[j])
// 				fmt.Println("Answer is; ", arr[i]*arr[j])
// 			}
// 		}
// 	}
// }

func prepProblem(f string) []int {
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var arr []int

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("error, argument", scanner.Text(), "is not a number!")
			os.Exit(1)
		}
		arr = append(arr, n)
	}
	fmt.Println(arr)
	return arr
}
