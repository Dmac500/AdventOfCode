package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	problem2()
}
func problem2() {
	content, err := os.ReadFile("Stupid_Elf_game.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	//TotalPoints := 0
	// PlayerMap := make(map[string][]string)
	// WinMap := make(map[string][]string)
	var copyArray []int
	var copysToUse []int

	for gameIndex, i := range lines {

		// var storage []int
		firstSplit := strings.Split(i, ":")
		sides := strings.Split(firstSplit[len(firstSplit)-1], "|")
		playersCards := strings.Split(sides[0], " ")
		winningCards := strings.Split(sides[len(sides)-1], " ")
		var playersCard []string
		var winCard []string

		for _, i := range playersCards {
			if isNumeric(i) {
				playersCard = append(playersCard, i)
			}

		}
		for _, i := range winningCards {
			num := strings.ReplaceAll(i, "\r", "")
			if isNumeric(num) {
				winCard = append(winCard, strings.TrimSpace(num))
			}

		}
		// Game := "Game" + strconv.Itoa(gamenum+1)

		// 1 W4 - 2 W2 C2 W2 - 3 W2 C3 W2 C3 W2 C3 W2  - 4 C1 C2 C2 C3 C3 C3 C3 - 5 C1 C3 C3 C3 C3 C4 C4 C4 C4 C4 C4 C4 C4 bust and 6 is a bust
		// 1so  -   2 so       -  4 so                  -  8 so                  - 14 so
		// per matching numbers deliver to k spots in front of the win
		// if on check to see if u have copys and distribute those winning too

		// PlayerMap[Game] = playersCard
		// printMap(PlayerMap)
		// WinMap[Game] = winCard
		// fmt.Println(len(playersCard))
		// fmt.Println(playersCard)
		// fmt.Println(len(winCard))
		// fmt.Println(winCard)
		pointcounter := 0
		for _, i := range playersCard {

			for _, t := range winCard {
				// fmt.Println("this is our card ", i, " :: this is the winning cards ", t)
				if i == t {
					pointcounter = pointcounter + 1
				}
			}

		}

		copysToUse = append(copysToUse, pointcounter) // giving out

		score := 1

		fmt.Println("players cards: ", playersCard)
		fmt.Println("winning cards: ", winCard)

		for i := gameIndex - 1; i >= 0; i-- { // see prev scores and add them to yours.
			fmt.Println("Rotation: ", gameIndex)
			if copysToUse[i] < 0 {
				continue
			} else {
				//fmt.Println("walking into the club with ", i-1, "bitches")

				if copysToUse[i] > 0 {
					fmt.Println("copys to be used : ", copysToUse)
					score += copyArray[i]
					copysToUse[i] = copysToUse[i] - 1
					// fmt.Println("Score : ", copyArray)

				}
			}
		}
		copyArray = append(copyArray, score)
		fmt.Println("Score : ", copyArray)

		//fmt.Println(copyArray)

	}
	total := 0
	fmt.Println("lenth ", len(copyArray))
	for i := 0; i < len(copyArray); i++ {
		total = total + copyArray[i]
	}
	fmt.Println(total)
}

func problem1() {
	content, err := os.ReadFile("Stupid_Elf_game.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	TotalPoints := 0

	for _, i := range lines {
		firstSplit := strings.Split(i, ":")
		sides := strings.Split(firstSplit[len(firstSplit)-1], "|")
		playersCards := strings.Split(sides[0], " ")
		winningCards := strings.Split(sides[len(sides)-1], " ")
		var playersCard []string
		var winCard []string

		for _, i := range playersCards {
			if isNumeric(i) {
				playersCard = append(playersCard, i)
			}

		}
		for _, i := range winningCards {
			num := strings.ReplaceAll(i, "\r", "")
			if isNumeric(num) {
				winCard = append(winCard, strings.TrimSpace(num))
			}

		}
		// fmt.Println(len(playersCard))
		// fmt.Println(playersCard)
		// fmt.Println(len(winCard))
		// fmt.Println(winCard)
		pointcounter := 0
		for _, i := range playersCard {

			for _, t := range winCard {
				// fmt.Println("this is our card ", i, " :: this is the winning cards ", t)
				if i == t {
					if pointcounter == 0 {
						fmt.Println("THIS IS THE WINNER ", i, " current points: ", pointcounter)
						pointcounter = 1
					} else {
						fmt.Println("THIS IS THE WINNER ", i, " current points: ", pointcounter)
						pointcounter = pointcounter * 2
					}

				}
			}

		}
		TotalPoints = TotalPoints + pointcounter
		fmt.Println(TotalPoints)

	}

}

func convertToDoubleDigit(arr []string) []string {
	// Initialize a new slice to store the converted cards
	var result []string

	for _, str := range arr {
		// Check if the string has only one digit
		if len(str) == 1 {
			// Convert the string to an integer
			num, err := strconv.Atoi(str)
			if err == nil {
				// Convert the integer back to a string with leading zero
				result = append(result, fmt.Sprintf("%02d", num))
			}
		} else {
			// If the string has more than one digit, simply append it
			result = append(result, str)
		}
	}

	return result
}
func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
