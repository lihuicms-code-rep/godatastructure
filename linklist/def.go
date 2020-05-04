package linklist

import "fmt"

//以数据域为int为例
type LNode struct {
	Data int
	Next *LNode
}

func NewLNode(data int) *LNode {
	return &LNode{
		Data: data,
		Next: nil,
	}
}

func (l *LNode) Print() {
	fmt.Println("*****当前链表情况******")
	for p := l; p != nil; p = p.Next {
		fmt.Printf("%d->", p.Data)
	}
	fmt.Println()
}
