package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)

	var numList []int
	for i := 0; i < n; i++ {
		scanner.Scan()
		mStr := scanner.Text()
		m, _ := strconv.Atoi(mStr)
		numList = append(numList, m)
	}

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if numList[i] > numList[j] {
				tmp := numList[i]
				numList[i] = numList[j]
				numList[j] = tmp
			}
		}
	}

	for _, num := range numList {
		fmt.Println(num)
	}
}
