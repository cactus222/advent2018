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

	var felt = make([][]int ,1000)
	for idx := 0; idx < 1000; idx++ {
		felt[idx] = make([]int, 1000)
	}
	var duped = make([]bool , 1254)
	scanner := bufio.NewScanner(file)
	var claimNum int = 1;
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
				var target = felt[rowNum][colNum];
				if (target != 0) {
					//fmt.Println(target)
					duped[target] = true
					duped[claimNum] = true
				}
				felt[rowNum][colNum] = claimNum
			}
		}
		claimNum++;
	}
	for i:=1; i < 1254; i++ {
		if (!duped[i]) {
			fmt.Println(i)
		}
	}
}