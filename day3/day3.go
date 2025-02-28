package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3Part1(content string) {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(content, -1)

	var total int
	for _, match := range matches {
		if len(match) == 3 { // match[0] is full match, match[1] and match[2] are numbers
			fmt.Printf("Found: mul(%s, %s)\n", match[1], match[2])
			num1, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Error converting number:", err)
				continue
			}
			num2, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error converting number:", err)
				continue
			}
			total += num1 * num2
		}
	}
	fmt.Println(total)
}
