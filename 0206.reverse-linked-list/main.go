package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	n5 := &ListNode{Val: 5}
	n4 := &ListNode{Val: 4, Next: n5}
	n3 := &ListNode{Val: 3, Next: n4}
	n2 := &ListNode{Val: 2, Next: n3}
	n1 := &ListNode{Val: 1, Next: n2}
	n := reverseList(n1)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
