package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    text := scanner.Text()
    
    splitStr := strings.Split(text, " ")
    // n, _ := strconv.Atoi(splitStr[0])
    m, _ := strconv.Atoi(splitStr[1])
    
    scanner.Scan()
    nn := scanner.Text()
    nums := strings.Split(nn, " ")
    var numList []int
    for _, numStr := range nums {
        num, _ := strconv.Atoi(numStr)
        numList = append(numList, num)
    }
   
    minVal := m
    target := 0
    for i := 0; i < len(numList)-2; i++ {
        for j := i+1; j < len(numList) - 1; j++ {
            for k := j + 1; k < len(numList); k++ {
                sum := numList[i] + numList[j] + numList[k]
		if m - sum < minVal && m >= sum {
		    minVal = m - sum
		    target = sum
                }
            }
        }
    }    
    fmt.Println(target)
}
