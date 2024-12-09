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
	file, err := os.Open("./day-1/part2/file.txt")
	if err != nil {
		log.Println("Error: ", err)
	}
	defer file.Close()

	// var left []int
	var right []int

	scanner := bufio.NewScanner(file)
	leftExists := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()

		data := strings.Fields(line)

		num1, _ := strconv.Atoi(data[0])
		num2, _ := strconv.Atoi(data[1])

		leftExists[num1] = true
		// fmt.Println(leftExists)
		right = append(right, num2)
		// fmt.Println(right)
	}

	// count := make(map[int]int)

	// // fmt.Println(count)
	// for _, value := range right {
	// 	count[value]++
	// 	fmt.Println(count[value])
	// }

	// for num, c := range count {
	// 	if c > 1 {
	// 		fmt.Printf("ตัวเลข %d ซ้ำ %d ครั้ง\n", num, c)
	// 	}
	// }

	var result int
	for _, r := range right {
		if leftExists[r] {
			result += r
		}
	}
	fmt.Println(result)

}
