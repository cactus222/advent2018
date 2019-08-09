package main

import (
    "bufio"
    "fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var twos = 0;
	var threes = 0;
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		
		var freqMap = make(map[rune]int);
		for _, ch := range(line) {
			freqMap[ch]++;
		}

		var has2 = false;
		var has3 = false
		for ch, count  := range(freqMap) {
			if count == 2 {
				has2 = true;
				fmt.Println(line , " CH2 IS ", (ch-'a') );
			} else if (count == 3) {
				has3 = true;
				fmt.Println(line, " CH3 IS ", (ch-'a') );
			}
		} 
		if (has2) {
			twos++;
		}
		if has3 {
			threes++;
		}
	}
	fmt.Println(twos*threes)
	//fmt.Println("done", freq)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	
}