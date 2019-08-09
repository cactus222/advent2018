package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
)

type Task struct {
	taskName byte
	timeLeft byte

}

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
	var doneTaskCount = 0;
	var seconds = 0;

	var numWorkers = 5
	var extraTime byte = 61

	var DEFAULT_TASK = Task{128, 128}
	var workers = make([]Task, numWorkers)

	var workable = findWorkable(possibleRoots)
	var initialIndex = 0

	for ; initialIndex < len(workable); initialIndex++ {
		fmt.Println("Start job ", workable[initialIndex])
		workers[initialIndex] = Task{taskName: workable[initialIndex], timeLeft: workable[initialIndex] + extraTime}
		possibleRoots[workable[initialIndex]] = append(possibleRoots[workable[initialIndex]], '.')
	}
	workable = make([]byte, 0)


	for ; initialIndex < numWorkers; initialIndex++ {
		workers[initialIndex] = DEFAULT_TASK
	}

	for ; doneTaskCount < numElements; {
		
		seconds++
		// fmt.Println(workers)
		var doneTasks = make([]Task, 0)
		for idx, worker := range workers {
			if (worker.taskName != 128) {
				worker.timeLeft--
				workers[idx] = worker
				if (worker.timeLeft == 0) {
					doneTasks = append(doneTasks, worker)
					workers[idx] = DEFAULT_TASK
					doneTaskCount++
				}
			}
		}
		//something finished
		if (len(doneTasks) > 0) {
			for _, doneTask := range doneTasks {
				output = markFinished(possibleRoots, doneTask, output)
			}
		
			workable = findWorkable(possibleRoots)
			for _, work := range workable {
				for idx, worker := range workers {
					if (worker.taskName == 128) {
						fmt.Println("Start job ", work)
						workers[idx] = Task{taskName: work, timeLeft: work + extraTime}
						possibleRoots[work] = append(possibleRoots[work], '.')
						break;
					}
				}
			}

			// reader := bufio.NewReader(os.Stdin)
			// fmt.Print("Enter text: ")
			// reader.ReadString('\n')
		}

		// for idx, dependencies := range possibleRoots {
		// 	if (len(dependencies) == 0) {
				
		// 		// output += string(idx + 'A')
		// 		// //error this one out
		// 		// dependencies = append(dependencies, '.')
		// 		// possibleRoots[idx] = dependencies
		// 		// removeDependency(possibleRoots, idx)
		// 		// break;
		// 	}
		// }
		// for idx, dependencies := range possibleRoots {
		// 	fmt.Println(dependencies, " " , string(idx + 'A'))
		// }
	}
	fmt.Println(output)
	fmt.Println(seconds)
}

func findWorkable(possibleRoots [][]byte) []byte {
	var workable = make([]byte, 0)
	for idx, dependencies := range possibleRoots {
		if (len(dependencies) == 0) {
			workable = append(workable, byte(idx))
		}
	}
	return workable
}

func markFinished(possibleRoots [][]byte, task Task, output string) string {
	output += string(task.taskName + 'A')
	// possibleRoots[task.taskName] = append(possibleRoots[task.taskName], '.')
	// for idx, dependencies := range possibleRoots {

			// possibleRoots[idx] = dependencies
	removeDependency(possibleRoots, task.taskName)
			// break;
		// }
	// }

	return output
	
}

func removeDependency(possibleRoots [][]byte, removedNode byte) {
	for idx, children := range possibleRoots {
		for valIndex, val := range children {
			if (val == removedNode) {
				children[len(children)-1], children[valIndex] = children[valIndex], children[len(children)-1]
				children = children[:len(children)-1]
				possibleRoots[idx] = children
				break;
			}
		}
	}
}