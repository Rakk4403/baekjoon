package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func main() {
    var count int
    _, _ = fmt.Scanf("%d", &count)
    
    reader := bufio.NewScanner(os.Stdin)
    reader.Scan()
    text := reader.Text()
    numList := []int{}
    splitText := strings.Split(text, " ")
    for _, t := range splitText {
        num, _ := strconv.Atoi(t)
        numList = append(numList, num)
    }

    var maxNum int
    for _, t := range numList {
        if maxNum < t {
            maxNum = t
        }
    }
    
    numMap := []bool{}
    for i := 0; i <= maxNum*maxNum; i++ {
        numMap = append(numMap, true)
    }
    
    for idx := 2; idx < maxNum*maxNum; idx++ {
        for i := 2; i*idx < maxNum*maxNum; i++ {
            numMap[i*idx] = false
        }
    }
   numMap[0] = false
   numMap[1] = false
 
    var counter int
    for _, num := range numList {
        if numMap[num] == true {
	    counter++
        }
    }
    
    fmt.Println(counter)
}
