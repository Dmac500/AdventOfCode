package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	problem1()
}
func problem1() {
	file, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	content := string(file)
	//fmt.Println(content)
	lines := strings.Split(content, "\r\n")
	// fmt.Println(lines)
	total := 0
	for index, line := range lines {
		fmt.Println(index, " : ", line)
		intArray, err := stringArrayToIntArray(stringToStringArray(line))
		if err != nil {
			panic(err)
		}
		//fmt.Println("this is the int Array : ", intArray)
		var Last_digit_array []int
		lastPattern := FindThePattern(intArray, Last_digit_array)
		//fmt.Println(lastPattern)
		total += lastPattern
	}
	fmt.Println(total)

}
func FindThePattern(array []int, lastDigitArray []int) int {

	var DiffenceArray []int
	// storeLastElement := array[len(array)-1]
	// lastDigitArray = append(lastDigitArray, storeLastElement)
	storeLastElement := array[0]
	lastDigitArray = append(lastDigitArray, storeLastElement)
	if allSame(array) {

		//return findTheMissingNumber(lastDigitArray)
		return findTheFirstMissingNumber(lastDigitArray)
	}
	for index := 0; index < len(array)-1; index++ {

		diffence := array[index+1] - array[index]

		DiffenceArray = append(DiffenceArray, diffence)
	}

	return FindThePattern(DiffenceArray, lastDigitArray)

}
func stringArrayToIntArray(strArray []string) ([]int, error) {
	intArray := make([]int, len(strArray))

	for i, str := range strArray {
		num, err := strconv.Atoi(str)
		if err != nil {
			// Handle the error (e.g., return an error or skip the problematic element)
			return nil, err
		}
		intArray[i] = num
	}

	return intArray, nil
}
func stringToStringArray(input string) []string {
	words := strings.Fields(input)

	// Convert each word to a string array
	strArray := make([]string, len(words))
	for i, word := range words {
		strArray[i] = word
	}

	return strArray
}
func allSame(arr []int) bool {
	// Check if the array is empty
	if len(arr) == 0 {
		return true
	}

	// Compare each element with the first one
	firstElement := arr[0]
	for _, element := range arr {
		if element != firstElement {
			return false
		}
	}

	return true
}
func findTheMissingNumber(arr []int) int {
	lastDidgit := 0
	for index := len(arr) - 1; index >= 0; index-- {
		lastDidgit = lastDidgit + arr[index]
		//fmt.Println(lastDidgit)
	}
	// fmt.Println("final solution: ", lastDidgit)
	return lastDidgit

}
func findTheFirstMissingNumber(arr []int) int {
	lastDidgit := 0
	fmt.Println(arr)

	for index := len(arr) - 1; index >= 0; index-- {
		lastDidgit = arr[index] - lastDidgit

		//fmt.Println(lastDidgit)
	}

	fmt.Println("final solution: ", lastDidgit)
	return lastDidgit

}
