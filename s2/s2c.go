package main

func Solution(node *ListNode, idx int) *ListNode {
	getNodeByIndex := func(node *ListNode, index int) *ListNode {
		for index > 0 {
			node = node.next
			index--
		}
		return node
	}
	if idx == 0 {
		node = node.next
	} else {
		previousNode := getNodeByIndex(node, idx-1)
		nextNode := getNodeByIndex(node, idx+1)
		previousNode.next = nextNode
	}

	return node
}