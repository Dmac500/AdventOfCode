package main

import "fmt"

func main() {
	problem1()
}
func problem1() {
	var times = [1]int{40709879}
	var racedistance = [1]int{215105121471005}
	wins := 1

	for races, i := range times {
		raceWins := 0
		for buttonHold := 0; buttonHold < i; buttonHold++ {
			travelTime := i - buttonHold
			distance := buttonHold * travelTime
			if distance > racedistance[races] {
				raceWins++
			}

		}
		wins *= raceWins
	}
	fmt.Println(wins)
}
