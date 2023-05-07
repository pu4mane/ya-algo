package main

import "fmt"

func Solution(head *ListNode) {
	for head.next != nil {
		fmt.Println(head.data)
		head = head.next
	}
	fmt.Println(head.data)
}
