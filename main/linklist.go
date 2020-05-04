package main

import (
	"fmt"
	"github.com/lihuicms-code-rep/godatastructure/linklist"
)

//链表部分测试
func main() {
	//fmt.Println("*************第一种方法**************")
	//testReverseLinkList()
	//fmt.Println("*************第二种方法**************")
	//testReverseLinkList2()
	//fmt.Println("*************逆序输出链表**************")
	//testReversePrintLinkList()
	testRemoveRepeatLinkNode()
}

func testReverseLinkList() {
	head := linklist.NewLNode(1)
	node2 := linklist.NewLNode(2)
	node3 := linklist.NewLNode(3)
	node4 := linklist.NewLNode(4)
	node5 := linklist.NewLNode(5)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head.Print()
	head = linklist.ReverseLinkList(head)
	head.Print()
}

func testReverseLinkList2() {
	head := linklist.NewLNode(1)
	node2 := linklist.NewLNode(2)
	node3 := linklist.NewLNode(3)
	node4 := linklist.NewLNode(4)
	node5 := linklist.NewLNode(5)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head.Print()
	head = linklist.ReverseLinkList2(head)
	head.Print()
}

func testReversePrintLinkList() {
	head := linklist.NewLNode(1)
	node2 := linklist.NewLNode(2)
	node3 := linklist.NewLNode(3)
	node4 := linklist.NewLNode(4)
	node5 := linklist.NewLNode(5)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head.Print()
	fmt.Println("*********逆序输出链表*********")
	linklist.ReversePrint(head)
}

func testRemoveRepeatLinkNode() {
	head := linklist.NewLNode(1)
	node2 := linklist.NewLNode(3)
	node3 := linklist.NewLNode(1)
	node4 := linklist.NewLNode(5)
	node5 := linklist.NewLNode(5)
	node6 := linklist.NewLNode(7)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	head.Print()
	linklist.RemoveRepeatListNode(head)
	head.Print()
}
