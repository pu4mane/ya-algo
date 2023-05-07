package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	item int
	next *Node
}

func NewNode() *Node {
	return &Node{0, nil}
}

type ListQueue struct {
	head *Node
	tail *Node
	size int
}

func NewListQueue() *ListQueue {
	return &ListQueue{NewNode(), NewNode(), 0}
}

func (q *ListQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *ListQueue) Put(item int, err error) {
	if q.IsEmpty() {
		q.tail = &Node{item, nil}
		q.head = q.tail
	} else {
		q.tail.next = &Node{item, nil}
		q.tail = q.tail.next
	}
	q.size++
}

func (q *ListQueue) Get() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	} else if q.size == 1 {
		item := q.head
		q.head = NewNode()
		q.tail = NewNode()
		q.size--
		return item.item, true
	} else {
		item := q.head
		q.head = q.head.next
		q.size--
		return item.item, true
	}
}

func (q *ListQueue) Size() int {
	return q.size
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())

	lq := NewListQueue()

	var result []string

	for i := 0; i < n; i++ {
		scanner.Scan()
		command := strings.Split(scanner.Text(), " ")
		if command[0] == "put" {
			lq.Put(strconv.Atoi(command[1]))
		} else if command[0] == "get" {
			item, res := lq.Get()
			if res {
				result = append(result, strconv.Itoa(item))
			} else {
				result = append(result, "error")
			}
		} else if command[0] == "size" {
			result = append(result, strconv.Itoa(lq.Size()))
		}
	}

	for _, elem := range result {
		fmt.Println(elem)
	}
}