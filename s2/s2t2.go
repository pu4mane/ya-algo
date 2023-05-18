/*
https://contest.yandex.ru/contest/22781/run-report/87457479/
-- ПРИНЦИП РАБОТЫ --
Я реализовал калькулятор на стеке.
Функция принимает на вход строку - выражение, записанное в обратной польской нотации и парсит ее на слайс токенов(операнды и операторы)
Далее проходим по слайсу токенов, если встречается операнд, то помещаем его в вершину стека,
если встречается оператор, то из вершины стека извлекаются два операнда и производится операция над ними используя текущий оператор.
Результат выполненной операции помещается на вершину стека.
После обработки входного набора токенов результат вычисления выражения находится в вершине стека - его мы и выводим.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Проход по слайсу токенов стоит О(n), n - колличество токенов.
Добавление и извлечение из стека стоит O(1).
Временная сложность будет равна O(n).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пространственная сложность будет равна 0(n), n - колличество токенов, каждый токен занимает k памяти.
*/

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

func (s *Stack) Push(item int) {
	s.stack = append(s.stack, item)
	s.size++
}

func (s *Stack) Pop() (int, bool) {
	if s.size == 0 {
		return 0, false
	}
	itemIdx := s.size - 1
	item := s.stack[itemIdx]
	s.stack = s.stack[:itemIdx]
	s.size--
	return item, true
}

func Calc(expression string) int {
	m := map[string]func(int, int) int{
		"+": func(a int, b int) int { return b + a },
		"-": func(a int, b int) int { return b - a },
		"*": func(a int, b int) int { return b * a },
		"/": func(a int, b int) int { return int(math.Floor(float64(b) / float64(a))) },
	}
	stack := NewStack()
	tokens := strings.Split(expression, " ")
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
		} else if operation, ok := m[token]; ok && stack.size > 1 {
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()
			stack.Push(operation(num1, num2))
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