package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(filename string) ([][2]int, [][]int) {
	var rules [][2]int
	var updates [][]int

	inputFile, _ := os.Open(filename)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			ruleStrings := strings.Split(line, "|")
			firstNumber, _ := strconv.Atoi(strings.TrimSpace(ruleStrings[0]))
			secondNumber, _ := strconv.Atoi(strings.TrimSpace(ruleStrings[1]))
			ruleNumbers := [2]int{firstNumber, secondNumber}
			rules = append(rules, ruleNumbers)
		} else if strings.Contains(line, ",") {
			pageNumbersStrings := strings.Split(line, ",")
			var pageNumbers []int
			for _, numberString := range pageNumbersStrings {
				number, _ := strconv.Atoi(strings.TrimSpace(numberString))
				pageNumbers = append(pageNumbers, number)
			}
			updates = append(updates, pageNumbers)
		}
	}
	return rules, updates
}

func testRule(rule [2]int, pages []int) bool {
	for i, page := range pages {
		if page == rule[0] {
			if slices.Contains(pages[:i], rule[1]) {
				return false
			}
		} else if page == rule[1] {
			if slices.Contains(pages[i:], rule[0]) {
				return false
			}
		}
	}
	return true
}

func testUpdate(rules [][2]int, update []int) bool {
	for _, rule := range rules {
		if !testRule(rule, update) {
			return false
		}
	}
	return true
}

func main() {
	rules, updates := parseInput("input.txt")
	sumOfMiddles := 0
	for _, update := range updates {
		if testUpdate(rules, update) {
			sumOfMiddles += update[len(update)/2]
		}
	}
	fmt.Println(sumOfMiddles)
}
