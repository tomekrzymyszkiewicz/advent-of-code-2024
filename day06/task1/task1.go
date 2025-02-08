package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Guard struct {
	i    int
	j    int
	dirI int
	dirJ int
}

var (
	mapWidth  int
	mapHeight int
)

func parseInput(filename string) (*[]bool, *Guard) {
	inputFile, _ := os.Open(filename)
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	var lines []string = make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	mapWidth = len(lines[0])
	mapHeight = len(lines)
	var obstaclesMap []bool = make([]bool, len(lines)*mapWidth)
	var guardI, guardJ int
	for i, row := range lines {
		for j, point := range row {
			switch point {
			case '#':
				obstaclesMap[mapWidth*i+j] = true
			case '^':
				guardI = i
				guardJ = j
			default:
				obstaclesMap[mapWidth*i+j] = false
			}
		}
	}
	guard := &Guard{i: guardI, j: guardJ, dirI: -1, dirJ: 0}
	return &obstaclesMap, guard
}

func move(guardMap *[]bool, guard *Guard) {
	switch (*guardMap)[(guard.i+guard.dirI)*mapWidth+guard.j+guard.dirJ] {
	case true:
		switch {
		case guard.dirI == -1 && guard.dirJ == 0: //up
			guard.dirI = 0
			guard.dirJ = 1
		case guard.dirI == 0 && guard.dirJ == 1: //right
			guard.dirI = 1
			guard.dirJ = 0
		case guard.dirI == 1 && guard.dirJ == 0: //down
			guard.dirI = 0
			guard.dirJ = -1
		case guard.dirI == 0 && guard.dirJ == -1: //left
			guard.dirI = -1
			guard.dirJ = 0
		}
	case false:
		guard.i += guard.dirI
		guard.j += guard.dirJ
	}
}

func printMap(guardMap *[]bool, guard *Guard, visited *[]bool) {
	for i := range mapHeight {
		for j := range mapWidth {
			switch {
			case guard.i == i && guard.j == j:
				switch {
				case guard.dirI == -1 && guard.dirJ == 0: //up
					fmt.Print("^")
				case guard.dirI == 0 && guard.dirJ == 1: //right
					fmt.Print(">")
				case guard.dirI == 1 && guard.dirJ == 0: //down
					fmt.Print("v")
				case guard.dirI == 0 && guard.dirJ == -1: //left
					fmt.Print("<")
				}
			case (*guardMap)[i*mapWidth+j]:
				fmt.Print("#")
			case (*visited)[i*mapWidth+j]:
				fmt.Print("X")
			case !(*guardMap)[i*mapWidth+j]:
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	guardMap, guard := parseInput("input.txt")
	var visited []bool = make([]bool, mapWidth*mapHeight)
	for i := range visited {
		visited[i] = false
	}
	visited[guard.i*mapWidth+guard.j] = true
	for {
		move(guardMap, guard)
		visited[guard.i*mapWidth+guard.j] = true
		// cmd := exec.Command("clear")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		// printMap(guardMap, guard, &visited)
		time.Sleep(1 * time.Millisecond)
		if guard.i+guard.dirI < 0 || guard.i+guard.dirI >= mapWidth || guard.j+guard.dirJ < 0 || guard.j+guard.dirJ >= mapHeight {
			break
		}
	}
	numberOVisited := 0
	for i := range visited {
		if visited[i] {
			numberOVisited++
		}
	}
	fmt.Println(numberOVisited)
}
