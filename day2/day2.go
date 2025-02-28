package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func isSafeLevel(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if diff > 0 {
			isDecreasing = false
		} else if diff < 0 {
			isIncreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func Day2Part2(content string) {
	rows := strings.Split(strings.TrimSpace(content), "\n")
	var safeReports int
	var fixedReports int

	for _, row := range rows {
		levels := strings.Fields(row)

		if len(levels) == 0 {
			continue
		}

		var numbers []int
		for _, level := range levels {
			num, err := strconv.Atoi(level)
			if err != nil {
				fmt.Println("Error converting number:", err)
				continue
			}
			numbers = append(numbers, num)
		}

		if isSafeLevel(numbers) {
			safeReports++
			fmt.Println("✅ Safe report:", numbers)
			continue
		}

		// Try removing each number and check if it becomes safe
		for i := 0; i < len(numbers); i++ {
			tempNumbers := append([]int{}, numbers[:i]...)
			if i+1 < len(numbers) {
				tempNumbers = append(tempNumbers, numbers[i+1:]...)
			}

			if isSafeLevel(tempNumbers) {
				fixedReports++
				break
			}
		}
	}

	fmt.Println("Total Safe Reports:", safeReports+fixedReports)
}
