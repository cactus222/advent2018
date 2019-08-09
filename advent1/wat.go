package main

import (
    "bufio"
    "fmt"
	"os"
	"strconv"
)

func main() {


	var freq int64 = 0;
	var seen = make(map[int64]bool);
	for ;; {
		file, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var line = scanner.Text()
			var sign int64 = 1
	
			if (line[0] == '-') {
				sign = -1
			}
			var value, _ = strconv.ParseInt(line[1:], 10, 32)
			freq += (value * sign)
			fmt.Println(freq);
			if _, exists := seen[freq]; exists {
				fmt.Println("again", freq);
				goto DONE;
			} else {
				seen[freq] = true;
			}
		}
		//fmt.Println("done", freq)
	
		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	DONE: 
		fmt.Println("again", freq)
	
}