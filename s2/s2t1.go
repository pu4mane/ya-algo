package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Deque struct {
	deque []int
	maxM  int
	front int
	back  int
	size  int
}

func NewDeque(m int) *Deque {
	return &Deque{
		deque: make([]int, m),
		maxM:  m,
		front: 0,
		back:  0,
		size:  0,
		}
}

func (d *Deque) PushBack(item int) bool {
	if d.size == d.maxM {
		return false
	} else {
		d.deque[d.back] = item
		d.back = (d.back + 1) % d.maxM
		d.size++
		return true
	}
}

func (d *Deque) PushFront(item int) bool {
	if d.size == d.maxM {
		return false
	} else {
		d.front = (d.front - 1 + d.maxM) % d.maxM
		d.deque[d.front] = item
		d.size++
		return true
	}
}

func (d *Deque) PopFront() (int, bool) {
	if d.size == 0 {
		return 0, false
	} else {
		item := d.deque[d.front]
		d.front = (d.front + 1) % d.maxM
		d.size--
		return item, true
	}
}

func (d *Deque) PopBack() (int, bool) {
	if d.size == 0 {
		return 0, false
	} else {
		d.back = (d.back - 1 + d.maxM) % d.maxM
		item := d.deque[d.back]
		d.size--
		return item, true
	}
}

func RunCommands(commands []string, deque *Deque) {
	for _, command := range commands {
		cmd := strings.Split(command, " ")
		switch cmd[0] {
		case "push_back":
			item, _ := strconv.Atoi(cmd[1])
			if !deque.PushBack(item) {
				fmt.Println("error")
			}
			case "push_front":
				item, _ := strconv.Atoi(cmd[1])
				if !deque.PushFront(item) {
					fmt.Println("error")
				}
				case "pop_front":
					item, ok := deque.PopFront()
					if ok {
						fmt.Println(item)
					} else {
						fmt.Println("error")
					}
					case "pop_back":
						item, ok := deque.PopBack()
						if ok {
							fmt.Println(item)
						} else {
							fmt.Println("error")
						}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	commands := make([]string, n)

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	deque := NewDeque(m)

	for i := 0; i < n; i++ {
		scanner.Scan()
		commands[i] = scanner.Text()
	}

	RunCommands(commands, deque)
}