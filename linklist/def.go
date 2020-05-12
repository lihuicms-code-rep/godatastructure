package linklist

import "fmt"

//以数据域为int为例
type ListNode struct {
	Val int
	Next *ListNode
}

func NewListNode(Val int) *ListNode {
	return &ListNode{
		Val: Val,
		Next: nil,
	}
}

func (l *ListNode) Print() {
	fmt.Println("*****当前链表情况******")
	for p := l; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Val)
	}
	fmt.Println()
}
