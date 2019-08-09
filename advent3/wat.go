package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var felt = make([][]int8 ,1000)
	for idx := 0; idx < 1000; idx++ {
		felt[idx] = make([]int8, 1000)
	}
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		fmt.Println(line)
		var split = strings.Split(line, " ")
		var col, _ = strconv.ParseInt(strings.Split(split[2], ",")[0], 10, 32)
		var row, _ = strconv.ParseInt(strings.TrimRight(strings.Split(split[2], ",")[1], ":"), 10, 32)
		var width, _ = strconv.ParseInt(strings.Split(split[3], "x")[0], 10, 32)
		var height, _ = strconv.ParseInt(strings.Split(split[3], "x")[1], 10, 32)
	
		for i := 0; i < int(height); i++ {
			for j := 0; j < int(width); j++ {
				var rowNum = int(row) + i;
				var colNum = int(col) + j;
				felt[rowNum][colNum] += 1; 
			}
		}
	}
	var sum = 0;
	for i:=0; i < 1000; i++ {
		for j:=0; j < 1000; j++ {
			if (felt[i][j] > 1) {
				sum += 1;
			}
		}
	}
	fmt.Println(sum);
}