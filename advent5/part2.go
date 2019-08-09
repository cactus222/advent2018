package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
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
		
	
		var shortest = 999999;

		for i := 0; i < 26; i++ {
			var replaceLarge = string([]byte{'a' + byte(i)})
			var replaceSmall = string([]byte{'A' + byte(i)})
			var newString = strings.Replace(strings.Replace(line, replaceLarge, "", -1), replaceSmall, "", -1)

			// var count = 0;
			var outputArray = make([]rune, 0)
			for _, curChar := range(newString) {
				var currentSize = len(outputArray)
				//Previous
				if (currentSize > 0) {
					if (curChar == outputArray[currentSize - 1] + caseDifference || 
						curChar == outputArray[currentSize - 1] - caseDifference) {
							//fmt.Println("match ", string(curChar), " ", string(outputArray[currentSize - 1]))
						//Trim last
						outputArray = outputArray[:currentSize  - 1]
						
					} else {
						outputArray = append(outputArray, curChar)
					}
				} else {
					outputArray = append(outputArray, curChar)
				}
			
				//fmt.Println(string(outputArray))
			}
			fmt.Println(outputArray)
			fmt.Println("size left", len(outputArray))
			if (len(outputArray) < shortest) {
				shortest = len(outputArray)
			}
		}
		fmt.Println(shortest)

	}

}