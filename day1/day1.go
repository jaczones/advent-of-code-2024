package day1

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day1Part1(content string) {
	rows := strings.Split(strings.TrimSpace(content), "\n")
	var columnA []int
	var columnB []int

	for _, row := range rows {
		columns := strings.Fields(row)

		if len(columns) == 2 {
			numA, _ := strconv.Atoi(columns[0])
			numB, _ := strconv.Atoi(columns[1])

			columnA = append(columnA, numA)
			columnB = append(columnB, numB)
		}
	}
	sort.Ints(columnA)
	sort.Ints(columnB)

	totalDifference := 0

	for i := 0; i < len(columnA); i++ {
		diff := int(math.Abs(float64(columnA[i] - columnB[i])))
		totalDifference += diff
		fmt.Printf("Row %d: |%d - %d| = %d\n", i+1, columnA[i], columnB[i], diff)
	}

	fmt.Println("Sorted Column A:", columnA)
	fmt.Println("Sorted Column B:", columnB)

	fmt.Println("Total Distance:", totalDifference)
}

func Day1Part2(content string) {
	rows := strings.Split(strings.TrimSpace(content), "\n")
	var columnA []int
	var columnB []int

	for _, row := range rows {
		columns := strings.Fields(row)

		if len(columns) == 2 {
			numA, _ := strconv.Atoi(columns[0])
			numB, _ := strconv.Atoi(columns[1])

			columnA = append(columnA, numA)
			columnB = append(columnB, numB)
		}
	}
	var simScore int
	for _, num := range columnA {
		var match int
		for _, numColB := range columnB {
			if num == numColB {
				match++
			}
		}
		if match > 0 {
			simScore += num * match
			fmt.Println("simScore = ", simScore)
		}
	}
	fmt.Println(simScore)
}
