package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
)

//start inclusive, end exclusive
type Interval struct {
	start int64
	end int64
}
func main() {
	file, err := os.Open("sortedinput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var guards = make(map[int64][]Interval, 0)
	
	scanner := bufio.NewScanner(file)
	var currentGuard int64 = 0;
	var currentStart int64 = 0;
	var currentEnd int64 = 0;
	for scanner.Scan() {
		var line = scanner.Text()
		fmt.Println(line)
		var split = strings.Split(line, " ")

		var timeSplit = strings.Split(strings.Trim(split[1], "]"), ":")
		var hour, _ = strconv.ParseInt(timeSplit[0], 10, 32)
		if (hour == 0) {
			hour = 24
		}
		var min, _ = strconv.ParseInt(timeSplit[1], 10, 32)
		var time = min + hour*60 
		fmt.Println(hour, " " , min)
		fmt.Println(split[2]);
		if (split[2] == "Guard") {
			currentGuard, _ = strconv.ParseInt(strings.Trim(split[3], "#"), 10, 32)
			// currentStart = time
		} else if (split[2] == "wakes") {
			currentEnd = time
			var interval = Interval{start:currentStart, end:currentEnd}
			guards[currentGuard] = append(guards[currentGuard], interval)
			currentStart = time
		} else if (split[2] == "falls") {
			// currentEnd = time
			// var interval = Interval{start:currentStart, end:currentEnd}
			// guards[currentGuard] = append(guards[currentGuard], interval)
			currentStart = time
		}
	}

	var max int64 = 0;
	var maxGuard int64 = 0;
	for guard, intervals := range guards {
		var sum int64 = 0
		for _, interval := range intervals {
			sum += (interval.end - interval.start)
		}
		fmt.Println("SUM ", sum, " guard ", guard)
		if (sum > max) {
			max = sum
			maxGuard = guard
		}
	}
	fmt.Println(max, " max ", maxGuard)
	var heatMap = make([]int, 120)

	for _, interval := range guards[maxGuard] {
		for i := interval.start; i < interval.end; i++ {
			heatMap[i-1420]++;
		}
		fmt.Println(interval)
	}

	var maxHeat = 0;
	var maxIdx = 0;
	for idx, heat := range heatMap {
		if (heat > maxHeat) {
			maxIdx = idx;
			maxHeat = heat;
		}
	}
	fmt.Println(maxIdx + 1420, " max ", maxHeat)

	fmt.Println( maxGuard * int64(maxIdx-1440+1420))

	
	// var sum = 0;
	// for i:=0; i < 1000; i++ {
	// 	for j:=0; j < 1000; j++ {
	// 		if (felt[i][j] > 1) {
	// 			sum += 1;
	// 		}
	// 	}
	// }
	// fmt.Println(sum);
}