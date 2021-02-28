package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numList []int

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	// n, _ := strconv.Atoi(text)

	scanner.Scan()
	text = scanner.Text()
	textSplits := strings.Split(text, " ")

	for _, splitText := range textSplits {
		rawNum, _ := strconv.Atoi(splitText)
		numList = append(numList, rawNum)
	}

	for i, pick := range numList {
	}
	
	
	scanner.Scan()
	text = scanner.Text()
	// m, _ := strconv.Atoi(text)

	scanner.Scan()
	text = scanner.Text()
	textSplits = strings.Split(text, " ")
	for _, splitText := range textSplits {
		rawNum, _ := strconv.Atoi(splitText)
		fmt.Println(tree.Search(rawNum))
	}
}
