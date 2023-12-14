package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Hand is the string of cards
// kind of Hand is the Value we will give to the Hand
// Value will be the bid

type Bid struct {
	Hand       string
	KindofHand int
	Value      int
}

func main() {
	problem1()
}
func problem1() {
	data, err := os.ReadFile("solution.txt")

	if err != nil {
		log.Fatal(err)
	}

	HandsAndBids := make(map[string]int)

	content := string(data)
	//fmt.Println(content)

	lines := strings.Split(content, "\n")
	for _, line := range lines {

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("FUCK!!")
		}
		key := parts[0]
		valueStr := parts[1]
		value, _ := strconv.Atoi(valueStr)
		HandsAndBids[key] = value
	}

	printMap(HandsAndBids)

}
func printMap(m map[string]int) {

	fmt.Println("working on Hands:")
	var FuckMyCock []Bid
	//pq := make(PriorityQueue, 0)
	Hand := 0
	//PQP := 0
	for key, val := range m {
		Hand++
		KindofHand := ""
		valueOfHand := 0

		paircount := 0
		threeCount := 0
		fourCount := 0
		fiveCount := 0

		JokerAdder := 0

		counts := make(map[rune]int)

		for index, i := range key {
			if key[index] == 'J' {
				JokerAdder += 1
				//fmt.Println(key, " : ", JokerAdder)
			}
			counts[i]++

		}
		newCounts := sortByValue(counts)
		for _, kv := range newCounts {
			if JokerAdder == 5 {
				fmt.Println("KEVIN IS SO COOL")
				KindofHand = assignedHand(string(kv.Key), 5)
				fmt.Println("KEVIN IS SO COOL ", KindofHand)
				fiveCount++
			}
			if kv.Key == 'J' {
				continue
			}
			if kv.Value >= 1 {
				if JokerAdder < 5 {
					KindofHand = assignedHand(string(kv.Key), kv.Value+JokerAdder)
					JokerAdder = 0
				}
			} else {
				KindofHand = assignedHand(string(kv.Key), kv.Value+JokerAdder)
			}

			if KindofHand == "pair" {
				paircount++
			}
			if KindofHand == "threeOfAKind" {
				threeCount++
			}
			if KindofHand == "FourOfAKind" {
				fourCount++
			}
			if KindofHand == "FiveOfAKind" {
				fiveCount++
			}
			//fmt.Println("Hand", Hand, " : kind of card :  ", kindOfHand)
		}
		if fiveCount == 1 {
			valueOfHand += 7
			//fmt.Println("5ofakind")
		} else if fourCount == 1 {
			valueOfHand += 6
			//fmt.Println("4ofakind")
		} else if threeCount == 1 && paircount == 1 {
			valueOfHand += 5
			fmt.Println("fullhouse")
		} else if threeCount == 1 {
			valueOfHand += 4
			//fmt.Println("3ofakind")
		} else if paircount == 2 {
			valueOfHand += 3
			//fmt.Println("2 pair")
		} else if paircount == 1 {
			valueOfHand += 2
			//fmt.Println("pair")
		} else {
			valueOfHand += 1
			//fmt.Println("ALLSHIT")
		}
		//heap.Push(&pq, bid{Hand: key, KindofHand: valueOfHand, Value: val, Priority: PQP})
		//PQP++
		LookingAtHand := Bid{Hand: key, KindofHand: valueOfHand, Value: val}
		FuckMyCock = append(FuckMyCock, LookingAtHand)
	}
	//fmt.Println(FuckMyCock)
	sort.SliceStable(FuckMyCock, func(i, j int) bool {
		// fmt.Println("Comparing: ")
		// spew.Dump(hands[i])
		// spew.Dump(hands[j])

		if FuckMyCock[i].KindofHand != FuckMyCock[j].KindofHand {

			return FuckMyCock[i].KindofHand < FuckMyCock[j].KindofHand
		} else {

			order := "J23456789TQKA"

			for k := 0; k <= 5; k++ {
				if strings.Index(order, string(FuckMyCock[i].Hand[k])) != strings.Index(order, string(FuckMyCock[j].Hand[k])) {
					return strings.Index(order, string(FuckMyCock[i].Hand[k])) < strings.Index(order, string(FuckMyCock[j].Hand[k]))
				}
			}
		}

		return true
	})
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, line := range FuckMyCock {
		_, err := fmt.Fprintln(file, line.Hand)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	total_sum := 0
	for muti := 0; muti < len(FuckMyCock); muti++ {
		total_sum += FuckMyCock[muti].Value * (muti + 1)
	}
	if total_sum == 248750248 {
		fmt.Println("TRUE!!!!!!!!")
	} else {
		fmt.Println("Your fucking retarted")
	}

}
func assignedHand(card string, howMany int) string {
	KindofHand := ""
	//fmt.Println(card, " : ", howMany)
	switch howMany {
	case 1:
		KindofHand += "highcard"
	case 2:
		KindofHand += "pair"
	case 3:
		KindofHand += "threeOfAKind"
	case 4:
		KindofHand += "FourOfAKind"
	case 5:
		KindofHand += "FiveOfAKind"
	}

	return KindofHand

}

type KeyValue struct {
	Key   rune
	Value int
}

func sortByValue(inputMap map[rune]int) []KeyValue {
	// Convert the map to a slice of key-value pairs
	var keyValueSlice []KeyValue
	for key, value := range inputMap {
		keyValueSlice = append(keyValueSlice, KeyValue{Key: key, Value: value})
	}

	// Sort the slice based on values
	sort.Slice(keyValueSlice, func(i, j int) bool {
		return keyValueSlice[i].Value > keyValueSlice[j].Value
	})

	return keyValueSlice
}

// type PriorityQueue []bid

// func (pq PriorityQueue) Len() int { return len(pq) }

// func (pq PriorityQueue) Less(i, j int) bool {
// 	return pq[i].Priority < pq[j].Priority
// }

// func (pq PriorityQueue) Swap(i int, j int) {
// 	fmt.Println("calling swap")

// 	spew.Dump(pq[i])
// 	fmt.Println("this is the second stuct: ", pq[j])
// 	pq[i].Priority, pq[j].Priority = pq[j].Priority, pq[i].Priority
// 	pq[i].index = j
// 	pq[j].index = i

// 	// pq[i], pq[j] = pq[j], pq[i]
// 	fmt.Println("operation called")
// 	fmt.Println("this is the first stuct: ", pq[i])
// 	fmt.Println("this is the second stuct: ", pq[j])
// }

// func (pq *PriorityQueue) Push(x interface{}) {
// 	item := x.(bid)
// 	n := len(*pq)
// 	*pq = append(*pq, item)
// 	item.index = n
// }

// func (pq *PriorityQueue) Pop() interface{} {
// 	old := *pq
// 	n := len(old)
// 	item := old[n-1]
// 	item.index = -1
// 	*pq = old[0 : n-1]
// 	return item
// }
// func PrintQueue(pq PriorityQueue) {
// 	fmt.Printf("Priority Queue: ")
// 	for _, item := range pq {
// 		fmt.Printf("(%s:%d:%d:%d) ", item.Hand, item.KindofHand, item.Value, item.Priority)
// 	}
// 	fmt.Println()
// }
// func adjustPrioritys(pq PriorityQueue) {
// 	fmt.Printf("Priority Queue: ")
// 	if pq.Len() > 1 {
// 		for _, item := range pq {
// 			for _, next := range pq {
// 				if next.Hand == item.Hand {
// 					continue
// 				} else if next.KindofHand == item.KindofHand {
// 					// check whats higher in the can hind
// 					for i := 0; i < len(item.Hand); i++ {
// 						if evalCard(string(next.Hand[i])) < evalCard(string(next.Hand[i])) {
// 							pq.Swap(next.Priority, item.Priority)
// 						}

// 					}
// 				} else if next.KindofHand < item.KindofHand {
// 					fmt.Println("Pre Swap")
// 					fmt.Printf("item: (%s:%d:%d:%d) ", item.Hand, item.KindofHand, item.Value, item.Priority)
// 					fmt.Printf("next :(%s:%d:%d:%d) ", next.Hand, next.KindofHand, next.Value, next.Priority)
// 					fmt.Println()
// 					pq.Swap(next.Priority, item.Priority)
// 					fmt.Println("Post Swap")
// 					fmt.Printf("item: (%s:%d:%d:%d) ", item.Hand, item.KindofHand, item.Value, item.Priority)
// 					fmt.Printf("next :(%s:%d:%d:%d) ", next.Hand, next.KindofHand, next.Value, next.Priority)
// 					fmt.Println()
// 				} else {
// 					fmt.Printf("")
// 				}

// 			}
// 		}
// 	}

// }
func evalCard(card string) int {

	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 12
	case "T":
		return 10
	case "9":
		return 9
	case "8":
		return 8
	case "7":
		return 7
	case "6":
		return 6
	case "5":
		return 5
	case "4":
		return 4
	case "3":
		return 3
	case "2":
		return 2
	}
	return 0
}
