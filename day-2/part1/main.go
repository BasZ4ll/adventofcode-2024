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
	file, err := os.Open("./day-2/part1/file.txt")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var num int
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)
		fmt.Println("data", data)

		if calculateDifference(data) == 1 {
			num++
		}
	}
	fmt.Println(num)

}

func calculateDifference(data []string) int {

	var previous int
	isFirst := true

	var isIncreasing bool
	var directionSet bool

	var diff int
	for _, v := range data {
		current, _ := strconv.Atoi(v)

		if isFirst {
			previous = current
			isFirst = false
			continue
		}

		diff = current - previous

		if diff < -3 || diff > 3 || diff == 0 {
			fmt.Printf("Unsafe: (diff: %d)\n", diff)
			return 0
		}
		if !directionSet {
			isIncreasing = diff > 0
			directionSet = true
		} else {
			if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
				fmt.Printf("Unsafe: \n")
				return 0
			}
		}
		previous = current

	}
	return 1
}
