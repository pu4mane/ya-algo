/*
https://contest.yandex.ru/contest/22781/run-report/87361848/
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
	stack := NewStack()
	tokens := strings.Split(expression, " ")
	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack.Push(num)
		} else {
			num1, _ := stack.Pop()
			num2, _ := stack.Pop()
			switch token {
			case "+":
				res := float64(num2) + float64(num1)
				stack.Push(int(math.Floor(res)))
				case "-":
					res := float64(num2) - float64(num1)
					stack.Push(int(math.Floor(res)))
					case "*":
						res := float64(num2) * float64(num1)
						stack.Push(int(math.Floor(res)))
						case "/":
							res := float64(num2) / float64(num1)
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