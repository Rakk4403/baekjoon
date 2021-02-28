package main

import (
    "fmt"
    "strings"
    "strconv"
    "bufio"
    "os"
)

func main() {
    var count int
    _, _ = fmt.Scanf("%d", &count)

    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
text = strings.Replace(text, "\n", "", -1)
	fmt.Println(strings.Replace(text, "\n", "", -1))
    splits := strings.Split(text, " ")
	fmt.Println(splits)
    
    numList := []int{}
    for _, splitStr := range splits {
        num, _ := strconv.Atoi(splitStr)
        numList = append(numList, num)
    }
   max := numList[0]
   min := numList[0]
   for _, num := range numList {
       if num > max {
            max = num
        }
        if num < min {
            min = num
        }
    }
    fmt.Printf("%d %d", min, max)
}
