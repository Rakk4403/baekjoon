package main

import ("fmt")

type Item struct {
    a int
    b int
}

func main() {
    var itemList []Item
    for {
        item := Item{}
        _, _ = fmt.Scanf("%d %d", &item.a, &item.b)
        if (item.a == 0 && item.b == 0) {
            break
        }
        itemList = append(itemList, item)
    }
    
    for _, item := range itemList {
        println(item.a+item.b)
    }
}
