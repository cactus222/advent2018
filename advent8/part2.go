package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	children []*Node
	metadata []int
	numChildren int
	numMeta int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var idx = 0;
	var root *Node;
	for scanner.Scan() {
		var line = scanner.Text()
		var split = strings.Split(line, " ")
		root, idx = parseNode(split, 0)

		//var sum = 0
		var stack = make([]*Node, 0)
		stack = append(stack, root)
		for ; idx < len(split); {
			var lastNode = stack[len(stack) - 1]
			// fmt.Println("STACK ", stack)
			// fmt.Println("lastnode", lastNode)
			// fmt.Println("curridx", idx)
			if (lastNode.numChildren - len(lastNode.children) == 0) {
				// fmt.Println("mathced")

				//Add metadata
				var metaArray = make([]int, lastNode.numMeta)
				for i := 0; i < lastNode.numMeta; i++ {
					var meta int
					meta, idx = parseMeta(split, idx)
					metaArray[i] = meta
					//sum += meta
				}
				// fmt.Println(sum)
				lastNode.metadata = metaArray

				stack = stack[:len(stack)-1]
				// fmt.Println(stack[len(stack) - 1])
				// fmt.Println("stack pop", stack)

			} else {
				// fmt.Println("diff", lastNode.numChildren - len(lastNode.children))
				// parse another node
				var node *Node;
				node, idx = parseNode(split, idx)
				
				lastNode.children = append(lastNode.children, node)
				// fmt.Println("appended", lastNode)
				stack = append(stack, node)

			}
			// reader := bufio.NewReader(os.Stdin)
			// //fmt.Print("Enter text: ")
			// reader.ReadString('\n')
		}
		//fmt.Println(sum)
	}
	
	//traverse
	var ans = traverse(root)
	
	fmt.Println(ans)



	
}

func traverse (node *Node) int {
	var sum = 0;
	fmt.Println(node)
	if (node.numChildren == 0) {
		
		for _, val := range node.metadata {
			sum += val
		}
		fmt.Println("yello", sum)
	} else {
		for i := 0; i < node.numMeta; i++ {
			var childIndex = node.metadata[i]
			if (childIndex <= node.numChildren && childIndex != 0) {
				sum += traverse(node.children[childIndex-1])
			}
		}
	} 
	
	return sum
}

func parseMeta(data []string, idx int) (int, int) {
	var meta,_ = strconv.ParseInt(data[idx], 10, 32)
	return int(meta), idx+1
}

func parseNode(data []string, idx int) (*Node, int) {
	var rootNumChildren,_ = strconv.ParseInt(data[idx], 10, 32)
	var rootNumMeta,_ = strconv.ParseInt(data[idx+1], 10, 32)
	return &Node{numChildren: int(rootNumChildren), numMeta: int(rootNumMeta)}, idx+2
}