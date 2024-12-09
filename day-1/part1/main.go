package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day-1/part1/file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var left []int
	var right []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Fields(line)

		num1, _ := strconv.Atoi(data[0])
		num2, _ := strconv.Atoi(data[1])

		left = append(left, num1)
		right = append(right, num2)
	}

	sort.Ints(left)
	sort.Ints(right)

	var result int
	for i := range left {
		if left[i] > right[i] {
			result += left[i] - right[i]
		} else {
			result += right[i] - left[i]
		}
	}

	fmt.Println(result)

}
