package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var white int = 0
var blue int = 0
var rect [][]int

func f(a, b, n, color int) {
	fmt.Println(a, b, n, color)
	if n == 0 {
		return
	}
	flag := false
	base := color
	for i := 0; i < n; i++ {
		j := 0
		for j = 0; j < n; j++ {
			if rect[a+i][b+j] != base {
				flag = true
				break
			}
		}
		if (j != n) {
			break
		}
	}

	if flag == false {
		if color == 0 {
			white++
		} else {
			blue++
		}
	} else {
		n /= 2
		f(a, b, n, color);
		f(a, b+n, n, color);
		f(a+n, b, n, color);
		f(a+n, b+n, n, color);
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nStr := scanner.Text()
	n, _ := strconv.Atoi(nStr)


	for i := 0; i < n; i++ {
		scanner.Scan()
		lineStr := scanner.Text()
		lineSplits := strings.Split(lineStr, " ")
		rect = append(rect, []int{})
		for _, itemStr := range lineSplits {
			item, _ := strconv.Atoi(itemStr)
			rect[i] = append(rect[i], item)
		}
	}

	f(0,0,n,1)
	f(0,0,n,0)
	fmt.Println(white)
	fmt.Println(blue)
}
