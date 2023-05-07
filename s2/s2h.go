package main

import "fmt"

type Stack struct {
	items []rune
}

func NewStack() *Stack {
	return &Stack{items: []rune{}}
}

func (s *Stack) push(item rune) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() rune {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem
}

func (s *Stack) peek() rune {
	lastIndex := len(s.items) - 1
	return s.items[lastIndex]
}

func (s *Stack) empty() bool {
	return len(s.items) == 0
}

func isCorrectBracketSeq(line string) bool {
	s := NewStack()
	for _, c := range line {
		if c == '[' || c == '(' || c == '{' {
			s.push(c)
		} else {
			if s.empty() {
				return false
			} else {
				peek := s.peek()
				if peek == '[' && c == ']' || peek == '(' && c == ')' || peek == '{' && c == '}' {
					s.pop()
				}
			}
		}
	}
	return s.empty()
}

func main() {
	var line string
	fmt.Scan(&line)

	if isCorrectBracketSeq(line) {
		fmt.Print("True")
	} else {
		fmt.Print("False")
	}
}