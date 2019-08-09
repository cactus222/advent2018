package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	next *Node
	prev *Node
	val int 
}

func main() {
	file, err := os.Open("miniinput.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	for scanner.Scan() {
		var line = scanner.Text()
		var split = strings.Split(line, " ")
		
		var numPlayers64,_ = strconv.ParseInt(split[0], 10, 32)
		var numMarbles64,_ = strconv.ParseInt(split[6], 10, 32)
		
		var numPlayers = int(numPlayers64)
		var numMarbles = int(numMarbles64)

		var players = make([]int, numPlayers)
		// var board = make([]int, numMarbles)
		//TODO just link root to tail so we dont hafta deal with junk
		var root = &Node{val:0}
		var third = &Node{val:1}
		root.next = &Node{prev: root, next:third, val:2}
		var currentNode = root

		var curPlayer = 2;
		printList(root)
		for idx := 3; idx <= numMarbles; idx++ {
			if (idx % 23 == 0) {
				players[curPlayer]+=idx

				//too lazy to do tail logic
				for i := 0; i < 6; i++ {
					currentNode = currentNode.prev
					if (currentNode == nil) {
						//go to end
						currentNode = root
						for ; currentNode.next != nil; {
							currentNode = currentNode.next
						}
					} 
				}
				
				players[curPlayer] += currentNode.val
				if (currentNode.prev != nil) {
					currentNode.prev.next = currentNode.next
					currentNode.next.prev = currentNode.prev
					currentNode = currentNode.prev
				} else {
					//wdf?
					fmt.Println("WE DED, todo loop back to end and change root")
				}
				fmt.Println(idx)
			} else {
				currentNode = currentNode.next
				if (currentNode == nil) {
					currentNode = root
				}
				currentNode = currentNode.next
				if (currentNode == nil) {
					currentNode = root
				}
				var newNode = &Node{prev:currentNode, next:currentNode.next, val:idx}
				if (currentNode.next != nil) {
					currentNode.next.prev = newNode
				}
				currentNode.next = newNode
				
			}
			curPlayer = (curPlayer+1) % numPlayers
			// fmt.Println(board)
	//	printList(root)
		}	
		
		var playerNum, score = getMax(players)
		fmt.Println("winner ", playerNum , " " , score)
	}
}

func printList(root *Node) {
	for ; root != nil; {
		fmt.Print(root.val, " ")
		root = root.next
	}
	fmt.Println()
}
func getMax(players []int) (idx int, val int) {
	idx = 0
	val = 0
	for i, score := range players {
		fmt.Println("player ", i , " score ", score)
		if (score > val) {
			val = score
			idx = i
		}
	}
	return idx, val
}