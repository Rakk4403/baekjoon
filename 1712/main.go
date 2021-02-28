package main

import (
    "fmt"
)

func main() {
    var a, b, c int
    _, _ = fmt.Scanf("%d %d %d", &a, &b, &c)
    
    x := a / (c - b)
    if x > 0 {
        fmt.Println(x)
    } else {
        fmt.Println(-1)
    }
}
