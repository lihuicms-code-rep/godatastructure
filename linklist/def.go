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


//多级双向链表
type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}

//根据输入创建多级双向链表
//数据元素为-1,表示nil
//[1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]


