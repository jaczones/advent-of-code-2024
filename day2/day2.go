package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2Part1(content string) {
	rows := strings.Split(strings.TrimSpace(content), "\n")
	var safeReports int

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

		isIncreasing := true
		isDecreasing := true
		isSafeLevel := true

		for i := 1; i < len(numbers); i++ {
			diff := numbers[i] - numbers[i-1]

			if diff < -3 || diff > 3 || diff == 0 {
				fmt.Printf("❌ Unsafe diff of %d between %d and %d\n", diff, numbers[i-1], numbers[i])
				isSafeLevel = false
				break
			}

			// Track whether the report is increasing or decreasing
			if diff > 0 {
				isDecreasing = false
			} else if diff < 0 {
				isIncreasing = false
			}

			// If both increasing and decreasing are false, the report is not safe
			if !isIncreasing && !isDecreasing {
				fmt.Println("❌ Unsafe report: numbers do not consistently increase/decrease")
				isSafeLevel = false
				break
			}
		}

		if isSafeLevel {
			fmt.Println("✅ Safe report:", numbers)
			safeReports++
		} else {
			fmt.Println("❌ Unsafe report:", numbers)
		}
	}

	fmt.Println("Total Safe Reports:", safeReports)
}
