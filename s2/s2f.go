package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StackMax struct {
	items []int
	maxs  []int
}

func NewStackMax() *StackMax {
	return &StackMax{items: make([]int, 0), maxs: make([]int, 0)}
}

func (s *StackMax) Push(item int) {
	s.items = append(s.items, item)
	if len(s.maxs) == 0 || item >= s.maxs[len(s.maxs)-1] {
		s.maxs = append(s.maxs, item)
	}
}

func (s *StackMax) Pop() (int, bool) {
	if len(s.items) == 0 {
		return -1, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	if item == s.maxs[len(s.maxs)-1] {
		s.maxs = s.maxs[:len(s.maxs)-1]
	}
	return item, true
}

func (s *StackMax) GetMax() (int, bool) {
	if len(s.maxs) == 0 {
		return -1, false
	}
	return s.maxs[len(s.maxs)-1], true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	s := NewStackMax()
	var result []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")
		if command[0] == "push" {
			item, _ := strconv.Atoi(command[1])
			s.Push(item)
		} else if command[0] == "pop" {
			_, res := s.Pop()
			if res == false {
				result = append(result, "error")
			}
		} else if command[0] == "get_max" {
			max, res := s.GetMax()
			if res == false {
				result = append(result, "None")
			} else {
				result = append(result, strconv.Itoa(max))
			}
		}
	}
	for _, r := range result {
		fmt.Println(r)
	}
}