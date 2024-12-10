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
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var safeReports int
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Fields(line)

		if isSafe(data) {
			safeReports++
		} else if canBeSafeByRemovingOne(data) {
			safeReports++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeReports)
}

func isSafe(data []string) bool {
	var previous int
	isFirst := true
	directionSet := false
	isIncreasing := false

	for _, v := range data {
		current, _ := strconv.Atoi(v)

		if isFirst {
			previous = current
			isFirst = false
			continue
		}

		diff := current - previous

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if !directionSet {
			isIncreasing = diff > 0
			directionSet = true
		} else {
			if (isIncreasing && diff < 0) || (!isIncreasing && diff > 0) {
				return false
			}
		}

		previous = current
	}

	return true
}

// canBeSafeByRemovingOne checks if removing one number from the report can make it safe
func canBeSafeByRemovingOne(data []string) bool {
	for i := range data {
		// Create a copy of the data excluding the current index
		newData := append([]string{}, data[:i]...)
		newData = append(newData, data[i+1:]...)

		// Check if the modified data is safe
		if isSafe(newData) {
			return true
		}
	}
	return false
}
