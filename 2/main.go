package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	var add int
	for l1 != nil || l2 != nil {
		var vl2, vl1 int
		if l1 == nil {
			vl1 = 0
		} else {
			vl1 = l1.Val
		}

		if l2 == nil {
			vl2 = 0
		} else {
			vl2 = l2.Val
		}

		v := vl1 + vl2 + add
		if v > 9 {
			cur.Val = v % 10
			add = 1
		} else {
			cur.Val = v
			add = 0
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1 != nil || l2 != nil {
			newNode := &ListNode{}
			cur.Next = newNode
			cur = newNode
		}
	}

	if add > 0 {
		newNode := &ListNode{}
		newNode.Val = 1
		cur.Next = newNode
	}

	return ret
}

func main() {
	l11 := &ListNode{3, nil}
	l12 := &ListNode{4, l11}
	l13 := &ListNode{2, l12}

	l21 := &ListNode{4, nil}
	l22 := &ListNode{6, l21}
	l23 := &ListNode{5, l22}

	r := addTwoNumbers(l13, l23)

	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
