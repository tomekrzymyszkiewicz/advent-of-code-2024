package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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

func testRule(rule [2]int, pages []int) (bool, int, int) {
	for i, page := range pages {
		if page == rule[0] {
			if slices.Contains(pages[:i], rule[1]) {
				j := slices.Index(pages[:], rule[1])
				return false, i, j
			}
		} else if page == rule[1] {
			if slices.Contains(pages[i:], rule[0]) {
				j := slices.Index(pages[:], rule[0])
				return false, i, j
			}
		}
	}
	return true, -1, -1
}

func testUpdate(rules [][2]int, update []int) (bool, []int) {
	hadCorrectOrder := true
	allRulesMatching := true
	for {
		allRulesMatching = true
		for _, rule := range rules {
			testResult, i, j := testRule(rule, update)
			if !testResult {
				swap := reflect.Swapper(update)
				swap(i, j)
				hadCorrectOrder = false
				allRulesMatching = false
			}
		}
		if allRulesMatching {
			break
		}
	}
	return hadCorrectOrder, update
}

func main() {
	rules, updates := parseInput("input.txt")
	sumOfMiddles := 0
	for _, update := range updates {
		result, pages := testUpdate(rules, update)
		if !result {
			sumOfMiddles += pages[len(pages)/2]
		}
	}
	fmt.Println(sumOfMiddles)
}
