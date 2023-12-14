package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var f *os.File

func main() {
	//problem1()
	problem2()
}

var starMap map[string][]int

func problem2() {
	f, err := os.Create("logger3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	array, err := readLines("day3input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	starMap = make(map[string][]int)
	for i, val := range array {
		fmt.Println(i, " : ", val)
		prevWasNum := false
		var sb strings.Builder
		var startingIndex int

		for index, charr := range val {
			if unicode.IsDigit(charr) {
				if !prevWasNum {
					startingIndex = index
				}
				sb.WriteString(string(charr))
				prevWasNum = true

				if index == len(val)-1 {
					prevWasNum = false
					s, err := strconv.Atoi(sb.String())
					if err != nil {
						panic(err)
					}
					if checkForSymbol(i, array, index, startingIndex, s) {
						sb.Reset()
						total = s + total
						fmt.Println("this is your total : ", total)
					}

				}

			} else if (prevWasNum) && !(unicode.IsDigit(charr)) {
				prevWasNum = false
				s, err := strconv.Atoi(sb.String())
				f.WriteString(sb.String())
				if checkForSymbol(i, array, index, startingIndex, s) {
					f.WriteString(" is valid ")
					f.WriteString("\n")
					sb.Reset()
					if err != nil {
						panic(err)
					}
					fmt.Println(s)

					total = s + total

					fmt.Println("this is your total : ", total)

				} else {
					f.WriteString(" not valid")
					f.WriteString("\n")
					sb.Reset()
				}
			}
		}

	}
	fmt.Println(starMap)
	totalbuns := 0
	for _, value := range starMap {
		apple := 0
		if len(value) >= 2 {
			apple = 1
			for _, titty := range value {
				fmt.Println(titty)
				apple = titty * apple
			}
		}
		totalbuns = apple + totalbuns
	}
	fmt.Println(totalbuns)

}

func problem1() {
	f, err := os.Create("logger3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	array, err := readLines("day3input.txt")
	if err != nil {
		panic(err)
	}
	total := 0
	for i, val := range array {
		fmt.Println(i, " : ", val)
		prevWasNum := false
		var sb strings.Builder
		var startingIndex int
		for index, charr := range val {
			if unicode.IsDigit(charr) {
				if !prevWasNum {
					startingIndex = index
				}
				sb.WriteString(string(charr))
				prevWasNum = true

				if index == len(val)-1 {
					prevWasNum = false
					s, err := strconv.Atoi(sb.String())
					if err != nil {
						panic(err)
					}
					if checkForSymbol(i, array, index, startingIndex, s) {
						sb.Reset()
						total = s + total
						fmt.Println("this is your total : ", total)
					}

				}

			} else if (prevWasNum) && !(unicode.IsDigit(charr)) {
				prevWasNum = false
				s, err := strconv.Atoi(sb.String())
				f.WriteString(sb.String())
				if checkForSymbol(i, array, index, startingIndex, s) {
					f.WriteString(" is valid ")
					f.WriteString("\n")
					sb.Reset()
					if err != nil {
						panic(err)
					}
					fmt.Println(s)

					total = s + total

					fmt.Println("this is your total : ", total)

				} else {
					f.WriteString(" not valid")
					f.WriteString("\n")
					sb.Reset()
				}
			}
		}

	}
}
func checkForSymbol(linenumber int, numlist []string, endingIndex int, startingIndex int, s int) bool {
	if startingIndex != 0 {
		startingIndex = startingIndex - 1
	}

	if linenumber == 0 {
		currentLine := numlist[linenumber]
		nextLine := numlist[linenumber+1]
		for i := startingIndex; i <= endingIndex; i++ {
			if (currentLine[i] != '.') && !unicode.IsDigit(rune(currentLine[i])) || (nextLine[i] != '.') && !unicode.IsDigit(rune(nextLine[i])) {
				if currentLine[i] == '*' {
					key := strconv.Itoa(linenumber) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				} else if nextLine[i] == '*' {
					key := strconv.Itoa(linenumber+1) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				}
				return true
			}
		}
		fmt.Println("edge case one ")
	} else if linenumber == len(numlist)-1 {
		currentLine := numlist[linenumber]
		prevLine := numlist[linenumber-1]
		for i := startingIndex; i <= endingIndex; i++ {
			if (currentLine[i] != '.') && !unicode.IsDigit(rune(currentLine[i])) || (prevLine[i] != '.') && !unicode.IsDigit(rune(prevLine[i])) {
				if currentLine[i] == '*' {
					key := strconv.Itoa(linenumber) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				} else if prevLine[i] == '*' {
					key := strconv.Itoa(linenumber-1) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				}
				return true
			}
		}
	} else {
		nextLine := numlist[linenumber+1]
		currentLine := numlist[linenumber]
		prevLine := numlist[linenumber-1]
		// fmt.Println("===========start==================")
		// fmt.Println("this is prevline", string(prevLine))
		// fmt.Println("this is currline", string(currentLine))
		// fmt.Println("this is nextline", string(nextLine))
		// fmt.Println("===========end==================")
		for i := startingIndex; i <= endingIndex; i++ {
			fmt.Println("this is cancer ", string(currentLine[i]), " | starting index: ", startingIndex, "| ending Index", endingIndex)
			// fmt.Println("this is prevline", string(prevLine[i]))
			// fmt.Println("this is currline", string(currentLine[i]))
			// fmt.Println("this is nextline", string(nextLine[i]))
			if (currentLine[i] != '.') && !unicode.IsDigit(rune(currentLine[i])) || (prevLine[i] != '.') && !unicode.IsDigit(rune(prevLine[i])) || (nextLine[i] != '.') && !unicode.IsDigit(rune(nextLine[i])) {
				// fmt.Println(endingIndex)
				if currentLine[i] == '*' {
					key := strconv.Itoa(linenumber) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				} else if prevLine[i] == '*' {
					key := strconv.Itoa(linenumber-1) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				} else if nextLine[i] == '*' {
					key := strconv.Itoa(linenumber+1) + ":" + strconv.Itoa(i)
					starMap[key] = append(starMap[key], s)
				}
				return true
			}
		}

	}
	return false

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
