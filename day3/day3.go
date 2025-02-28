package day3

import (
	"fmt"
	"regexp"
	"sort"
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

type Match struct {
	Type  string // Identifies which regex matched
	Value string // The actual matched string
	Index int    // The position in the original string
}

func Day3Part2(content string) {

	patterns := map[string]*regexp.Regexp{
		"mul":   regexp.MustCompile(`mul\((\d+),(\d+)\)`),
		"do":    regexp.MustCompile(`do\(\)`),
		"don't": regexp.MustCompile(`don\'t\(\)`),
	}

	var matches []Match

	// Find matches for each pattern
	for key, re := range patterns {
		locs := re.FindAllStringIndex(content, -1)
		for _, loc := range locs {
			matches = append(matches, Match{
				Type:  key,
				Value: content[loc[0]:loc[1]], // Extract matched substring
				Index: loc[0],                 // Start index
			})
		}
	}

	// Sort matches by order in the string
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Index < matches[j].Index
	})

	includeMul := true
	runningTotal := 0

	for _, match := range matches {
		switch match.Type {
		case "mul":
			if includeMul {
				nums := extractNumbers(match.Value)
				sum := nums[0] * nums[1]
				runningTotal += sum
				fmt.Printf("✅ Including %s -> +%d (Total: %d)\n", match.Value, sum, runningTotal)
			} else {
				fmt.Printf("❌ Excluding %s\n", match.Value)
			}
		case "don't":
			includeMul = false
			fmt.Println("🚫 Encountered don't(), excluding future mul(#,#)")
		case "do":
			includeMul = true
			fmt.Println("✅ Encountered do(), including future mul(#,#)")
		}
	}

	fmt.Println("Final Total:", runningTotal)
}

func extractNumbers(input string) []int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 3 {
		return []int{0, 0} // Fallback
	}
	n1, _ := strconv.Atoi(matches[1])
	n2, _ := strconv.Atoi(matches[2])
	return []int{n1, n2}
}
