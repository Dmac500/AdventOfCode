// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"sync"
// )

// func main() {
// 	fmt.Println("hello world")
// 	problem1()
// }

// var maps map[string]map[string][]int

// const MaxUint = ^uint(0)
// const MaxInt = int(MaxUint >> 1)

// func isNumeric(str string) bool {
// 	_, err := strconv.Atoi(str)
// 	return err == nil
// }
// func problem1() {
// 	var seeds []int
// 	seedToSoil := make(map[string][]int)
// 	soilToFertilizer := make(map[string][]int)
// 	fertilizerToWater := make(map[string][]int)
// 	waterToLight := make(map[string][]int)
// 	lightToTemp := make(map[string][]int)
// 	temperatureToHumidity := make(map[string][]int)
// 	humidityToLocation := make(map[string][]int)
// 	content, err := os.Open("solution.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	scanner := bufio.NewScanner(content)

// 	keyvals := 0
// 	flagCounter := 0
// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		words := strings.Fields(line)
// 		// for _, t := range words {
// 		// 	fmt.Println(t)
// 		// }
// 		//fmt.Println(line)
// 		if line == "" {
// 			flagCounter = flagCounter + 1
// 			keyvals = 0
// 			continue
// 		}
// 		if flagCounter == 0 {

// 			for _, t := range words[1:] {
// 				num, err := strconv.Atoi(t)
// 				if err != nil {
// 					panic(err)
// 				}
// 				seeds = append(seeds, num)
// 			}
// 		} else if flagCounter == 1 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				seedToSoil[key] = arr
// 			}

// 		} else if flagCounter == 2 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				soilToFertilizer[key] = arr
// 			}

// 		} else if flagCounter == 3 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				fertilizerToWater[key] = arr

// 			}

// 		} else if flagCounter == 4 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				waterToLight[key] = arr

// 			}

// 		} else if flagCounter == 5 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				lightToTemp[key] = arr

// 			}

// 		} else if flagCounter == 6 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				temperatureToHumidity[key] = arr

// 			}

// 		} else if flagCounter == 7 {
// 			if !isNumeric(words[0]) {
// 				continue
// 			} else {
// 				var arr []int
// 				for _, sadness := range words {
// 					num, err := strconv.Atoi(sadness)
// 					if err != nil {
// 						panic(err)
// 					}
// 					arr = append(arr, num)

// 				}

// 				keyvals = keyvals + 1
// 				key := "key" + strconv.Itoa(keyvals)
// 				humidityToLocation[key] = arr
// 			}

// 		}

// 		// if words[0] == "seed-to-soil" {

// 		// }

// 		// if isword {

// 		// 	var arr []int
// 		// 	for _, sadness := range words {
// 		// 		num, err := strconv.Atoi(sadness)
// 		// 		if err != nil {
// 		// 			panic(err)
// 		// 		}
// 		// 		arr = append(seeds, num)

// 		// 	}

// 		// 	keyvals = keyvals + 1
// 		// 	key := "key" + string(keyvals)
// 		// 	seedToSoil[key] = arr

// 		// }

// 		// words := strings.Fields(line)

// 	}

// 	//var lowestNumArray []int
// 	//var newSeeds []int

// 	//arrayofMaps := make([]map[int]bool, len(seeds)/2)
// 	//counter := 0
// 	//outof := len(seeds) / 2
// 	//basket := make(map[int][]int)
// 	var lowestNum = MaxInt
// 	fmt.Println("LET SHIT FLY !!!!!")
// 	for i := 0; i < len(seeds); i = i + 2 {

// 		fuckhead := seeds[i] + seeds[i+1]
// 		//	rangemap := make(map[int]bool)

// 		for t := seeds[i]; t < fuckhead; t++ {
// 			//newSeeds = append(newSeeds, t)
// 			dickhead := pussylips(t, seedToSoil, soilToFertilizer, fertilizerToWater, waterToLight, lightToTemp, temperatureToHumidity, humidityToLocation)
// 			if dickhead < lowestNum {
// 				lowestNum = dickhead
// 			}
// 		}

// 		//arrayofMaps[counter] = rangemap
// 		// basket[counter] = newSeeds
// 		// counter++
// 		// fmt.Println("progress: ", counter, " / ", outof)
// 		fmt.Println("this is the new lowest num: ", lowestNum)
// 	}
// 	fmt.Println("Finished bud ")
// 	// 	var wg sync.WaitGroup
// 	// 	var mutex sync.Mutex
// 	// 	var resultChan = make(chan int)

// 	// 	// Number of workers
// 	// 	numWorkers := len(seeds) / 2

// 	// 	// Set the number of operating system threads
// 	// 	runtime.GOMAXPROCS(numWorkers)

// 	// 	// Start multiple goroutines (workers)
// 	// 	giveMeZero := 0
// 	// 	for i := 1; i <= numWorkers; i++ {
// 	// 		wg.Add(2)
// 	// 		go worker(i, &wg, &mutex, seedToSoil, soilToFertilizer, fertilizerToWater,
// 	// 			waterToLight, lightToTemp, temperatureToHumidity, humidityToLocation, basket[giveMeZero], resultChan)
// 	// 		giveMeZero++
// 	// 	}

// 	// 	// Wait for all goroutines to finish
// 	// 	go func() {
// 	// 		wg.Wait()
// 	// 		close(resultChan)
// 	// 	}()
// 	// 	for result := range resultChan {
// 	// 		fmt.Printf("Received result: %d\n", result)
// 	// 	}

// 	// 	fmt.Println("All workers have finished.")

// 	// }
// }

// func containsValue(arr []int, val int) bool {
// 	for _, element := range arr {
// 		if element == val {
// 			return true
// 		}
// 	}
// 	return false
// }

// func printMap(inputMap map[string][]int) {
// 	for key, value := range inputMap {
// 		fmt.Printf("Key: %s -> Value: %v\n", key, value)
// 	}
// }
// func findLowestNum(sad []int) {
// 	smallest := sad[0]
// 	for _, value := range sad {
// 		if value < smallest {
// 			smallest = value
// 		}
// 	}
// 	fmt.Println("jesus found where the small seed goes ", smallest)
// }

// func worker(id int, wg *sync.WaitGroup, mutex *sync.Mutex, seedToSoil map[string][]int, soilToFertilizer map[string][]int, fertilizerToWater map[string][]int,
// 	waterToLight map[string][]int, lightToTemp map[string][]int, temperatureToHumidity map[string][]int, humidityToLocation map[string][]int, rangemap []int, resultChan chan<- int) {
// 	defer wg.Done()

// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	fmt.Printf("Worker %d started\n", id)
// 	var lowestNum = MaxInt
// 	for f := range rangemap {

// 		e := f
// 		//fmt.Println(e)

// 		for _, value := range seedToSoil {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			// if in range add diff
// 			new_num := value[0]
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				//fmt.Println("seedToSoil: ", e)
// 				break
// 			}

// 		}
// 		for _, value := range soilToFertilizer {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			new_num := value[0]
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				break
// 			}

// 		}
// 		for _, value := range fertilizerToWater {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			// if in range add diff
// 			new_num := value[0]
// 			// fmt.Println("this is high: ", highNum)
// 			// fmt.Println("this is low: ", lowernum)
// 			// fmt.Println("this is new_num: ", new_num)
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				//fmt.Println("fertilizerToWater: ", e)
// 				break
// 			}

// 		}
// 		for _, value := range waterToLight {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			// if in range add diff
// 			new_num := value[0]
// 			// fmt.Println("this is high: ", highNum)
// 			// fmt.Println("this is low: ", lowernum)
// 			// fmt.Println("this is new_num: ", new_num)
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				//fmt.Println("waterToLight: ", e)
// 				break
// 			}

// 		}
// 		for _, value := range lightToTemp {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			// if in range add diff
// 			new_num := value[0]
// 			// fmt.Println("this is high: ", highNum)
// 			// fmt.Println("this is low: ", lowernum)
// 			// fmt.Println("this is new_num: ", new_num)
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				//	fmt.Println("lightToTemp: ", e)
// 				break
// 			}

// 		}
// 		for _, value := range temperatureToHumidity {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			// if in range add diff
// 			new_num := value[0]
// 			// fmt.Println("this is high: ", highNum)
// 			// fmt.Println("this is low: ", lowernum)
// 			// fmt.Println("this is new_num: ", new_num)
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				//fmt.Println("temperatureToHumidity: ", e)
// 				break
// 			}

// 		}
// 		for _, value := range humidityToLocation {

// 			lowernum := value[1]
// 			highNum := value[1] + value[2]
// 			// if in range add diff
// 			new_num := value[0]
// 			// fmt.Println("this is high: ", highNum)
// 			// fmt.Println("this is low: ", lowernum)
// 			// fmt.Println("this is new_num: ", new_num)
// 			if (lowernum <= e) && (e < highNum) {
// 				t := new_num + (e - lowernum)
// 				e = t
// 				//fmt.Println("humidityToLocation: ", e)
// 				break
// 			}

// 		}

// 		if e < lowestNum {
// 			lowestNum = e
// 		}

// 	}
// 	fmt.Printf("Worker %d done\n", id)
// 	// fmt.Printf("this is the lowest num :", lowestNum)
// 	resultChan <- lowestNum
// }
// func pussylips(seedmath int, seedToSoil map[string][]int, soilToFertilizer map[string][]int, fertilizerToWater map[string][]int,
// 	waterToLight map[string][]int, lightToTemp map[string][]int, temperatureToHumidity map[string][]int, humidityToLocation map[string][]int) int {

// 	e := seedmath

// 	for _, value := range seedToSoil {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		// if in range add diff
// 		new_num := value[0]
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			//fmt.Println("seedToSoil: ", e)
// 			break
// 		}

// 	}
// 	for _, value := range soilToFertilizer {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		new_num := value[0]
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			break
// 		}

// 	}
// 	for _, value := range fertilizerToWater {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		// if in range add diff
// 		new_num := value[0]
// 		// fmt.Println("this is high: ", highNum)
// 		// fmt.Println("this is low: ", lowernum)
// 		// fmt.Println("this is new_num: ", new_num)
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			//fmt.Println("fertilizerToWater: ", e)
// 			break
// 		}

// 	}
// 	for _, value := range waterToLight {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		// if in range add diff
// 		new_num := value[0]
// 		// fmt.Println("this is high: ", highNum)
// 		// fmt.Println("this is low: ", lowernum)
// 		// fmt.Println("this is new_num: ", new_num)
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			//fmt.Println("waterToLight: ", e)
// 			break
// 		}

// 	}
// 	for _, value := range lightToTemp {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		// if in range add diff
// 		new_num := value[0]
// 		// fmt.Println("this is high: ", highNum)
// 		// fmt.Println("this is low: ", lowernum)
// 		// fmt.Println("this is new_num: ", new_num)
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			//	fmt.Println("lightToTemp: ", e)
// 			break
// 		}

// 	}
// 	for _, value := range temperatureToHumidity {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		// if in range add diff
// 		new_num := value[0]
// 		// fmt.Println("this is high: ", highNum)
// 		// fmt.Println("this is low: ", lowernum)
// 		// fmt.Println("this is new_num: ", new_num)
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			//fmt.Println("temperatureToHumidity: ", e)
// 			break
// 		}

// 	}
// 	for _, value := range humidityToLocation {

// 		lowernum := value[1]
// 		highNum := value[1] + value[2]
// 		// if in range add diff
// 		new_num := value[0]
// 		// fmt.Println("this is high: ", highNum)
// 		// fmt.Println("this is low: ", lowernum)
// 		// fmt.Println("this is new_num: ", new_num)
// 		if (lowernum <= e) && (e < highNum) {
// 			t := new_num + (e - lowernum)
// 			e = t
// 			//fmt.Println("humidityToLocation: ", e)
// 			break
// 		}

// 	}

//		return e
//	}
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func mapped_value_fromsection(mapping [][]int, key int) int {
	for _, row := range mapping {
		if key >= row[0] && key <= row[1] {
			distance := key - row[0]
			return row[2] + distance
		}
	}

	return key
}

func main() {

	data, err := os.ReadFile("solution.txt")

	if err != nil {
		log.Fatal(err)
	}

	content := string(data)

	sections := strings.Split(content, "\r\n\r\n")

	seeds := strings.Split(strings.TrimSpace(strings.Split(sections[0], ":")[1]), " ")

	// a-to-b -> sourceval -> destination
	mappings := make(map[string][][]int)

	for _, section := range sections[1:] {

		sections_lines := strings.Split(section, "\r\n")

		section_name, section_rows := strings.Split(sections_lines[0], " ")[0], sections_lines[1:]

		fmt.Println("Starting section: ", section_name)

		rows := [][]int{}

		for _, row := range section_rows {
			// fmt.Print("Line: ", idx)

			parsed_row := strings.Split(row, " ")
			destination_start, _ := strconv.Atoi(parsed_row[0])
			source_start, _ := strconv.Atoi(parsed_row[1])
			length, _ := strconv.Atoi(parsed_row[2])

			temp := []int{source_start, source_start + length, destination_start, length}

			rows = append(rows, temp)

			// fmt.Println(" Done")

		}

		mappings[section_name] = rows
	}

	// spew.Dump(mappings)

	min_loc := math.MaxInt

	start := time.Now()

	fmt.Println(seeds)
	for i := 0; i < len(seeds); i += 2 {
		start_seed, _ := strconv.Atoi(seeds[i])
		distance, _ := strconv.Atoi(seeds[i+1])

		t := time.Now()
		elapsed := t.Sub(start)

		fmt.Print("Checking: ", start_seed, ":", start_seed+distance, " Seconds: ", elapsed.Seconds())

		for seed_val_temp := start_seed; seed_val_temp < start_seed+distance; seed_val_temp++ {
			// seedval,  := strconv.Atoi(seed)

			seed_val := mapped_value_from_section(mappings["seed-to-soil"], seed_val_temp)

			// fmt.Println("\tsoil: ", seed_val)

			seed_val = mapped_value_from_section(mappings["soil-to-fertilizer"], seed_val)

			// fmt.Println("\tfertilizer: ", seed_val)

			seed_val = mapped_value_from_section(mappings["fertilizer-to-water"], seed_val)
			// fmt.Println("\twater: ", seed_val)

			seed_val = mapped_value_from_section(mappings["water-to-light"], seed_val)
			// fmt.Println("\tlight: ", seed_val)

			seed_val = mapped_value_from_section(mappings["light-to-temperature"], seed_val)
			// fmt.Println("\ttemperature: ", seed_val)

			seed_val = mapped_value_from_section(mappings["temperature-to-humidity"], seed_val)
			// fmt.Println("\thumidity: ", seed_val)

			seed_val = mapped_value_from_section(mappings["humidity-to-location"], seed_val)
			// fmt.Println("Seed: ", seed_val, " \tlocation: ", seed_val)

			if seed_val < min_loc {
				min_loc = seed_val
			}
		}
		fmt.Println(" Smallest: ", min_loc)

	}

	// for , seed := range seeds {

	// }

	fmt.Println("Smallest: ", min_loc)

	// spew.Dump(mappings)

}
func mapped_value_from_section(mapping [][]int, key int) int {
	for _, row := range mapping {
		if key >= row[0] && key <= row[1] {
			distance := key - row[0]
			return row[2] + distance
		}
	}

	return key
}
