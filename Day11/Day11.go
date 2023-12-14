package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// |x2 - x1|

func main() {
	problem1()
}

type Row struct {
	rownum     int
	hasGalaxey bool
}
type Col struct {
	colnum     int
	hasGalaxey bool
}
type pointLoc struct {
	gal int
	x   int
	y   int
}

func problem1() {
	// file, err := os.ReadFile("littleSpace.txt")
	file, err := os.Open("bigSpace.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	//content := string(file)
	//lines := strings.Split(content, "\n")
	scanner := bufio.NewScanner(file)
	var OuterSoace [][]string

	for scanner.Scan() {

		line := strings.Split(strings.TrimSpace(scanner.Text()), "")
		OuterSoace = append(OuterSoace, line)
	}

	LookForMissingGalaxy(OuterSoace)

}
func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Println("this is the board")
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j])
		}
		fmt.Println()
	}

}
func markLocations(board [][]string) []pointLoc {
	var Points []pointLoc
	count := 0
	for i := 0; i < len(board); i++ {
		//fmt.Println("this is the board")
		for j := 0; j < len(board[i]); j++ {
			//fmt.Print(board[i][j])
			if board[i][j] == "#" {
				count++
				//fmt.Println(count)
				var point pointLoc
				point.gal = count
				point.x = i
				point.y = j
				Points = append(Points, point)
				board[i][j] = strconv.Itoa(count)
			}
		}

	}
	return Points
}
func LookForMissingGalaxy(board [][]string) {
	var rowArray []Row
	var ColArray []Col

	for i := 0; i < len(board); i++ {

		for j := 0; j < len(board[i]); j++ {
			//fmt.Print(board[i][j])
			if !rowContains(rowArray, i) {
				var row Row
				row.rownum = i
				row.hasGalaxey = false
				rowArray = append(rowArray, row)
			}
			if !ColContains(ColArray, j) {
				var col Col
				col.colnum = j
				col.hasGalaxey = false
				ColArray = append(ColArray, col)
			}
			if board[i][j] == "#" {
				rowArray[i].hasGalaxey = true
				ColArray[j].hasGalaxey = true

			}

		}
		//fmt.Println()
	}
	fmt.Println(rowArray)
	fmt.Println(ColArray)
	var numsToExpandCol []int
	var numsToExpandRow []int

	numsToExpandRow, numsToExpandCol = checkWhoGetsTheExpantion(rowArray, ColArray)

	//

	//board = expandBoard(board, numsToExpandCol, numsToExpandRow, 0)
	points := markLocations(board)
	total := 0
	new_points := pointexpand(points, numsToExpandCol, numsToExpandRow, 999999)
	fmt.Println(new_points)
	final := doMathStuff(new_points, total, numsToExpandRow, numsToExpandCol, 0)
	fmt.Println(final)
}

// func doMathStuffPart2(board [][]string, numsToExpandRow []int, numsToExpandCol []int, points []pointLoc, total int) int {
// 	startingPoint := points[0]
// 	//fmt.Println("number: ", startingPoint.gal, "(", startingPoint.x, ",", startingPoint.y, ")")
// 	if len(points) == 1 {

// 		return total
// 	}

// 	newPointsArray := points[1:]
// 	for i := 0; i < len(newPointsArray); i++ {
// 		abX := abs(newPointsArray[i].x - startingPoint.x)
// 		abY := abs(newPointsArray[i].y - startingPoint.y)
// 		matDistance := abX + abY
// 		total = total + matDistance
// 	}
// 	return doMathStuff(newPointsArray, total)

// }
func pointexpand(points []pointLoc, numsToExpandCol []int, numsToExpandRow []int, adder int) []pointLoc {
	var NewPoints []pointLoc
	for index, point := range points {
		var newPoint pointLoc
		//fmt.Println(point)
		newYPoint := point.y
		newXPoint := point.x
		for _, rowNum := range numsToExpandRow {
			if rowNum < point.x {
				newXPoint += adder
			}
		}
		for _, colNum := range numsToExpandCol {
			//fmt.Println(colNum)
			if colNum < point.y {
				fmt.Println(point.x)

				newYPoint += adder
			}
		}
		newPoint.gal = index
		newPoint.x = newXPoint
		newPoint.y = newYPoint
		NewPoints = append(NewPoints, newPoint)
	}
	return NewPoints
}
func doMathStuff(points []pointLoc, total int, numsToExpandRow []int, numsToExpandCol []int, adder int) int {
	startingPoint := points[0]
	fmt.Println("number: ", startingPoint.gal, "(", startingPoint.x, ",", startingPoint.y, ")")
	if len(points) == 1 {
		return total
	}

	newPointsArray := points[1:]
	for i := 0; i < len(newPointsArray); i++ {
		x1 := startingPoint.x
		y1 := startingPoint.y
		x2 := newPointsArray[i].x
		y2 := newPointsArray[i].y
		// for j := 0; j < len(numsToExpandRow); j++ {
		// 	if numsToExpandRow[j] < startingPoint.y {
		// 		y1 += adder
		// 	}
		// 	if numsToExpandRow[j] < newPointsArray[i].y {
		// 		y2 += adder
		// 	}
		// }
		// for j := 0; j < len(numsToExpandCol); j++ {
		// 	if numsToExpandCol[j] < startingPoint.x {
		// 		fmt.Println("should get 10 ", numsToExpandCol[j], " < ", startingPoint.x)
		// 		x1 += adder
		// 	}
		// 	if numsToExpandCol[j] < newPointsArray[i].x {
		// 		fmt.Println("should get 10 ", numsToExpandCol[j], " < ", newPointsArray[i].x)
		// 		x2 += adder
		// 	}
		// }
		//fmt.Println("Pussy Fart number: ", newPointsArray[i].gal, "(", x2, ",", y2, ")")
		abX := abs(x2 - x1)
		abY := abs(y2 - y1)
		matDistance := abX + abY
		total = total + matDistance
	}
	return doMathStuff(newPointsArray, total, numsToExpandRow, numsToExpandCol, adder)

}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func expandBoard(grid [][]string, numsToExpandCol []int, numsToExpandRow []int, mutiplyer int) [][]string {

	var newRow []string
	for i := 0; i < len(grid); i++ {
		newRow = append(newRow, ".")
	}
	count := 0
	for trueindex, index := range numsToExpandRow {
		//fmt.Println(index)
		for i := 0; i < mutiplyer; i++ {
			grid = append(grid[:index], append([][]string{newRow}, grid[index:]...)...)

		}
		count += mutiplyer
		if trueindex+1 < len(numsToExpandRow) {
			numsToExpandRow[trueindex+1] += count
		}
	}

	counter := 0
	for trueindex, index := range numsToExpandCol {
		for i := 0; i < mutiplyer; i++ {
			for i := 0; i < len(grid); i++ {
				// Insert the new element at the specified column index

				row := grid[i]
				row = append(row, "")
				copy(row[index+1:], row[index:])
				row[index] = "."
				grid[i] = row

			}
		}
		counter += mutiplyer
		if trueindex+1 < len(numsToExpandCol) {
			numsToExpandCol[trueindex+1] += counter
		}

	}

	return grid
}
func checkWhoGetsTheExpantion(rowArray []Row, colArray []Col) ([]int, []int) {

	var returningValsforCol []int
	var returningValsforRow []int

	for index := range rowArray {
		if !rowArray[index].hasGalaxey {
			returningValsforRow = append(returningValsforRow, rowArray[index].rownum)
		}
	}

	for index := range colArray {
		if !colArray[index].hasGalaxey {
			returningValsforCol = append(returningValsforCol, colArray[index].colnum)
		}
	}

	return returningValsforRow, returningValsforCol

}

func rowContains(s []Row, e int) bool {
	for _, a := range s {
		if a.rownum == e {
			return true
		}
	}
	return false
}
func ColContains(s []Col, e int) bool {
	for _, a := range s {
		if a.colnum == e {
			return true
		}
	}
	return false
}

// func debug(board [][]string) {
// 	file, err := os.Create("debug.txt")
// 	if err != nil {
// 		fmt.Println("Error creating file:", err)
// 	}
// 	defer file.Close()
// 	for i := 0; i < len(board); i++ {
// 		//fmt.Println("this is the board")
// 		for j := 0; j < len(board[i]); j++ {
// 			_, err := fmt.Fprint(file, board[i][j])

// 			if err != nil {
// 				fmt.Println("Error writing to file:", err)
// 			}
// 		}
// 		fmt.Fprint(file, "\n")
// 	}
// }
