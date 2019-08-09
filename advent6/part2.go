package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
)

type Point struct {
	x int
	y int
	name string
}

type DisPoint struct {
	point Point
	distance int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lowestX = 9999999;
	var highestX = 0;
	var lowestY = 9999999;
	var highestY = 0;

	var points = make([]Point, 0)

	scanner := bufio.NewScanner(file)

	// var DEFAULT_POINT = Point{x:0, y:0, name:"."}
	var name = []byte{'A'};
	for scanner.Scan() {
		var line = scanner.Text()
		var split = strings.Split(line, ", ")
		var tempx, _ = strconv.ParseInt(split[0], 10, 32)
		var tempy, _ = strconv.ParseInt(split[1], 10, 32)
		var x = int(tempx)
		var y = int(tempy)
		var point = Point{x:x, y:y, name:string(name)}
		name[0]++
		points = append(points, point)
		if (x < lowestX) {
			lowestX = x;
		}
		if (x > highestX) {
			highestX = x;
		}
		if (y < lowestY) {
			lowestY = y;
		}
		if (y > highestY) {
			highestY = y;
		}
	}
	var fieldHeight =  highestY - lowestY
	var fieldWidth = highestX - lowestX
	var field = make([][]DisPoint, fieldHeight)
	for i := 0; i < fieldHeight; i++ {
		field[i] = make([]DisPoint, fieldWidth)
	}

	var count = 0
	for row := 0; row < fieldHeight; row++ {
		for col := 0; col < fieldWidth; col++ {
			var sum = 0
			for _, point := range points {
				var realX = (point.x-lowestX) - col
				if (realX < 0) {
					realX = -realX
				}
				var realY = (point.y-lowestY) - row
				if (realY < 0) {
					realY = -realY
				}
				var total = realX + realY
				sum += total
			}
			if (sum < 10000) {
				count++
			}
		}
	}
	fmt.Println(count)

}