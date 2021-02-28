package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	Items []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val int) {
	q.Items = append(q.Items, val)
}

func (q *Queue) Pop() int {
	if len(q.Items) == 0 {
		return -1
	}
	first := q.Items[0]
	q.Items = q.Items[1:]
	return first
}

func (q *Queue) Size() int {
	return len(q.Items)
}

func (q *Queue) Front() int {
	if len(q.Items) == 0 {
		return -1
	}
	return q.Items[0]
}

func (q *Queue) Back() int {
	if len(q.Items) == 0 {
		return -1
	}
	return q.Items[(len(q.Items) - 1)]
}

func main() {
	q := NewQueue()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	countStr := scanner.Text()
	count, _ := strconv.Atoi(countStr)
	cmdList := []string{}
	for i := 0; i < count; i++ {
		scanner.Scan()
		cmd := scanner.Text()
		cmdList = append(cmdList, cmd)
	}

	for _, cmd := range cmdList {
		if cmd == "front" {
			fmt.Println(q.Front())
		} else if cmd == "back" {
			fmt.Println(q.Back())
		} else if cmd == "size" {
			fmt.Println(q.Size())
		} else if cmd == "empty" {
			if q.Size() == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		} else if cmd == "pop" {
			fmt.Println(q.Pop())
		} else {
			splitStr := strings.Split(cmd, " ")
			value, _ := strconv.Atoi(splitStr[1])
			q.Push(value)
		}
	}
}
