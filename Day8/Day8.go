package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Point struct {
	X, Y int
}

func bfs(grid [][]string, condition func(string) bool) int {
	rows, cols := len(grid), len(grid[0])
	visited := make(map[Point]bool)
	queue := []Point{{X: 0, Y: 0}}
	count := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if !isValid(current, rows, cols) || visited[current] || !condition(grid[current.X][current.Y]) {
			continue
		}

		// Process the current element (e.g., increment count)
		count++

		// Mark the current element as visited
		visited[current] = true

		// Add neighbors to the queue
		neighbors := getNeighbors(current, rows, cols)
		for _, neighbor := range neighbors {
			queue = append(queue, neighbor)
		}
	}

	return count
}

func getNeighbors(point Point, rows, cols int) []Point {
	neighbors := []Point{
		{point.X - 1, point.Y},
		{point.X + 1, point.Y},
		{point.X, point.Y - 1},
		{point.X, point.Y + 1},
	}

	return neighbors
}

func isValid(point Point, rows, cols int) bool {
	return point.X >= 0 && point.X < rows && point.Y >= 0 && point.Y < cols
}

func main() {
	file, err := os.ReadFile("solution.txt")

	if err != nil {
		log.Fatal(err)
	}

	content := string(file)

	sections := strings.Split(content, "\r\n\r\n")
	actions, graph_values := sections[0], sections[1]

	//var starter []string
	//var starter []string
	TheDoraMapFromDora := make(map[string][]string)
	var allStartingNodes []string

	for _, line := range strings.Split(graph_values, "\n") {
		node := line[0:3]
		left := line[7:10]
		right := line[12:15]
		if line[2] == 'A' {
			allStartingNodes = append(allStartingNodes, node)
		}
		TheDoraMapFromDora[node] = []string{left, right}
		fmt.Println(node, left, right)
	}
	//printStringArrayMapSortedByKey(TheDoraMapFromDora)
	//spew.Dump(TheDoraMapFromDora)
	fmt.Println(actions)
	count := 0

	//fmt.Println(Finalcount)
	fmt.Println(allStartingNodes)
	var endCounts []int
	for _, starting_nodes := range allStartingNodes {
		Finalcount := 0
		Finalcount = goThroughMap(actions, TheDoraMapFromDora, count, starting_nodes)
		endCounts = append(endCounts, Finalcount)

	}
	LCMNumber := lcmArray(endCounts)

	fmt.Println(endCounts)
	fmt.Println(LCMNumber)

}
func goThroughMap(directions string, TheDoraMapFromDora map[string][]string, count int, key string) int {
	//[AAA QRA KQA DFA DBA HJA]
	if key[2] == 'Z' {
		return count
	}
	if directions[0] == 'L' {
		count++
		updatedirections := directions[1:] + "L"

		newKey := TheDoraMapFromDora[key][0]
		return goThroughMap(updatedirections, TheDoraMapFromDora, count, newKey)

	} else if directions[0] == 'R' {
		count++
		updatedirections := directions[1:] + "R"
		newKey := TheDoraMapFromDora[key][1]
		return goThroughMap(updatedirections, TheDoraMapFromDora, count, newKey)

	}
	return 0
}
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
func lcmArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]

	for i := 1; i < len(arr); i++ {
		result = LCM(result, arr[i])
	}

	return result
}

// func printStringArrayMapSortedByKey(m map[string][]string) {
// 	// Create a slice to store keys
// 	var keys []string
// 	for key := range m {
// 		keys = append(keys, key)
// 	}

// 	// Sort the slice of keys
// 	sort.Strings(keys)

// 	fmt.Println("Printing Map (Sorted by Key):")
// 	for _, key := range keys {
// 		fmt.Printf("%s: %v\n", key, m[key])
// 	}
// }
