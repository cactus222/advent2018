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

	var caseDifference = 'a' - 'A'
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var outputArray = make([]rune, 0)
		var count = 0;

		for _, curChar := range(line) {
			var currentSize = len(outputArray)
			//Previous
			if (currentSize > 0) {
				if (curChar == outputArray[currentSize - 1] + caseDifference || 
					curChar == outputArray[currentSize - 1] - caseDifference) {
						//fmt.Println("match ", string(curChar), " ", string(outputArray[currentSize - 1]))
					//Trim last
					outputArray = outputArray[:currentSize  - 1]
					count--;
				} else {
					count++;
					outputArray = append(outputArray, curChar)
				}
			} else {
				count++;
				outputArray = append(outputArray, curChar)
			}
		
			//fmt.Println(string(outputArray))
		}
		fmt.Println(count)
		fmt.Println(outputArray)
		fmt.Println("size left", len(outputArray))
	}

}