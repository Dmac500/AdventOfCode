package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// type Point struct {
// 	X, Y int
// }

// func bfs(grid [][]string, condition func(string) bool) int {
// 	rows, cols := len(grid), len(grid[0])
// 	visited := make(map[Point]bool)
// 	queue := []Point{{X: 0, Y: 0}}
// 	count := 0

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]

// 		if !isValid(current, rows, cols) || visited[current] || !condition(grid[current.X][current.Y]) {
// 			continue
// 		}

// 		// Process the current element (e.g., increment count)
// 		count++

// 		// Mark the current element as visited
// 		visited[current] = true

// 		// Add neighbors to the queue
// 		neighbors := getNeighbors(current, rows, cols)
// 		for _, neighbor := range neighbors {
// 			queue = append(queue, neighbor)
// 		}
// 	}

// 	return count
// }

// func getNeighbors(point Point, rows, cols int) []Point {
// 	neighbors := []Point{
// 		{point.X - 1, point.Y},
// 		{point.X + 1, point.Y},
// 		{point.X, point.Y - 1},
// 		{point.X, point.Y + 1},
// 	}

// 	return neighbors
// }

//	func isValid(point Point, rows, cols int) bool {
//		return point.X >= 0 && point.X < rows && point.Y >= 0 && point.Y < cols
//	}
func main() {
	problem1()
}
func problem1() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	var gameBoard [][]string

	for _, line := range lines {

		if line == "" {
			continue
		}
		row := strings.Split(line, "")
		gameBoard = append(gameBoard, row)
	}
	starting_row, starting_col := findStartingPoint(gameBoard)
	fmt.Println("this is the starting row: ", starting_row)
	fmt.Println("this is the starting starting_col: ", starting_col)
	fmt.Println(find_farthest_path(gameBoard))

}

// type Queue struct {
// 	List []string
// }

func find_farthest_path(board [][]string) int {
	// | is a vertical pipe connecting north and south.
	// - is a horizontal pipe connecting east and west.
	// L is a 90-degree bend connecting north and east.
	// J is a 90-degree bend connecting north and west.
	// 7 is a 90-degree bend connecting south and west.
	// F is a 90-degree bend connecting south and east.
	// . is ground; there is no pipe in this tile.
	// S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

	starting_row, starting_col := findStartingPoint(board)
	//count := 0
	//currNode := board[starting_row][starting_col]
	Loop_Not_completed_Flag := true
	// find starting dir
	current_row := starting_row
	current_col := starting_col
	var direction string
	count := 0
	board[current_row][current_col] = "1"

	Checknorth := starting_row - 1

	if Checknorth == -1 {
		Checknorth = 0
	}

	if board[Checknorth][starting_col] == "|" || board[Checknorth][starting_col] == "7" || board[Checknorth][starting_col] == "F" {
		//head north
		fmt.Println("your fucking dumb")
		current_row = starting_row - 1
		current_col = starting_col
		direction = "N"

		count++
	} else if board[starting_row][starting_col+1] == "-" || board[starting_row][starting_col+1] == "J" || board[starting_row][starting_col+1] == "7" {
		current_row = starting_row
		current_col = starting_col + 1
		direction = "E"
		fmt.Println("go east")
		count++
	} else if board[starting_row+1][starting_col] == "|" || board[starting_row+1][starting_col] == "L" || board[starting_row+1][starting_col] == "J" {
		//head south
		current_row = starting_row + 1
		current_col = starting_col
		direction = "S"
		fmt.Println("go south")

		count++
	} else if board[starting_row][starting_col-1] == "-" || board[starting_row][starting_col-1] == "L" || board[starting_row][starting_col-1] == "F" {
		current_row = starting_row
		current_col = starting_col - 1
		direction = "W"
		fmt.Println("go west")

		count++
	} else {
		return 0
	}

	for Loop_Not_completed_Flag {

		if count != 0 && current_row == starting_row && current_col == starting_col {
			file, err := os.Create("example.txt")
			if err != nil {
				fmt.Println("Error creating file:", err)
			}
			defer file.Close()
			for i := 0; i < len(board); i++ {
				//fmt.Println("this is the board")
				for j := 0; j < len(board[i]); j++ {
					_, err := fmt.Fprint(file, board[i][j])

					if err != nil {
						fmt.Println("Error writing to file:", err)
					}
				}
				fmt.Fprint(file, "\n")
			}

			return count / 2
		}
		if direction == "N" {

			if board[current_row][current_col] == "|" {
				current_row = current_row - 1
				board[current_row+1][current_col] = "^"
				count++
			} else if board[current_row][current_col] == "7" {
				direction = "W"
				current_col--
				board[current_row][current_col+1] = "<"
				count++
			} else if board[current_row][current_col] == "F" {
				direction = "E"
				current_col++
				board[current_row][current_col-1] = ">"
				count++
			}

		} else if direction == "E" {

			if board[current_row][current_col] == "-" {
				current_col = current_col + 1
				board[current_row][current_col-1] = ">"
				count++
			} else if board[current_row][current_col] == "J" {
				direction = "N"
				current_row--
				board[current_row+1][current_col] = "^"
				count++
			} else if board[current_row][current_col] == "7" {
				direction = "S"
				current_row++
				board[current_row-1][current_col] = "v"
				count++
			}
		} else if direction == "S" {

			if board[current_row][current_col] == "|" {
				current_row = current_row + 1
				board[current_row-1][current_col] = "v"
				count++
			} else if board[current_row][current_col] == "J" {
				direction = "W"
				current_col--
				board[current_row][current_col+1] = "<"
				count++
			} else if board[current_row][current_col] == "L" {
				direction = "E"
				current_col++
				board[current_row][current_col-1] = ">"
				count++
			}

		} else if direction == "W" {
			if board[current_row][current_col] == "-" {
				current_col = current_col - 1
				board[current_row][current_col+1] = "<"
				count++
			} else if board[current_row][current_col] == "L" {
				direction = "N"
				current_row--
				board[current_row+1][current_col] = "^"
				count++
			} else if board[current_row][current_col] == "F" {
				direction = "S"
				current_row++
				board[current_row-1][current_col] = "v"
				count++
			}
		}

	}

	return 0
}
func findStartingPoint(board [][]string) (int, int) {
	for i := 0; i < len(board); i++ {
		fmt.Println("this is the board")
		for j := 0; j < len(board[i]); j++ {
			if "S" == board[i][j] {
				return i, j
			}
		}
		fmt.Println()
	}
	return 0, 0

}
func markOpens() {
	file, err := os.ReadFile("example.txt")

	if err != nil {
		log.Fatal(err)
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	var gameBoard [][]string

	for _, line := range lines {

		if line == "" {
			continue
		}
		row := strings.Split(line, "")
		gameBoard = append(gameBoard, row)
	}
	printBoard(gameBoard)

	for i := 0; i < len(gameBoard); i++ {
		for j := 0; j < len(gameBoard[i]); j++ {
			if gameBoard[i][j] != "1" && gameBoard[i][j] != "O" {
				north := i - 1
				south := i + 1
				east := j + 1
				west := j - 1
				if north == -1 && gameBoard[i][j] != "1" {
					gameBoard[i][j] = "O"
				} else if south == len(gameBoard) && gameBoard[i][j] != "1" {
					gameBoard[i][j] = "O"
				} else if east == len(gameBoard[i])-1 && gameBoard[i][j] != "1" {
					gameBoard[i][j] = "O"
				} else if west == -1 && gameBoard[i][j] != "1" {
					gameBoard[i][j] = "O"
				} else if gameBoard[north][j] == "O" {
					gameBoard[i][j] = "O"
				} else if gameBoard[south][j] == "O" {
					gameBoard[i][j] = "O"
				} else if gameBoard[i][east] == "O" {
					gameBoard[i][j] = "O"
				} else if gameBoard[i][west] == "O" {
					gameBoard[i][j] = "O"
				}

			}

		}
	}
	printBoard(gameBoard)

}
func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		//fmt.Println("this is the board")
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}

}
func problem2() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	content := string(file)
	lines := strings.Split(content, "\n")

	var gameBoard [][]string

	for _, line := range lines {

		if line == "" {
			continue
		}
		row := strings.Split(line, "")
		gameBoard = append(gameBoard, row)
	}
}
