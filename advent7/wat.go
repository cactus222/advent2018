package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
)


func main() {
	file, err := os.Open("input.txt")
	var numElements = 26

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var possibleRoots = make([][]byte, numElements)
	
	for i := 0; i < numElements; i++ {
		possibleRoots[i] = make([]byte, 0)
	}

	for scanner.Scan() {
		var line = scanner.Text()
		var split = strings.Split(line, " ")
		var parent = split[1]
		var child = split[7]
		
		possibleRoots[child[0]-'A'] = append(possibleRoots[child[0]-'A'], parent[0] - 'A')
		
		fmt.Println(parent, " " , child)
	}
	var output = ""

	for i := 0; i < numElements; i++ {
		for idx, dependencies := range possibleRoots {
			if (len(dependencies) == 0) {
				output += string(idx + 'A')
				//error this one out
				dependencies = append(dependencies, '.')
				possibleRoots[idx] = dependencies
				//remove dependency
				for idx2, children := range possibleRoots {
					for valIndex, val := range children {
						if (val == byte(idx)) {
							//swap with last and then set
							children[len(children)-1], children[valIndex] = children[valIndex], children[len(children)-1]
							children = children[:len(children)-1]
							// fmt.Println(children)
							possibleRoots[idx2] = children
							break;
						}
					}
				}
				break;
			}
		}
		// for idx, dependencies := range possibleRoots {
		// 	fmt.Println(dependencies, " " , string(idx + 'A'))
		// }
	}
	fmt.Println(output)

	
}