package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	stack []int
	size  int
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]int, 0),
		size:  0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	itemIdx := s.size - 1
	item := s.stack[itemIdx]
	s.stack = s.stack[:itemIdx]
	s.size--
	return item, true
}

func Calc(expression string) int {
	stack := NewStack()
	tokens := strings.Split(expression, " ")
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
		} else {
			num1, _ := stack.Pop()
			b := float64(num1)
			num2, _ := stack.Pop()
			a := float64(num2)
			switch token {
			case "+":
				res := a + b
				stack.Push(int(math.Floor(res)))
			case "-":
				res := a - b
				stack.Push(int(math.Floor(res)))
			case "*":
				res := a * b
				stack.Push(int(math.Floor(res)))
			case "/":
				res := a / b
				stack.Push(int(math.Floor(res)))
			}
		}
	}
	result, _ := stack.Pop()
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expression := scanner.Text()
	fmt.Print(Calc(expression))
}
