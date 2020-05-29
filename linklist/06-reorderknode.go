package linklist

//题目6:单链表中的倒数第k个结点
/*
  找出单链表中的倒数第k个元素，例如给定单链表：1->2->3->4->5->6->7，则单链表的倒数第k=3个元素为5。
 */

//解法1:先遍历一遍得知链表长度为n,倒数第k个结点就是正数的:n-k+1

//解法2:双指针
func KthToLast(head *ListNode, k int) int {
	if head == nil || k <= 0 {
		return -1
	}

	p1, p2 := head, head

	for i := 1; i <= k; i++ {      //注意k的范围
		if p2 != nil {
			p2 = p2.Next
		}
	}

	if p2 == nil { //走到尾部
		return p1.Val
	}

	for {
		if p2 == nil {
			break
		}

		p2 = p2.Next
		p1 = p1.Next
	}

	return p1.Val
}
