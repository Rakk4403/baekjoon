package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// stack
type Stack struct {
	Items []int
}

func (s *Stack) Push(value int) {
	s.Items = append(s.Items, value)
}

func (s *Stack) Pop() int {
	if len(s.Items) == 0 {
		return -1
	}
	last := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return last
}

func (s *Stack) Top() int {
	if len(s.Items) == 0 {
		return -1
	}
	return s.Items[len(s.Items)-1]
}

func (s *Stack) Size() int {
	return len(s.Items)
}

func NewStack() *Stack {
	return &Stack{}
}

func main() {
	s := NewStack()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	count, _ := strconv.Atoi(text)
	cmdList := []string{}
	for i := 0; i < count; i++ {
		scanner.Scan()
		text := scanner.Text()
		cmdList = append(cmdList, text)
	}
	for _, cmd := range cmdList {
		if cmd == "top" {
			fmt.Println(s.Top())
		} else if cmd == "size" {
			fmt.Println(s.Size())
		} else if cmd == "empty" {
			if s.Size() == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		} else if cmd == "pop" {
			fmt.Println(s.Pop())
		} else {
			// push x
			splitCmd := strings.Split(cmd, " ")
			pushVal, _ := strconv.Atoi(splitCmd[1])
			s.Push(pushVal)
		}
	}
}
