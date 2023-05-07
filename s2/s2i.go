package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	queue []int
	maxN  int
	head  int
	tail  int
	size  int
}

func NewQueue(n int) *Queue {
	return &Queue{
		queue: make([]int, n),
		maxN:  n,
		head:  0,
		tail:  0,
		size:  0,
	}
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Push(item int) bool {
	if q.size != q.maxN {
		q.queue[q.tail] = item
		q.tail = (q.tail + 1) % q.maxN
		q.size++
		return true
	}
	return false
}

func (q *Queue) Pop() (int, bool) {
	if q.IsEmpty() {
		return -1, false
	}
	item := q.queue[q.head]
	q.queue[q.head] = 0
	q.head = (q.head + 1) % q.maxN
	q.size--
	return item, true
}

func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	item := q.queue[q.head]
	return item, true
}

func (q *Queue) Size() int {
	return q.size
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	q := NewQueue(n)
	var result []string
	for i := 0; i < c; i++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")
		if command[0] == "push" {
			item, _ := strconv.Atoi(command[1])
			if !q.Push(item) {
				result = append(result, "error")
			}
		} else if command[0] == "pop" {
			item, res := q.Pop()
			if !res {
				result = append(result, "None")
			} else {
				result = append(result, strconv.Itoa(item))
			}
		} else if command[0] == "peek" {
			item, res := q.Peek()
			if !res {
				result = append(result, "None")
			} else {
				result = append(result, strconv.Itoa(item))
			}
		} else if command[0] == "size" {
			result = append(result, strconv.Itoa(q.Size()))
		}
	}
	for _, r := range result {
		fmt.Println(r)
	}
}
