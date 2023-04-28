package main

func solve(list1 []int, list2 []int) []int {
	i, j := 0, 0
	n := (len(list1) + len(list2))
	result := make([]int, n)
	x1, x2 := 0, 0
	for (i + j) < n {
		if i < len(list1) {
			x1 = list1[i]
		}
		if j < len(list1) {
			x2 = list2[j]
		}
		if x1 < x2 {
			result[i+j] = x1
			i++
		} else {
			result[i+j] = x2
			j++
		}
	}

	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func solveOnLeetcode(list1 *ListNode, list2 *ListNode) *ListNode {

	var head *ListNode
	var tail *ListNode

	head1 := list1
	head2 := list2

	hasNext := func() bool {
		return head1 != nil || head2 != nil
	}

	next := func() *ListNode {
		var curr *ListNode
		if head1 == nil {
			curr = head2
			head2 = head2.Next
			return curr
		} else if head2 == nil {
			curr = head1
			head1 = head1.Next
			return curr
		} else {
			if head1.Val < head2.Val {
				curr = head1
				head1 = head1.Next
				return curr
			} else {
				curr = head2
				head2 = head2.Next
				return curr
			}
		}
	}

	insert := func(v int) {
		if head == nil {
			head = &ListNode{v, nil}
			tail = head
		} else {
			tail.Next = &ListNode{v, nil}
			tail = tail.Next
		}
	}

	for hasNext() {
		insert(next().Val)
	}

	return head
}
