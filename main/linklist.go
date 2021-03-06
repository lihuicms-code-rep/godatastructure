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
	//testRemoveRepeatLinkNode()
	//testRemoveRepeatLinkNode2()
	//testRemoveRepeatInSortedLinkList()
	//testLinkListSum1()
	//testLinkListSum2()
	//testReSortLinkList()
	//testKthNode()
	//testHasCircle()
	//testPairReverse()
	//testKGroupReverse()
	//testMergeTwoLinkList()
	//testRemoveNode()
	TestGetIntersectionNode()
}

//测试反转链表
func testReverseLinkList() {
	head := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head.Print()
	head = linklist.ReverseLinkList(head)
	head.Print()
}

//测试反转链表第二种方式
func testReverseLinkList2() {
	head := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head.Print()
	head = linklist.ReverseLinkList2(head)
	head.Print()
}

//测试倒序打印链表
func testReversePrintLinkList() {
	head := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	head.Print()
	fmt.Println("*********逆序输出链表*********")
	linklist.ReversePrint(head)
}

//测试删除无序链表中重复的结点
func testRemoveRepeatLinkNode() {
	head := linklist.NewListNode(1)
	node2 := linklist.NewListNode(3)
	node3 := linklist.NewListNode(1)
	node4 := linklist.NewListNode(5)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(7)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	head.Print()
	linklist.RemoveRepeatListNode(head)
	head.Print()
}

//测试删除无序链表中重复的结点
//方法2
func testRemoveRepeatLinkNode2() {
	head := linklist.NewListNode(1)
	node2 := linklist.NewListNode(3)
	node3 := linklist.NewListNode(1)
	node4 := linklist.NewListNode(5)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(7)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	head.Print()
	linklist.RemoveRepeatListNode2(head)
	head.Print()
}

//测试删除有序链表中重复的结点
func testRemoveRepeatInSortedLinkList() {
	head := linklist.NewListNode(1)
	node2 := linklist.NewListNode(1)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(5)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(7)
	head.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	head.Print()
	linklist.RemoveRepeateListNodeInSorted(head)
	head.Print()
}

//测试单链表相加:方法一,可能会溢出的情况
func testLinkListSum1() {
	head1 := linklist.NewListNode(3)
	node2 := linklist.NewListNode(4)
	node3 := linklist.NewListNode(5)
	node4 := linklist.NewListNode(6)
	node5 := linklist.NewListNode(7)
	node6 := linklist.NewListNode(9)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	fmt.Println("*************链表1*************")
	head1.Print()

	fmt.Println("*************链表2*************")
	head2 := linklist.NewListNode(9)
	node22 := linklist.NewListNode(8)
	node32 := linklist.NewListNode(2)
	node42 := linklist.NewListNode(6)
	node52 := linklist.NewListNode(5)
	head2.Next = node22
	node22.Next = node32
	node32.Next = node42
	node42.Next = node52
	head2.Print()

	fmt.Println("*************求和的链表**********")
	newHead := linklist.AddTwoNumbers(head1, head2)
	newHead.Print()

	//特殊用例
	head3 := linklist.NewListNode(0)
	head4 := linklist.NewListNode(0)
	newHead2 := linklist.AddTwoNumbers(head3, head4)
	fmt.Println("************新表************")
	newHead2.Print()


	head5 := linklist.NewListNode(1)
	node552 := linklist.NewListNode(0)
	node53 := linklist.NewListNode(9)
	head5.Next = node552
	node552.Next = node53

	head6 := linklist.NewListNode(5)
	node62 := linklist.NewListNode(7)
	node63 := linklist.NewListNode(8)
	head6.Next = node62
	node62.Next = node63

	newHead3 := linklist.AddTwoNumbers(head5, head6)
	fmt.Println("************新表2************")
	newHead3.Print()
}

//测试单链表相加:方法二,链表加法
func testLinkListSum2() {
	head1 := linklist.NewListNode(3)
	node2 := linklist.NewListNode(4)
	node3 := linklist.NewListNode(5)
	node4 := linklist.NewListNode(6)
	node5 := linklist.NewListNode(7)
	node6 := linklist.NewListNode(9)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	fmt.Println("*************链表1*************")
	head1.Print()

	fmt.Println("*************链表2*************")
	head2 := linklist.NewListNode(9)
	node22 := linklist.NewListNode(8)
	node32 := linklist.NewListNode(2)
	node42 := linklist.NewListNode(6)
	node52 := linklist.NewListNode(5)
	head2.Next = node22
	node22.Next = node32
	node32.Next = node42
	node42.Next = node52
	head2.Print()

	fmt.Println("*************求和的链表**********")
	newHead := linklist.AddTwoNumbers2(head1, head2)
	newHead.Print()

	//特殊用例
	head3 := linklist.NewListNode(0)
	head4 := linklist.NewListNode(0)
	newHead2 := linklist.AddTwoNumbers2(head3, head4)
	fmt.Println("************新表************")
	newHead2.Print()


	head5 := linklist.NewListNode(1)
	node552 := linklist.NewListNode(0)
	node53 := linklist.NewListNode(9)
	head5.Next = node552
	node552.Next = node53

	head6 := linklist.NewListNode(5)
	node62 := linklist.NewListNode(7)
	node63 := linklist.NewListNode(8)
	head6.Next = node62
	node62.Next = node63

	newHead3 := linklist.AddTwoNumbers2(head5, head6)
	fmt.Println("************新表2************")
	newHead3.Print()
}

//测试重排链表
func testReSortLinkList() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(6)
	node7 := linklist.NewListNode(7)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("******当前链表情况********")
	head1.Print()

	linklist.ReorderList(head1)
	fmt.Println("******重排链表************")
	head1.Print()
}


//测试链表的倒数第k个结点
func testKthNode() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(6)
	node7 := linklist.NewListNode(7)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("******当前链表情况********")
	head1.Print()

	val1 := linklist.KthToLast(head1, 1)
	fmt.Printf("倒数第一个结点为:%d\n", val1)

	val2 := linklist.KthToLast(head1, 2)
	fmt.Printf("倒数第二个结点为:%d\n", val2)

	val3 := linklist.KthToLast(head1, 3)
	fmt.Printf("倒数第三个结点为:%d\n", val3)

	val4 := linklist.KthToLast(head1, 4)
	fmt.Printf("倒数第四个结点为:%d\n", val4)

	val5 := linklist.KthToLast(head1, 5)
	fmt.Printf("倒数第四个结点为:%d\n", val5)

	val6 := linklist.KthToLast(head1, 6)
	fmt.Printf("倒数第六个结点为:%d\n", val6)

	val7 := linklist.KthToLast(head1, 7)
	fmt.Printf("倒数第七个结点为:%d\n", val7)

	val8 := linklist.KthToLast(head1, 8)
	fmt.Printf("倒数第七个结点为:%d\n", val8)
}

//测试链表中是否有环
func testHasCircle() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(6)
	node7 := linklist.NewListNode(7)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("******head1当前链表情况********")
	head1.Print()

	hasC := linklist.HasCycle(head1)
	fmt.Printf("方法1测试 head1 是否有环:%t\n", hasC)

	hasC = linklist.HasCycle2(head1)
	fmt.Printf("方法2测试 head1 是否有环:%t\n", hasC)

	head2 := linklist.NewListNode(1)
	node22 := linklist.NewListNode(2)
	node23 := linklist.NewListNode(3)
	node24 := linklist.NewListNode(4)
	node25 := linklist.NewListNode(5)
	node26 := linklist.NewListNode(6)
	node27 := linklist.NewListNode(7)
	head2.Next = node22
	node22.Next = node23
	node23.Next = node24
	node24.Next = node25
	node25.Next = node26
	node26.Next = node27
	node27.Next = node25

	fmt.Println("******head2 当前链表情况********")
	hasC = linklist.HasCycle(head2)
	fmt.Printf("方法1 测试 head2 是否有环:%t\n", hasC)

	hasC = linklist.HasCycle2(head2)
	fmt.Printf("方法2 测试 head2 是否有环:%t\n", hasC)
}

//测试两两交换
func testPairReverse() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(6)
	node7 := linklist.NewListNode(7)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("******当前链表情况********")
	head1.Print()

	newHead := linklist.SwapPairs(head1)
	fmt.Println("******两两交换后的链表情况********")
	newHead.Print()
}

//测试K个一组进行翻转
func testKGroupReverse() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(6)
	node7 := linklist.NewListNode(7)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("******当前链表情况********")
	head1.Print()

	newHead := linklist.ReverseKGroup(head1, 2)
	fmt.Println("******两个一组进行交换后的链表情况********")
	newHead.Print()
}

//测试合并两个有序链表
func testMergeTwoLinkList() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(4)
	head1.Next = node2
	node2.Next = node3
	fmt.Println("**********有序链表1*************")
	head1.Print()

	head2 := linklist.NewListNode(1)
	node4 := linklist.NewListNode(3)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(9)
	head2.Next = node4
	node4.Next = node5
	node5.Next = node6

	fmt.Println("**********有序链表2*************")
	head2.Print()

	head3 := linklist.MergeTwoLists(head1, head2)
	fmt.Println("**********合并后的链表***********")
	head3.Print()
}

//测试删除结点
func testRemoveNode() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	node4 := linklist.NewListNode(4)
	node5 := linklist.NewListNode(5)
	node6 := linklist.NewListNode(6)
	node7 := linklist.NewListNode(7)
	head1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7

	fmt.Println("******当前链表情况********")
	head1.Print()

	linklist.RemoveNode(node5)
	fmt.Println("******删除结点5后当前链表情况********")
	head1.Print()
}

//测试单链表是否交叉,并返回交叉点
func TestGetIntersectionNode() {
	head1 := linklist.NewListNode(1)
	node2 := linklist.NewListNode(2)
	node3 := linklist.NewListNode(3)
	head1.Next = node2
	node2.Next = node3

	head2 := linklist.NewListNode(4)
	node4 := linklist.NewListNode(5)
	head2.Next = node4

	node5 := linklist.NewListNode(6)
	node6 := linklist.NewListNode(7)
	node7 := linklist.NewListNode(8)
	node8 := linklist.NewListNode(9)
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8

	node3.Next = node5
	node4.Next = node5

	fmt.Println("****第一条链表********")
	head1.Print()

	fmt.Println("****第二条链表********")
	head2.Print()

	iNode := linklist.GetIntersectionNode(head1, head2)
	fmt.Printf("intersection node:%v\n", iNode)

	//
	head3 := linklist.NewListNode(1)
	node31 := linklist.NewListNode(2)
	node32 := linklist.NewListNode(3)
	head3.Next = node31
	node31.Next = node32

	fmt.Println("*******第三条链表******")
	head3.Print()

	fmt.Println("*******第四条链表******")
	head4 := linklist.NewListNode(1)
	head4.Print()

	iNode = linklist.GetIntersectionNode(head3, head4)
	fmt.Printf("intersection node:%v\n", iNode)



}
