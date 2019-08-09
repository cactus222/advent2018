package main

import (
    "bufio"
    "fmt"
	"os"
)


type trie struct {
	children []*trie
}

func (root *trie) addWord(word string) {
	var curNode = root;
	for _, letter := range(word) {
		var index = letter - 'a';
		if curNode.children[index] != nil {
			curNode = curNode.children[index]
		} else {
			curNode.children[index] = &trie{make([]*trie, 26)}
			curNode = curNode.children[index]
		}
	}
}

func addChar(word string, b byte) string {
	return string(append([]byte(word), b))
}

func (root *trie) contains1Err(word string, idx int, count int, wordSoFar string) (bool, string) {
	if (idx == len(word)) {
		return true, wordSoFar
	}
	if (count == 0) {
		for charIndex, child := range (root.children) {
			if (child != nil) {
				if ((word[idx] - 'a') == byte(charIndex)) {
					var success, word = child.contains1Err(word, idx+1, count, addChar(wordSoFar, byte(charIndex+'a')) )
					if (success) {
						return success, word
					}	
				} else {
					var success, word = child.contains1Err(word, idx+1, count+1,  addChar(wordSoFar, byte(charIndex+'a')) )
					if (success) {
						return success, word
					}
				}
			}
		}
	} else if (count == 1) {
		var child = root.children[word[idx] - 'a']
		if child != nil {
			var success, word = child.contains1Err(word, idx+1, count, addChar(wordSoFar,  word[idx]))
			if (success) {
				return success, word
			}
		}
	}
	return false, ""
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	
	var root = trie{make([]*trie, 26)};

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line = scanner.Text()
		var success, word = root.contains1Err(line, 0, 0, "")

		if (success) {
			fmt.Println(word)
			fmt.Println(line);
			break
		} else {
			root.addWord(line)
		}
	}

	

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	
}